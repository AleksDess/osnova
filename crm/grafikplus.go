package crm

import (
	"database/sql"
	"fmt"
	"osnova/times"
	"time"
)

type CarPlus struct {
	Car      string
	Vin      string
	Mapon_id int64
}

type CarPlusSQL struct {
	Car      sql.NullString
	Vin      sql.NullString
	Mapon_id sql.NullInt64
}

type ListCarPlus []GrafikPlus
type ListCarPlusSQL []GrafikPlusSQL

type GrafikPlus struct {
	City            string
	Date            string
	Status          string
	IsLatestVersion bool
	IsDeleted       bool
	Car             string
	Vin             string
	Mapon_id        int64
	Model           string
	Driver          string
	InnerStatus     string
	Driver_Accaunt  string
	Driver_Accaunts []string
	DayOff          bool
	Remont          bool
	DTP             bool
}

type GrafikPlusSQL struct {
	City            sql.NullString
	Status          sql.NullString
	IsLatestVersion sql.NullBool
	IsDeleted       sql.NullBool
	Car             sql.NullString
	Vin             sql.NullString
	Mapon_id        sql.NullInt64
	Model           sql.NullString
	Driver          sql.NullString
	InnerStatus     sql.NullString
	DriverAccaunt   sql.NullString
}

type ListGrafikPlus []GrafikPlus
type ListGrafikPlusSQL []GrafikPlusSQL

func (a *CarPlus) Create(b *GrafikPlusSQL) {

	if b.Car.Valid {
		a.Car = b.Car.String
	}
	if b.Vin.Valid {
		a.Vin = b.Vin.String
	}
	if b.Mapon_id.Valid {
		a.Mapon_id = b.Mapon_id.Int64
	}
}

func (a *GrafikPlus) Create(b *GrafikPlusSQL) {
	if b.City.Valid {
		a.City = b.City.String
	}
	if b.Status.Valid {
		a.Status = b.Status.String
	}
	if b.IsLatestVersion.Valid {
		a.IsLatestVersion = b.IsLatestVersion.Bool
	} else {
		a.IsLatestVersion = true
	}
	if b.IsDeleted.Valid {
		a.IsDeleted = b.IsDeleted.Bool
	}
	if b.Car.Valid {
		a.Car = b.Car.String
	}
	if b.Vin.Valid {
		a.Vin = b.Vin.String
	}
	if b.Mapon_id.Valid {
		a.Mapon_id = b.Mapon_id.Int64
	}
	if b.Model.Valid {
		a.Model = b.Model.String
	}
	if b.Driver.Valid {
		a.Driver = b.Driver.String
	}
	if b.InnerStatus.Valid {
		a.InnerStatus = b.InnerStatus.String
	}
	if b.DriverAccaunt.Valid {
		a.Driver_Accaunt = b.DriverAccaunt.String
	}
}

func (a *ListCarPlus) Create(b *ListCarPlusSQL) {
	for _, g := range *b {
		rs := GrafikPlus{}
		rs.Create(&g)
		*a = append(*a, rs)
	}
}

func (a *ListGrafikPlus) Create(b *ListGrafikPlusSQL, car *ListCarPlus) {

	r := ListGrafikPlus{}
	for _, g := range *b {
		rs := GrafikPlus{}
		rs.Create(&g)
		r = append(r, rs)
	}

	for _, c := range *car {
		rs := GrafikPlus{Vin: c.Vin}
		for _, g := range r {
			if !g.IsLatestVersion || g.IsDeleted {
				continue
			}
			if g.Vin == c.Vin {
				rs.City = g.City
				rs.Car = g.Car
				rs.Status = g.Status
				rs.Mapon_id = g.Mapon_id
				rs.Driver = g.Driver
				rs.Driver_Accaunts = append(rs.Driver_Accaunts, g.Driver_Accaunt)
				if g.Status == "ROAD_ACCIDENT" {
					rs.DTP = true
				}
				if g.Status == "ON_SERVICE_STATION" {
					rs.Remont = true
				}
			}
			if g.Status == "DRIVER_DAY_OFF" {
				if g.Driver == rs.Driver {
					rs.DayOff = true
				}
			}
		}
		*a = append(*a, rs)
	}
}

func (a *ListCarPlus) GetCityCarsDayV1(city string) {
	z := ""
	z += "SELECT DISTINCT\n"
	z += "cars.license_plate,  cars.vin, cars.mapon_id\n"
	z += "FROM cars\n"
	z += fmt.Sprintf("WHERE cars.auto_park_id = '%s'\n", city)
	z += "ORDER BY cars.license_plate\n"

	// fmt.Println(z)

	sqlr := ListCarPlusSQL{}

	rows, err := CrmDB.Query(z)
	if err != nil {
		fmt.Println(err)
		return
	}

	sqlr.ReadRows(rows)
	rows.Close()

	a.Create(&sqlr)
}

