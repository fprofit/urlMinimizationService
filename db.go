package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var (
	psqlInfo string
	dbName   = "dbminiandorigurl"
	tabName  = "tabminiandorigurl"
)

func getMiniToOrig(origUrl string) (miniUrl string) {
	conf := psqlInfo + " dbname=" + dbName
	db, err := sql.Open("postgres", conf)
	if err != nil {
		return "false"
	}
	defer db.Close()

	s := fmt.Sprintf("SELECT miniUrl FROM %s WHERE origUrl='%s';", tabName, origUrl)
	err = db.QueryRow(s).Scan(&miniUrl)
	if err != nil {
		miniUrl = UrlGenerator(origUrl)
		insDB := fmt.Sprintf(`INSERT INTO %s (miniUrl, origUrl)	
		VALUES($1, $2)`, tabName)
		_, err := db.Exec(insDB, miniUrl, origUrl)
		if err != nil {
			miniUrl = "false"
		}
	}
	return
}

func getOrigToMini(miniUrl string) (origUrl string) {
	conf := psqlInfo + " dbname=" + dbName
	db, err := sql.Open("postgres", conf)
	if err != nil {
		return "false"
	}
	defer db.Close()

	s := fmt.Sprintf("SELECT origUrl FROM %s WHERE miniUrl='%s';", tabName, miniUrl)
	err = db.QueryRow(s).Scan(&origUrl)
	if err != nil {
		return "false"
	}
	return
}

func openDB() bool {
	db, err := sql.Open("postgres", psqlInfo+" dbname="+dbName)
	if err != nil {
		return false
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		createDB()
	}
	return true
}

func createDB() {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	_, err = db.Exec("create database " + dbName)
	if err != nil {
		panic(err)
	}
	createTab()
}

func createTab() {
	db, err := sql.Open("postgres", psqlInfo+" dbname="+dbName)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	creaTab := fmt.Sprintf("CREATE TABLE %s (miniurl varchar(10) UNIQUE NOT NULL, origurl varchar UNIQUE NOT NULL, PRIMARY KEY(miniurl, origurl));", tabName)
	_, err = db.Exec(creaTab)
	if err != nil {
		panic(err)
	}
}
