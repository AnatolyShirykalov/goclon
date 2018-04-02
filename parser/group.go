package parser

import (
	"../app/models"
	"github.com/PuerkitoBio/goquery"
)

func group(doc *goquery.Document) (group *models.Group) {
	s := doc.Find(".performer_name a")
	names := make([]string, s.Length())
	urls := s.Map(func(i int, se *goquery.Selection) string {
		url, exists := se.Attr("href")
		if !exists {
			panic("Нет ссылки на исполнителя")
		}
		names[i] = se.Text()
		return url
	})
	//fmt.Println(names, urls)

	performerIDs := make([]int64, len(urls))
	performers := make([]*models.Performer, len(urls))
	for i, url := range urls {
		performer := &models.Performer{}
		models.DB.FirstOrInit(performer, models.Performer{ClassicOnlineId: url})
		performer.Name = names[i]
		models.DB.Save(performer)
		performerIDs[i] = performer.ID
		performers[i] = performer
	}
	rows, err := models.DB.Table("group_performers").
		Select("group_id, count(*)").
		Where("performer_id IN (?)", performerIDs).
		Group("group_id").
		Rows()
	if err != nil {
		panic(err)
	}
	group = &models.Group{}
	found := false
	for rows.Next() {
		var groupID, count int64
		if err := rows.Scan(&groupID, &count); err != nil {
			panic(err)
		}
		if count == int64(len(performerIDs)) {
			models.DB.Where(models.Group{ID: groupID}).First(group)
			found = true
		}
	}
	if found == false {
		group = &models.Group{}
		models.DB.Create(group)
		for _, performer := range performers {
			gp := &models.GroupPerformer{}
			gp.Performer = *performer
			gp.Group = *group
			models.DB.Save(gp)
		}
	}

	return
}
