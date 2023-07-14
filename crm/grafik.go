package crm

import (
	"database/sql"
	"fmt"
	"log"
	"osnova/times"
	"time"
)

// запрос для списка статусов
// SELECT DISTINCT event_type
// FROM schedule

// "event_type"
// "AUTO_POUND"
// "ROAD_ACCIDENT"
// "ON_SERVICE_STATION"
// "RENTAL"
// "BUSY_WITH_CREW"
// "OTHER"
// "DRIVER_DAY_OFF"
// "BUSY_WITH_PRIVATE_TRADER"

type GrafikSQL struct {
	Id                 sql.NullString
	Auto_park_id       sql.NullString
	Company_id         sql.NullString
	Car_id             sql.NullString
	Parent_id          sql.NullString
	Event_type         sql.NullString
	Event_period_start sql.NullTime
	Event_period_end   sql.NullTime
	Comment            sql.NullString
	Created_at         sql.NullTime
	Updated_at         sql.NullTime
	Driver_mode        sql.NullString
	Is_latest_version  sql.NullBool
	Public_id          sql.NullString
	Is_deleted         sql.NullBool
	Edited_by_user     sql.NullString
	Driver_id          sql.NullString
	Shift_type         sql.NullString
	Rent_debt          sql.NullFloat64
}

type Grafik struct {
	Id                 string
	Auto_park_id       string
	Company_id         string
	Car_id             string
	Parent_id          string
	Event_type         string
	Event_period_start time.Time
	Event_period_end   time.Time
	Comment            string
	Created_at         time.Time
	Updated_at         time.Time
	Driver_mode        string
	Is_latest_version  bool
	Public_id          string
	Is_deleted         bool
	Edited_by_user     string
	Driver_id          string
	Driver_            Driver
	Shift_type         string
	Rent_debt          sql.NullFloat64
}

func (a *Grafik) Create(db *sql.DB, b *GrafikSQL) {

	_, offset := time.Now().Zone()

	if b.Id.Valid {
		a.Id = b.Id.String
	}
	if b.Auto_park_id.Valid {
		a.Auto_park_id = b.Auto_park_id.String
	}
	if b.Company_id.Valid {
		a.Company_id = b.Company_id.String
	}
	if b.Car_id.Valid {
		a.Car_id = b.Car_id.String
	}
	if b.Parent_id.Valid {
		a.Parent_id = b.Parent_id.String
	}
	if b.Event_type.Valid {
		a.Event_type = b.Event_type.String
	}
	if b.Event_period_start.Valid {
		a.Event_period_start = b.Event_period_start.Time
		a.Event_period_start = a.Event_period_start.Add(time.Duration(offset) * time.Second)
	}
	if b.Event_period_end.Valid {
		a.Event_period_end = b.Event_period_end.Time
		a.Event_period_end = a.Event_period_end.Add(time.Duration(offset) * time.Second)
	}
	if b.Comment.Valid {
		a.Comment = b.Comment.String
	}
	if b.Created_at.Valid {
		a.Created_at = b.Created_at.Time
		a.Created_at = a.Created_at.Add(time.Duration(offset) * time.Second)
	}
	if b.Updated_at.Valid {
		a.Updated_at = b.Updated_at.Time
		a.Updated_at = a.Updated_at.Add(time.Duration(offset) * time.Second)
	}
	if b.Driver_mode.Valid {
		a.Driver_mode = b.Driver_mode.String
	}
	if b.Is_latest_version.Valid {
		a.Is_latest_version = b.Is_latest_version.Bool
	}
	if b.Public_id.Valid {
		a.Public_id = b.Public_id.String
	}
	if b.Is_deleted.Valid {
		a.Is_deleted = b.Is_deleted.Bool
	}
	if b.Edited_by_user.Valid {
		a.Edited_by_user = b.Edited_by_user.String
	}
	if b.Driver_id.Valid {
		a.Driver_id = b.Driver_id.String
	}
	if b.Shift_type.Valid {
		a.Shift_type = b.Shift_type.String
	}
	if b.Rent_debt.Valid {
		a.Rent_debt = b.Rent_debt
	}
	// a.Driver_.GetId(db, a.Driver_id)
}

func GetGrafik(db *sql.DB, city string, t time.Time) {
	fmt.Println(t.Format(times.TNS2))

	cars, err := GetCarCityMap(db, city)
	if err != nil {
		fmt.Println(err)
	}
	drivers, err := GetDriversCityMap(db, city)
	if err != nil {
		fmt.Println(err)
	}

	rows, err := db.Query("select * from schedule where auto_park_id = $1 AND event_period_start <= $2 AND event_period_end >= $2", city, t.Format(times.TNS2))
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var grafiks []GrafikSQL
	for rows.Next() {
		var g GrafikSQL
		err = rows.Scan(
			&g.Id,
			&g.Auto_park_id,
			&g.Company_id,
			&g.Car_id,
			&g.Parent_id,
			&g.Event_type,
			&g.Event_period_start,
			&g.Event_period_end,
			&g.Comment,
			&g.Created_at,
			&g.Updated_at,
			&g.Driver_mode,
			&g.Is_latest_version,
			&g.Public_id,
			&g.Is_deleted,
			&g.Edited_by_user,
			&g.Driver_id,
			&g.Shift_type,
			&g.Rent_debt,
		)
		if err != nil {
			log.Fatal(err)
		}
		grafiks = append(grafiks, g)
	}

	for _, g := range grafiks {
		r := Grafik{}
		r.Create(db, &g)
		fmt.Println(r.Created_at.Format(times.TNSF))
		fmt.Println(r.Event_period_start.Format(times.TNSF))
		fmt.Println(r.Event_period_end.Format(times.TNSF))
		fmt.Println(r.Event_type)
		fmt.Println(cars[r.Car_id].License_plate)
		fmt.Println(cars[r.Car_id].Mapon_id)
		fmt.Println(drivers[r.Driver_id].Full_name)
		fmt.Println(r.Driver_id)
		fmt.Println()
	}
}
