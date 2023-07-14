package gcardb

import (
	"database/sql"
	"fmt"
)

type Cars struct {
	Mapon int64
	Nomer string
	Vin   string
	Price int64
	Model string
}

type ListCars []Cars
type MapCars map[int64]Cars

func NewListCars() (r ListCars) {
	return ListCars{}
}

func (a *ListCars) ReadDB(db *sql.DB) (err error) {
	z := "SELECT * FROM cars"
	rows, err := db.Query(z)
	if err != nil {
		return err
	}
	for rows.Next() {
		var id int64
		var r Cars
		err := rows.Scan(&id, &r.Mapon, &r.Nomer, &r.Vin, &r.Price, &r.Model)
		if err != nil {
			fmt.Println(err)
			continue
		}
		*a = append(*a, r)
	}
	return nil
}

func (a MapCars) ReadDB(db *sql.DB) (err error) {
	z := "SELECT * FROM cars"
	rows, err := db.Query(z)
	if err != nil {
		return err
	}
	for rows.Next() {
		var id int64
		var r Cars
		err := rows.Scan(&id, &r.Mapon, &r.Nomer, &r.Vin, &r.Price, &r.Model)
		if err != nil {
			fmt.Println(err)
			continue
		}
		a[r.Mapon] = r
	}
	return nil
}

func (a *Cars) RecDB(db *sql.DB) (err error) {
	_, err = db.Exec("INSERT INTO cars (mapon, nomer, vin, price, model) VALUES ($1, $2, $3)", a.Mapon, a.Nomer, a.Vin, a.Price, a.Model)
	if err != nil {
		return
	}
	return nil
}

func (a *Cars) DelDB(db *sql.DB) (err error) {
	_, err = db.Exec("DELETE FROM cars WHERE mapon = $1", a.Mapon)
	if err != nil {
		return
	}
	return nil
}
