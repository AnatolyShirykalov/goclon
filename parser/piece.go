package parser

import (
	"../app/models"
	"github.com/PuerkitoBio/goquery"
)

func piece(doc *goquery.Document, c *models.Composer) (ret *models.Piece) {
	ret = &models.Piece{}
	s := doc.Find(".composer_name a").Last()
	url, exist := s.Attr("href")
	if !exist {
		panic("Нет ссылки на пьесу")
	}
	models.DB.FirstOrInit(ret, models.Piece{ClassicOnlineId: url[15:]})
	ret.Name = s.Text()
	ret.Composer = *c
	models.DB.Save(ret)
	return
}