func (a *ListGrafikPlus) GetCityDayV1(city string, day time.Time) {

	z := ""
	z += "-- в работе с водилами\n"
	z += "SELECT DISTINCT \n"
	z += "auto_parks.name,\n"
	z += "schedule.event_type, schedule.is_latest_version, schedule.is_deleted,\n"
	z += "cars.license_plate,  cars.vin, cars.mapon_id,\n"
	z += "car_models.model,\n"
	z += "drivers.full_name, drivers.inner_status,\n"
	z += "driver_sync_to_external_ids.integration_type\n"
	z += "FROM schedule\n"
	z += "LEFT JOIN cars \n"
	z += "ON schedule.car_id = cars.id\n"
	z += "JOIN auto_parks\n"
	z += "ON schedule.auto_park_id = auto_parks.id\n"
	z += "LEFT JOIN drivers\n"
	z += "ON schedule.driver_id = drivers.id\n"
	z += "LEFT JOIN driver_sync_to_external_ids\n"
	z += "ON driver_sync_to_external_ids.driver_id = schedule.driver_id\n"
	z += "LEFT JOIN car_models\n"
	z += "ON car_models.id = cars.model_id\n"
	z += fmt.Sprintf("WHERE schedule.auto_park_id = '%s'\n", city)
	z += fmt.Sprintf("AND event_period_start <= '%s'\n", day.Format(times.TNS2))
	z += fmt.Sprintf("AND event_period_end >= '%s'\n", day.Format(times.TNS2))
	z += "ORDER BY schedule.event_type"

	sqlr := ListGrafikPlusSQL{}

	rows, err := CrmDB.Query(z)
	if err != nil {
		fmt.Println(err)
		return
	}

	sqlr.ReadRows(rows)
	rows.Close()

	car := ListCarPlus{}
	car.GetCityCarsDayV1(city)

	a.Create(&sqlr, &car)
}

func (grafiks *ListCarPlusSQL) ReadRows(rows *sql.Rows) {
	for rows.Next() {
		var g GrafikPlusSQL
		err := rows.Scan(
			&g.Car,
			&g.Vin,
			&g.Mapon_id,
		)
		if err != nil {
			fmt.Println(err)
			continue
		}
		*grafiks = append(*grafiks, g)
	}
}

func (grafiks *ListGrafikPlusSQL) ReadRows(rows *sql.Rows) {
	for rows.Next() {
		var g GrafikPlusSQL
		err := rows.Scan(
			&g.City,
			&g.Status,
			&g.IsLatestVersion,
			&g.IsDeleted,
			&g.Car,
			&g.Vin,
			&g.Mapon_id,
			&g.Model,
			&g.Driver,
			&g.InnerStatus,
			&g.DriverAccaunt,
		)
		if err != nil {
			fmt.Println(err)
			continue
		}
		*grafiks = append(*grafiks, g)
	}
}

func (a *GrafikPlus) ToString() (r []string) {
	r = make([]string, 2)
	if a.Status == "BUSY_WITH_PRIVATE_TRADER" {
		if len(a.Driver_Accaunts) == 1 {
			if a.Driver_Accaunts[0] == "BOLT" {
				r[0] = a.Driver
				if a.DayOff {
					r[1] = "Вихідний"
				} else {
					r[1] = a.Driver
				}
				return
			}
			if a.Driver_Accaunts[0] == "UBER" {
				r[0] = "(UB) " + a.Driver
				r[1] = "(UB) " + a.Driver
				return
			}
			if a.Driver_Accaunts[0] == "UKLON" {
				r[0] = "(UK) " + a.Driver
				r[1] = "(UK) " + a.Driver
				return
			}
		} else if len(a.Driver_Accaunts) > 1 {
			if a.Driver_Accaunts[0] == "BOLT" {
				r[0] = a.Driver
			} else if a.Driver_Accaunts[0] == "UBER" {
				r[0] = "(UB) " + a.Driver
			} else if a.Driver_Accaunts[0] == "UKLON" {
				r[0] = "(UK) " + a.Driver
			}
			if a.Driver_Accaunts[1] == "BOLT" {
				r[1] = a.Driver
			} else if a.Driver_Accaunts[1] == "UBER" {
				r[1] = "(UB) " + a.Driver
			} else if a.Driver_Accaunts[1] == "UKLON" {
				r[1] = "(UK) " + a.Driver
			}
			if a.DayOff {
				r[1] = "Вихідний"
			}
			return
		}
	} else if a.DTP {
		r[0] = "ДТП"
		r[1] = "ДТП"
		return
	} else if a.Remont {
		r[0] = "Ремонт"
		r[1] = "Ремонт"
		return
	}
	return
}
