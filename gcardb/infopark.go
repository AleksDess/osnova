package gcardb

import (
	"database/sql"
	"fmt"
)

type InfoPark struct {
	Day           string
	City          string
	NumbersOfCars int
	PricePark     float64
}

type ListInfoPark []InfoPark
type MapInfoPark map[string]InfoPark

func NewListInfoPark() (r ListInfoPark) {
	return ListInfoPark{}
}

func (a *ListInfoPark) ReadDB(db *sql.DB) (err error) {
	z := "SELECT * FROM shedule.\"infopark\""
	rows, err := db.Query(z)
	if err != nil {
		return err
	}
	for rows.Next() {
		var id int64
		var r InfoPark
		err := rows.Scan(&id, &r.Day, &r.City, &r.NumbersOfCars, &r.PricePark)
		if err != nil {
			fmt.Println(err)
			continue
		}
		*a = append(*a, r)
	}
	return nil
}

func (a MapInfoPark) ReadDB(db *sql.DB) (err error) {
	z := "SELECT * FROM shedule.\"infopark\""
	rows, err := db.Query(z)
	if err != nil {
		return err
	}
	for rows.Next() {
		var id int64
		var r InfoPark
		err := rows.Scan(&id, &r.Day, &r.City, &r.NumbersOfCars, &r.PricePark)
		if err != nil {
			fmt.Println(err)
			continue
		}
		a[r.Day+"|"+r.City] = r
	}
	return nil
}

func (a *InfoPark) RecDB(db *sql.DB) (err error) {
	_, err = db.Exec("INSERT INTO shedule.\"infopark\" (day, city, number_of_cars, price_park) VALUES ($1, $2, $3, $4)", a.Day, a.City, a.NumbersOfCars, a.PricePark)
	if err != nil {
		return
	}
	return nil
}

func (a *InfoPark) DelDB(db *sql.DB) (err error) {
	_, err = db.Exec("DELETE FROM shedule.\"infopark\" WHERE day = '$1' AND city = '$2'", a.Day, a.City)
	if err != nil {
		return
	}
	return nil
}
