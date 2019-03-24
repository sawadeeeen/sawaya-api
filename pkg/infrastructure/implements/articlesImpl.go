package implements

import (
	"github.com/jmoiron/sqlx"
	"github.com/sawadeeeen/sawaya-api/pkg/domain/model"
)

// Store ...
func Store(articles *model.Articles) error {
	db, err := sqlx.Open("mysql", "root:root@tcp(localhost:3306)/ianBlog")
	if err != nil {
		panic(err.Error())
	}
	//	defer db.Close()

	//	c.Request.ParseForm()

	res, err := db.Prepare(
		`insert into article(title, content, published, published_at) values(?,?,?,now())`,
	)
	//_, err = db.Exec()
	if err != nil {
		panic(err.Error())
	}

	//	_, err = res.Exec(inputData.title, inputData.content, inputData.published, inputData.published_at)
	_, err = res.Exec(&articles.Title, articles.Content, articles.Published, articles.Published_at)
	return err
}
