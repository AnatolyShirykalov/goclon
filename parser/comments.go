package parser

import (
	"../app/models"
	"github.com/PuerkitoBio/goquery"
	"github.com/pbnjay/strptime"
	"gopkg.in/resty.v1"
	"regexp"
	"strconv"
	"time"
)

func comments(client *resty.Client, doc *goquery.Document, perform *models.Perform) {
	doc.Find(".comments .comment").Each(func(i int, s *goquery.Selection) {
		data, _ := s.Find("a[name]").First().Attr("name")
		id := data[8:]
		comment := &models.Comment{}
		models.DB.FirstOrInit(comment, models.Comment{ClassicOnlineId: id})
		comment.User = *commenter(s)
		comment.Perform = *perform
		comment.Date = commentDate(s)
		comment.Text = commentText(s)
		comment.Likes = commentLikes(s, id)
		models.DB.Save(comment)
	})
	nl := doc.Find("#NextLink")
	if nl.Length() == 0 {
		return
	}
	u, _ := nl.First().Attr("href")
	comments(client, GetDoc(client, "http://classic-online.ru/archive/"+u), perform)
}

func commentDate(s *goquery.Selection) (ret *time.Time) {
	txt := s.Find("div").First().Text()
	re := regexp.MustCompile(`\d{2}\.\d{2}\.\d{4} +\d{2}:\d{2}`)
	dateStr := re.FindString(txt)
	format := `%d.%m.%Y %H:%M %z`
	tm, err := strptime.Parse(dateStr+" +0300", format)
	if err != nil {
		panic(err)
	}
	ret = &tm
	return
}

func commenter(s *goquery.Selection) (user *models.User) {
	userA := s.Find("a[href]").FilterFunction(func(i int, ss *goquery.Selection) bool {
		if href, _ := ss.Attr("href"); len(href) >= 13 && href[:13] == "/users/?item=" {
			return true
		}
		return false
	}).First()
	url, _ := userA.Attr("href")
	user = &models.User{}
	models.DB.FirstOrInit(user, models.User{ClassicOnlineId: url[13:]})
	user.Name = userA.Text()
	models.DB.Save(user)
	return
}

func commentText(s *goquery.Selection) (ret string) {
	ret = s.Find("td.cont").First().Text()
	return
}

func commentLikes(s *goquery.Selection, id string) (ret int64) {
	txt := s.Find("#agrees_cnt_" + id).First().Text()
	re := regexp.MustCompile(`\d+`)
	str := re.FindString(txt)
	l, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	return l
}
