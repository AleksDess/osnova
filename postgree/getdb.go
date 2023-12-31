package postgree

import (
	"database/sql"
	"fmt"
)

var GcarDB *sql.DB

func RunDB() (err error) {
	GcarDB, err = GetDB()
	return
}

var PostHost = ""
var PostPort = ""
var PostPass = ""
var PostName = ""

func GetDB() (db *sql.DB, err error) {

	connStr := fmt.Sprintf("host=%s port=%s user=postgres dbname='%s' password=%s sslmode=disable", PostHost, PostPort, PostName, PostPass)

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("err postgree not connect", err)
		return
	}
	return
}
