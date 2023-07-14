package speedover

import (
	"fmt"
	"osnova/imap"
	"osnova/times"
	"strings"
	"time"
)

type AdressSpeedOver struct {
	Adress   string
	Link     string
	City     bool // в городе
	Overcity bool // за городом
	In_work  bool // в работе
}

type AdressSpeedDB struct {
	Adress   string
	Link     string
	City     int
	Overcity int
	In_work  int
}

type ListAdressSpeedOver []AdressSpeedOver

func (a *ListAdressSpeedOver) Print() {
	for _, i := range *a {
		i.Print()
	}
}

func (a *ListAdressSpeedOver) СheckФvailability(s string) bool {
	for _, i := range *a {
		if i.Adress == s {
			return true
		}
	}
	return false
}

func (a *AdressSpeedOver) IsCity() bool {
	r := a.Adress
	r = strings.Trim(r, " ")
	s := strings.Split(r, ", ")
	if len(s) == 0 {
		return false
	}
	b := []byte(s[len(s)-1])
	if len(b) != 5 {
		return false
	}
	for _, i := range b {
		if i < 48 || i > 57 {
			return false
		}
	}
	return true
}

// Печать элементаAdressSpeedOver
func (a *AdressSpeedOver) Print() {
	fmt.Println()
	fmt.Println("AdressSpeedOver     :")
	fmt.Println("Adress              :", a.Adress)
	fmt.Println("Link                :", a.Link)
	fmt.Println("City                :", a.City)
	fmt.Println("Overcity            :", a.Overcity)
	fmt.Println("In_work             :", a.In_work)
}

// Создание таблицы БД
// для записи элементов AdressSpeedOver
// CREATE TABLE Adress (
//     Adress      VARCHAR(40) PRIMARY KEY,
//     Link        VARCHAR(40),
//     City        INTEGER,
//     Overcity    INTEGER,
//     In_work     INTEGER
// );

type SpeedOver struct {
	Adress              string
	Time                time.Time
	Date                string
	City                string
	Car                 string
	Link                string
	Car_Id              int64
	Speed               int
	Under_consideration bool // на рассмотрении
	In_the_city         bool // в городе
	In_work             bool // в работе
	Mapon_Tested        bool // прошел проверку мапон
}

type ListSpeedOver []SpeedOver

type SpeedDB struct {
	Adress              string
	Time                string
	Date                string
	City                string
	Car                 string
	Link                string
	Car_Id              int
	Speed               int
	Under_consideration int
	In_the_city         int
	In_work             int
	Mapon_Tested        int
}

// Печать элементаSpeed
func (a *SpeedOver) Print() {
	fmt.Println()
	fmt.Println("Speed               :")
	fmt.Println("Adress              :", a.Adress)
	fmt.Println("Time                :", a.Time.Format(times.TNSF))
	fmt.Println("Date                :", a.Date)
	fmt.Println("City                :", a.City)
	fmt.Println("Car                 :", a.Car)
	fmt.Println("Link                :", a.Link)
	fmt.Println("Car_Id              :", a.Car_Id)
	fmt.Println("Speed               :", a.Speed)
	fmt.Println("Under_consideration :", a.Under_consideration)
	fmt.Println("In_the_city         :", a.In_the_city)
	fmt.Println("In_work             :", a.In_work)
	fmt.Println("Mapon_Tested        :", a.Mapon_Tested)
}

// Создание таблицы БД
// для записи элементов Speed
// CREATE TABLE Speed (
// id INTEGER PRIMARY KEY AUTOINCREMENT,
// Adress                            VARCHAR(40) ,
// Time                              VARCHAR(20),
// Date                              VARCHAR(12),
// City                              VARCHAR(20),
// Car                               VARCHAR(10),
// Link                              VARCHAR(40),
// Car_Id                            INTEGER,
// Speed                             INTEGER,
// Under_consideration               INTEGER,
// In_the_city                       INTEGER,
// In_work                           INTEGER,
// Mapon_Tested                      INTEGER
// );

// Восстановление элемента Speed
// из строки БД
func (a *AdressSpeedOver) Reestablish(p *AdressSpeedDB) {
	a.Adress = p.Adress
	a.Link = p.Link
	if p.City == 1 {
		a.City = true
	}
	if p.Overcity == 1 {
		a.Overcity = true
	}
	if p.In_work == 1 {
		a.In_work = true
	}
}

// Восстановление элемента Speed
// из строки БД
func (a *SpeedOver) Reestablish(p *SpeedDB) {
	a.Adress = p.Adress
	a.Time, _ = time.Parse(times.TNSF, p.Time)
	a.Date = p.Time
	a.City = p.City
	a.Car = p.Car
	a.Link = p.Link
	a.Car_Id = int64(p.Car_Id)
	a.Speed = p.Speed
	if p.Under_consideration == 1 {
		a.Under_consideration = true
	}
	if p.In_the_city == 1 {
		a.In_the_city = true
	}
	if p.In_work == 1 {
		a.In_work = true
	}
	if p.Mapon_Tested == 1 {
		a.Mapon_Tested = true
	}
}

func (a *AdressSpeedOver) CreateSpeed(p *imap.MessageSpeed) {
	a.Adress = p.Adress
	a.Link = p.Link
}
