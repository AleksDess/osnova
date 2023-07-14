package postgree

import (
	"fmt"
	"time"
)

type CarCity struct {
	Day  string
	City string
	Car  int64
}

type ListCarCity []CarCity

func (a *ListCarCity) Print() {
	for _, i := range *a {
		fmt.Println("CarCity: ", i)
	}
}

// записать элемент
func (a *CarCity) Rec() {
	t, _ := time.Parse("02.01.2006", a.Day)
	created_time := time.Date(t.Year(), t.Month(), t.Day(), 8, 0, 0, 0, time.Local)

	_, err := GcarDB.Exec("INSERT INTO carcity (day, city, car, created_time) VALUES ($1, $2, $3, $4)", a.Day, a.City, a.Car, created_time)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Data inserted successfully.")
	}
}

// прочитать город за дату
func (a *ListCarCity) GetCityDay(day, city string) {

	z := fmt.Sprintf("SELECT day, city, car FROM carcity WHERE day = '%s' AND city = '%s'", day, city)
	//fmt.Println(z)

	rows, err := GcarDB.Query(z)
	if err != nil {
		fmt.Println(err)
		return
	}

	for rows.Next() {
		r := CarCity{}
		err := rows.Scan(&r.Day, &r.City, &r.Car)
		if err != nil {
			fmt.Println(err)
			continue
		}
		*a = append(*a, r)
	}
}
