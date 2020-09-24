package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/miraikeitai2020/backend-summer-vacation/pkg/config"
)

var Con *sql.DB

func init() {
	var err error
	Con, err = sql.Open("mysql", config.GetConnectionToken())
	if err != nil {
		log.Fatal(err)
	}
}
