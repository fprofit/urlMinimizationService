package dbpstsql

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

func OpenDB(psgInf string) bool {
	psqlInfo = psgInf + " sslmode=disable"
	db, err := sql.Open("postgres", psqlInfo+" dbname="+dbName)
	if err != nil {
		return false
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		createDB()

	}
	createTab()
	return true
}

// create DB
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
}

func createTab() {
	db, err := sql.Open("postgres", psqlInfo+" dbname="+dbName)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	creaTab := fmt.Sprintf("CREATE TABLE if not exists %s (miniurl varchar(10) UNIQUE NOT NULL, origurl varchar UNIQUE NOT NULL, PRIMARY KEY(miniurl, origurl));", tabName)
	_, err = db.Exec(creaTab)
	if err != nil {
		panic(err)
	}
}
