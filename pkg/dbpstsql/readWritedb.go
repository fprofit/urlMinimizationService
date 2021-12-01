package dbpstsql

import (
	"database/sql"
	"fmt"

	"github.com/fprofit/urlMinimizationService/tree/development/pkg/miniurlgenerator"
	_ "github.com/lib/pq"
)

// get MiniUrl to OrigUrl & create
func GetMiniToOrig(origUrl string) (miniUrl string) {
	conf := psqlInfo + " dbname=" + dbName
	db, err := sql.Open("postgres", conf)
	if err != nil {
		return "false"
	}
	defer db.Close()

	s := fmt.Sprintf("SELECT miniUrl FROM %s WHERE origUrl='%s';", tabName, origUrl)
	err = db.QueryRow(s).Scan(&miniUrl)
	if err != nil {
		miniUrl = miniurlgenerator.UrlGenerator(origUrl)
		insDB := fmt.Sprintf(`INSERT INTO %s (miniUrl, origUrl)	
		VALUES($1, $2)`, tabName)
		_, err := db.Exec(insDB, miniUrl, origUrl)
		if err != nil {
			miniUrl = "false"
		}
	}
	return
}

func GetOrigToMini(miniUrl string) (origUrl string) {
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
