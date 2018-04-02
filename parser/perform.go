package parser

import (
	"../app/models"
	"github.com/PuerkitoBio/goquery"
	"github.com/pbnjay/strptime"
	"regexp"
	"strconv"
	"time"
)

func perform(doc *goquery.Document, piece *models.Piece, group *models.Group) (ret *models.Perform) {
	ret = &models.Perform{}
	url, _ := doc.Find(`a[href]`).FilterFunction(func(i int, s *goquery.Selection) bool {
		if href, _ := s.Attr("href"); len(href) >= 18 && href[:18] == "/archive/?file_id=" {
			return true
		}
		return false
	}).First().Attr("href")
	models.DB.FirstOrInit(ret, models.Perform{ClassicOnlineId: url[18:]})
	ret.Piece = *piece
	ret.Group = *group
	ret.User = *uploader(doc)
	ret.Date = date(doc)
	ret.Likes = likes(doc, url[18:])
	models.DB.Save(ret)
	return
}

func likes(doc *goquery.Document, id string) int64 {
	txt := doc.Find("#thanks_cnt_" + id).First().Text()
	re := regexp.MustCompile(`\d+`)
	str := re.FindString(txt)
	l, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	return l
}

func date(doc *goquery.Document) (ret *time.Time) {
	txt := doc.Find("div.listen").First().Text()
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

func uploader(doc *goquery.Document) (ret *models.User) {
	ret = &models.User{}
	a := doc.Find("a.add_to_playlist").First()
	url, _ := a.Attr("href")
	if url == "" {
		models.DB.FirstOrCreate(ret, models.User{Name: "Администратор"})
		return
	}
	models.DB.FirstOrInit(ret, models.User{ClassicOnlineId: url[13:]})
	ret.Name = a.Text()
	models.DB.Save(ret)
	return
}
