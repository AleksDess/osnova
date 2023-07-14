package gcardb

import (
	"database/sql"
	"fmt"
	"osnova/times"
	"time"
)

type CarCity struct {
	Day   string
	City  string
	Mapon int64
}

type ListCarCity []CarCity

func NewListCarCity() (r ListCarCity) {
	return ListCarCity{}
}

func (a *ListCarCity) ReadDBCar(db *sql.DB, mapon int64) (err error) {
	z := fmt.Sprintf("SELECT * FROM shedule.\"carcity\" WHERE car = %d", mapon)
	rows, err := db.Query(z)
	if err != nil {
		return err
	}
	for rows.Next() {
		var id int64
		var r CarCity
		err := rows.Scan(&id, &r.Day, &r.City, &r.Mapon)
		if err != nil {
			continue
		}
		*a = append(*a, r)
	}
	return nil
}

func (a *ListCarCity) ReadDBCarCityDays(db *sql.DB, city string, days []time.Time) (err error) {
	z := fmt.Sprintf("SELECT * FROM shedule.\"carcity\" WHERE  city = '%s' AND (", city)
	for n, d := range days {
		if n == 0 {
			z += fmt.Sprintf("day = '%s'", d.Format(times.TNS))
		} else {
			z += fmt.Sprintf(" OR day = '%s'", d.Format(times.TNS))
		}
	}
	z += ")"
	rows, err := db.Query(z)
	if err != nil {
		return err
	}
	for rows.Next() {
		var id int64
		var r CarCity
		err := rows.Scan(&id, &r.Day, &r.City, &r.Mapon)
		if err != nil {
			continue
		}
		*a = append(*a, r)
	}
	return nil
}

func (a *CarCity) RecDB(db *sql.DB) (err error) {
	_, err = db.Exec("INSERT INTO shedule.\"carcity\" (day, city, car) VALUES ($1, $2, $3)", a.Day, a.City, a.Mapon)
	if err != nil {
		return
	}
	return nil
}

func (a *ListCarCity) InfoPark(db *sql.DB, days []time.Time) (r ListInfoPark) {

	car := MapCars{}
	car = make(map[int64]Cars)
	car.ReadDB(db)
	park := ""

	for _, d := range days {
		count := 0
		price := 0
		for _, i := range *a {
			if i.Day == d.Format(times.TNS) {
				price += int(car[i.Mapon].Price)
				count++
				park = i.City
			}
		}
		// fmt.Println(d.Format(times.TNS), count, price)
		if count > 0 {
			r = append(r, InfoPark{Day: d.Format(times.TNS), City: park, NumbersOfCars: count, PricePark: float64(price)})
		}
	}
	return
}
