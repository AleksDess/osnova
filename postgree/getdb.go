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

func GetDB() (db *sql.DB, err error) {

	connStr := "host=92.119.231.174 port=5432 user=postgres dbname='Gcar compani' password=fd5d6f9v2de4g58i4k25F1DR2G1 sslmode=disable"

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("err postgree not connect", err)
		return
	}
	return
}
