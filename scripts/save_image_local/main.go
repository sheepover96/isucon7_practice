package main

import (
	"log"
	"path"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"io/ioutil"
)

const BASE_PATH = "/home/isucon/isubata/webapp/public/icons"

func main() {
	dsn := "isucon:isucon@tcp(localhost:3306)/isubata?charset=utf8mb4&parseTime=true&loc=Local"
	db, err := sqlx.Connect("mysql", dsn)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	rows,err := db.Query("select `name`, `data` from image")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var name string
	var data []byte
	for rows.Next() {
		err2 := rows.Scan(&name, &data)
		if err2 != nil {
			log.Println(err2)
		}
		filePath := path.Join(BASE_PATH, name)
		err3 := ioutil.WriteFile(filePath, data, 0666)
		if err3 != nil {
			log.Fatal(err3)
		}
	}
}
