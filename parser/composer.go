package parser

import (
	"../app/models"
	"github.com/PuerkitoBio/goquery"
)

func composer(doc *goquery.Document) (ret *models.Composer) {
	ret = &models.Composer{}
	s := doc.Find(".composer_name a").First() //.Attr("href")
	url, exist := s.Attr("href")
	if !exist {
		panic("Нет ссылки на композитора")
	}
	models.DB.FirstOrInit(ret, models.Composer{ClassicOnlineId: url})
	if ret == nil {
		println("url", url)
		panic("Не получилось найти или создать композитора")
	}
	ret.Name = s.Text()
	models.DB.Save(ret)
	return
}
