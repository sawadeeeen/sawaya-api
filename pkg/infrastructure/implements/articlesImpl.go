package implements

import (
	"github.com/sawadeeeen/sawaya-api/pkg/domain/model"
	"github.com/sawadeeeen/sawaya-api/pkg/infrastructure/database"
)

type ArticlesImpl struct {
	database.SQLHandler
}

// Store ...
//func (article *ArticlesImpl) Store(articles model.Articles) error {
func Store(articles model.Articles) error {
	db := database.NewSQLHandler()
	//	defer db.Close()

	//	c.Request.ParseForm()

	//	inputData := model.Articles{
	//id := c.Request.Form["id"],
	//		title:        c.Request.Form["title"],
	//		content:      c.Request.Form["content"],
	//		published:    c.Request.Form["published"],
	//		published_at: c.Request.Form["published_at"],
	//	}

	res, err := db.Prepare(
		`insert into article(title, content, published, published_at) values(?,?,?,?,?)`,
	)
	//_, err = db.Exec()
	if err != nil {
		panic(err.Error())
	}

	//	_, err = res.Exec(inputData.title, inputData.content, inputData.published, inputData.published_at)
	_, err = res.Exec(articles.Title, articles.Content, articles.Published, articles.Published_at)
	return err
}
