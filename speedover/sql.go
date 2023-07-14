package speedover

import (
	"database/sql"
	"fmt"
	"osnova/times"
	"strings"
	"time"
)

// Запись элемента Speed в БД
func (a *SpeedOver) RecDB(db *sql.DB) {

	var d1, d2, d3, d4 int
	if a.Under_consideration {
		d1 = 1
	}
	if a.In_the_city {
		d2 = 1
	}
	if a.In_work {
		d3 = 1
	}
	if a.Mapon_Tested {
		d4 = 1
	}

	name := "Adress, Time, Date, City, Car, Link, Car_Id, Speed, Under_consideration, In_the_city, In_work, Mapon_Tested"
	recdb := fmt.Sprintf("INSERT INTO Speed (%s) VALUES ('%s', '%s', '%s', '%s', '%s','%s', %d, %d, %d, %d, %d, %d)", name, a.Adress, a.Time, a.Date, a.City, a.Car, a.Link, a.Car_Id, a.Speed, d1, d2, d3, d4)

	//fmt.Println(recdb)
	_, err := db.Exec(recdb)
	if err != nil {
		fmt.Println(err, recdb)
	}
}

// Запись элемента Speed в БД
func (a *AdressSpeedOver) RecDB(db *sql.DB) {

	var d1, d2, d3 int
	if a.City {
		d1 = 1
	}
	if a.Overcity {
		d2 = 1
	}
	if a.In_work {
		d3 = 1
	}

	z := fmt.Sprintf("INSERT INTO Adress (Adress, Link, City, Overcity, In_work) VALUES ('%s', '%s', %d, %d, %d)", a.Adress, a.Link, d1, d2, d3)
	fmt.Println(z)

	_, err := db.Exec(z)
	if err != nil {
		if strings.Contains(fmt.Sprint(err), "UNIQUE constraint failed") {
			return
		}
		fmt.Println(err, z)
	}
}

func (a *ListAdressSpeedOver) ReadDb(db *sql.DB) {

	zap := "SELECT * FROM Adress"

	rows, err := db.Query(zap)
	if err != nil {
		fmt.Println(err, zap)
		return
	}

	for rows.Next() {
		p := AdressSpeedDB{}
		err = rows.Scan(&p.Adress, &p.Link, &p.City, &p.Overcity, &p.In_work)
		if err != nil {
			fmt.Println("---------", err)
		}
		r := AdressSpeedOver{}
		r.Reestablish(&p)
		*a = append(*a, r)
	}
}

func (a *ListAdressSpeedOver) ReadDbCity(db *sql.DB, cit string) {

	z := fmt.Sprintf("SELECT * FROM Adress WHERE Adress LIKE '|%s|'", cit)
	z = strings.ReplaceAll(z, "|", "%")
	fmt.Println(z)

	rows, err := db.Query(z)
	if err != nil {
		fmt.Println(err, z)
		return
	}

	for rows.Next() {
		p := AdressSpeedDB{}
		err = rows.Scan(&p.Adress, &p.Link, &p.City, &p.Overcity, &p.In_work)
		if err != nil {
			fmt.Println("---------", err)
		}
		r := AdressSpeedOver{}
		r.Reestablish(&p)
		*a = append(*a, r)
	}
}

func (a *AdressSpeedOver) DelDb(db *sql.DB) {

	z := fmt.Sprintf("DELETE FROM Adress WHERE Adress = '%s'", a.Adress)
	fmt.Println(z)
	_, err := db.Exec(z)
	if err != nil {
		fmt.Println(err, z)
		return
	}

}

func (a *AdressSpeedOver) ReadDbAdress(db *sql.DB, adress string) {

	zap := fmt.Sprintf("SELECT * FROM Speed WHERE Adress = '%s'", adress)

	row := db.QueryRow(zap)
	p := SpeedDB{}
	err := row.Scan(&p)
	if err != nil {
		fmt.Println(err)
	}
	r := SpeedOver{}
	r.Reestablish(&p)
}

func (a *ListSpeedOver) ReadDb(db *sql.DB, date []time.Time) {

	zap := "SELECT * FROM Speed WHERE "
	for n, t := range date {
		if n == 0 {
			zap = zap + fmt.Sprintf("Date = '%s'", t.Format(times.TNS))
		} else {
			zap = zap + fmt.Sprintf(" OR Date = '%s'", t.Format(times.TNS))
		}
	}

	rows, err := db.Query(zap)
	if err != nil {
		fmt.Println(err, zap)
		return
	}

	for rows.Next() {
		p := SpeedDB{}
		err = rows.Scan(&p)
		if err != nil {
			fmt.Println(err)
		}
		r := SpeedOver{}
		r.Reestablish(&p)
		*a = append(*a, r)
	}
}

func (a *ListSpeedOver) ReadDbCity(db *sql.DB, date []time.Time, cit string) {

	zap := fmt.Sprintf("SELECT * FROM Speed WHERE City = '%s' AND ", cit)
	for n, t := range date {
		if n == 0 {
			zap = zap + fmt.Sprintf("Date = '%s'", t.Format(times.TNS))
		} else {
			zap = zap + fmt.Sprintf(" OR Date = '%s'", t.Format(times.TNS))
		}
	}

	rows, err := db.Query(zap)
	if err != nil {
		fmt.Println(err, zap)
		return
	}

	for rows.Next() {
		p := SpeedDB{}
		err = rows.Scan(&p)
		if err != nil {
			fmt.Println(err)
		}
		r := SpeedOver{}
		r.Reestablish(&p)
		*a = append(*a, r)
	}
}
