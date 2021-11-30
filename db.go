package main

import (
	"database/sql"
	"fmt"

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

// CREATE TABLE miniurltorigrurl_table (
// 	miniUrl varchar(10) UNIQUE NOT NULL,
//  origUrl varchar NOT UNIQUE NULL,
// PRIMARY KEY(miniUrl, origUrl));
