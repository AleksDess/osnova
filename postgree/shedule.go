package postgree

import (
	"fmt"
	"strings"
	"time"
)

type Shedule struct {
	Day    string
	City   string
	Car    int64
	Nomer  string
	V1     string
	V2     string
	Status string
}

// Печать элементаShedule
func (a *Shedule) Print() {
	fmt.Println()
	fmt.Println("Shedule             :")
	fmt.Println("Day                 :", a.Day)
	fmt.Println("City                :", a.City)
	fmt.Println("Car                 :", a.Car)
	fmt.Println("Nomer               :", a.Nomer)
	fmt.Println("V1                  :", a.V1)
	fmt.Println("V2                  :", a.V2)
	fmt.Println("Status              :", a.Status)
}

type ListShedule []Shedule

func (a *ListShedule) Rec() (err error) {
	// Построение текста запроса
	valueStrings := []string{}
	valueArgs := []interface{}{}
	i := 0
	for _, g := range *a {
		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d)", i+1, i+2, i+3, i+4, i+5, i+6, i+7, i+8))
		t, _ := time.Parse("02.01.2006", g.Day)
		created_time := time.Date(t.Year(), t.Month(), t.Day(), 8, 0, 0, 0, time.Local)
		valueArgs = append(valueArgs, g.Day, g.City, g.Car, g.Nomer, g.V1, g.V2, g.Status, created_time)
		i += 8
	}
	query := fmt.Sprintf("INSERT INTO shedule (day, city, car, nomer, v1, v2, status, time)  VALUES  %s",
		strings.Join(valueStrings, ", "))
	_, err = GcarDB.Exec(query, valueArgs...)
	if err != nil {
		return
	}
	fmt.Println("Save ListShedule succesfull!!!")
	return nil
}

func (a *ListShedule) Print() {
	for _, i := range *a {
		i.Print()
	}
}

// записать элемент
func (a *Shedule) Rec() {
	t, _ := time.Parse("02.01.2006", a.Day)
	created_time := time.Date(t.Year(), t.Month(), t.Day(), 8, 0, 0, 0, time.Local)
	_, err := GcarDB.Exec("INSERT INTO shedule (day, city, car, nomer, v1, v2, status, time) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)", a.Day, a.City, a.Car, a.Nomer, a.V1, a.V2, a.Status, created_time)
	if err != nil {
		fmt.Println(err)
	}
}

func ShedyleDelCityDay(day, city string) {
	_, err := GcarDB.Exec("DELETE FROM shedule WHERE day = $1 AND city = $2", day, city)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Data %s city %s deleted successfully.\n", day, city)
	}
}

func ShedyleDelDay(day string) (err error) {
	_, err = GcarDB.Exec("DELETE FROM shedule WHERE day = $1 ", day)
	if err != nil {
		return
	}
	return nil
}

// func RecGrafikPeriod() {
// 	for _, i := range times.List_Day("01.01.2023", "28.06.2023") {

// 		for _, cit := range ListCity {

// 			gr := Get_grafik(cit.Name, []time.Time{i})
// 			postgree.ShedyleDelCityDay(i.Format(TNS), cit.Id_BGQ)

// 			cc := postgree.ListCarCity{}
// 			cc.GetCityDay(i.Format(TNS), cit.Id_BGQ)

// 			for _, i := range cc {
// 				fl := false
// 				for _, j := range gr {
// 					if i.Car == j.id_mapon {
// 						fl = true
// 						// fmt.Println()
// 						// fmt.Println(i, j)
// 						r := postgree.Shedule{}
// 						r.Day = i.Day
// 						r.City = i.City
// 						r.Car = i.Car
// 						r.Nomer = j.car_nomer
// 						r.V1 = j.name_vod1
// 						r.V2 = j.name_vod2
// 						r.Status = j.Status()
// 						// r.Print()
// 						r.Rec()
// 						break
// 					}
// 				}
// 				if !fl {
// 					fmt.Println(i, "------------------------")
// 					numcars, _ := postgree.GetNumbersCar(i.Car)
// 					fmt.Println(numcars)
// 					flag := false
// 					for _, nc := range numcars {
// 						if flag {
// 							break
// 						}
// 						for _, j := range gr {
// 							if flag {
// 								break
// 							}
// 							if nc == j.car_nomer {
// 								fl = true
// 								r := postgree.Shedule{}
// 								r.Day = i.Day
// 								r.City = i.City
// 								r.Car = i.Car
// 								r.Nomer = j.car_nomer
// 								r.V1 = j.name_vod1
// 								r.V2 = j.name_vod2
// 								r.Status = j.Status()
// 								r.Print()
// 								r.Rec()
// 								flag = true
// 								break
// 							}
// 						}
// 					}
// 					if !flag {
// 						fmt.Println(i, "*********************")
// 					}
// 				}
// 			}
// 		}
// 	}
// }
