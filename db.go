package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var psqlInfo string

func getMiniToOrig(origUrl string) (miniUrl string) {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return "false"
	}
	defer db.Close()

	s := fmt.Sprintf("SELECT miniUrl FROM miniurltorigrurl_table WHERE origUrl='%s';", origUrl)
	err = db.QueryRow(s).Scan(&miniUrl)
	if err != nil {
		miniUrl = UrlGenerator(origUrl)
		insDB := `
		INSERT INTO miniurltorigrurl_table (miniUrl, origUrl)
		VALUES($1, $2)`
		_, err := db.Exec(insDB, miniUrl, origUrl)
		if err != nil {
			miniUrl = "false"
		}
	}
	return
}

func getOrigToMini(miniUrl string) (origUrl string) {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return "false"
	}
	defer db.Close()

	s := fmt.Sprintf("SELECT origUrl FROM miniurltorigrurl_table WHERE miniUrl='%s';", miniUrl)
	err = db.QueryRow(s).Scan(&origUrl)
	if err != nil {
		return "false"
	}
	return
}

func createTab() {
	db, err := sql.Open("postgres", "user=postgres password=1234 host=localhost sslmode=disable")
	if err != nil {
		fmt.Println("errr")
	}
	defer db.Close()
	dbName := "MiniUrlAndOrigUrl"
	_, err = db.Exec("create database " + dbName)
	if err != nil {
		//handle the error
		log.Fatal(err)
	}
	//dbname=MiniUrlToOrigUrl

	_, err = db.Exec("CREATE TABLE miniurltorigrurl (miniUrl varchar(10) UNIQUE NOT NULL, origUrl varchar UNIQUE NOT NULL, PRIMARY KEY(miniUrl, origUrl));")
	if err != nil {
		fmt.Println("errr2")
		log.Fatal(err)
	}
}
