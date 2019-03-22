package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type SQLHandler struct {
	Conn *sqlx.DB
}

func NewSQLHandler() *SQLHandler {
	db, err := sqlx.Open("mysql", "root:root@tcp(localhost:3306)/ianBlog")
	if err != nil {
		panic(err.Error())
	}
	sqlHandler := new(SQLHandler)
	sqlHandler.Conn = db
	return sqlHandler
}

// Prepare ...
func (s *SQLHandler) Prepare(query string) (*sql.Stmt, error) {
	result, err := s.Conn.Prepare(query)
	if err != nil {
		return nil, err
	}
	return result, nil
}
