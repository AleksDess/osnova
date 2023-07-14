package crm

import (
	"database/sql"
	"fmt"
	"time"
)

type CarSQL struct {
	id                       sql.NullString
	auto_park_id             sql.NullString
	company_id               sql.NullString
	manually_car_id          sql.NullString
	license_plate            sql.NullString
	vin                      sql.NullString
	registration_certificate sql.NullString
	place_quantity           sql.NullInt16
	color                    sql.NullString
	issue_year               sql.NullInt16
	created_at               sql.NullTime
	updated_at               sql.NullTime
	added_by_integration_id  sql.NullString
	added_by_user_id         sql.NullString
	is_duplicate             sql.NullBool
	model_id                 sql.NullString
	rent_price               sql.NullFloat64
	mapon_id                 sql.NullInt32
	odometer                 sql.NullInt32
	speed                    sql.NullInt32
	speed_updated_at         sql.NullTime
	high_rent_price          sql.NullFloat64
}

type Car struct {
	Id                       string
	Auto_park_id             string
	Company_id               string
	Manually_car_id          string
	License_plate            string
	Vin                      string
	Registration_certificate string
	Place_quantity           int
	Color                    string
	Issue_year               int
	Created_at               time.Time
	Updated_at               time.Time
	Added_by_integration_id  string
	Added_by_user_id         string
	Is_duplicate             bool
	Model_id                 string
	Rent_price               float64
	Mapon_id                 int
	Odometer                 int
	Speed                    int
	Speed_updated_at         time.Time
	High_rent_price          float64
}

// Печать элементаCar
func (a *Car) Print() {
	fmt.Println()
	fmt.Println("Car                 :")
	fmt.Println("Id                  :", a.Id)
	fmt.Println("Auto_park_id        :", a.Auto_park_id)
	fmt.Println("Company_id          :", a.Company_id)
	fmt.Println("Manually_car_id     :", a.Manually_car_id)
	fmt.Println("License_plate       :", a.License_plate)
	fmt.Println("Vin                 :", a.Vin)
	fmt.Println("Registration_certificate:", a.Registration_certificate)
	fmt.Println("Place_quantity      :", a.Place_quantity)
	fmt.Println("Color               :", a.Color)
	fmt.Println("Issue_year          :", a.Issue_year)
	fmt.Println("Created_at          :", a.Created_at)
	fmt.Println("Updated_at          :", a.Updated_at)
	fmt.Println("Added_by_integration_id:", a.Added_by_integration_id)
	fmt.Println("Added_by_user_id    :", a.Added_by_user_id)
	fmt.Println("Is_duplicate        :", a.Is_duplicate)
	fmt.Println("Model_id            :", a.Model_id)
	fmt.Println("Rent_price          :", a.Rent_price)
	fmt.Println("Mapon_id            :", a.Mapon_id)
	fmt.Println("Odometer            :", a.Odometer)
	fmt.Println("Speed               :", a.Speed)
	fmt.Println("Speed_updated_at    :", a.Speed_updated_at)
	fmt.Println("High_rent_price     :", a.High_rent_price)
}

// Печать элементаCar
func (a *Car) SmPrint() {
	fmt.Println()
	fmt.Println("Car                 :")
	fmt.Println("License_plate       :", a.License_plate)
	fmt.Println("Vin                 :", a.Vin)
	fmt.Println("Mapon_id            :", a.Mapon_id)
}

type Cars []Car

func GetCarCityMap(db *sql.DB, park string) (res map[string]Car, err error) {
	res = make(map[string]Car)
	r := Cars{}
	err = r.GetCarCity(db, park)
	if err != nil {
		return
	}
	for _, i := range r {
		res[i.Id] = i
	}
	return
}

func (a *Cars) GetCarCity(db *sql.DB, park string) (err error) {

	rows, err := db.Query("SELECT * FROM cars where auto_park_id = $1", park)
	if err != nil {
		return
	}
	defer rows.Close()

	var cars []CarSQL
	for rows.Next() {
		var car CarSQL
		err = rows.Scan(
			&car.id,
			&car.auto_park_id,
			&car.company_id,
			&car.manually_car_id,
			&car.license_plate,
			&car.vin,
			&car.registration_certificate,
			&car.place_quantity,
			&car.color,
			&car.issue_year,
			&car.created_at,
			&car.updated_at,
			&car.added_by_integration_id,
			&car.added_by_user_id,
			&car.is_duplicate,
			&car.model_id,
			&car.rent_price,
			&car.mapon_id,
			&car.odometer,
			&car.speed,
			&car.speed_updated_at,
			&car.high_rent_price,
		)
		if err != nil {
			fmt.Println(err)
			continue
		}
		cars = append(cars, car)
	}

	for _, car := range cars {
		r := Car{}
		r.Create(&car)
		*a = append(*a, r)
		//r.SmPrint()
	}
	return
}

func (a *Car) Create(b *CarSQL) {

	_, offset := time.Now().Zone()

	if b.id.Valid {
		a.Id = b.id.String
	}
	if b.auto_park_id.Valid {
		a.Auto_park_id = b.auto_park_id.String
	}
	if b.company_id.Valid {
		a.Company_id = b.company_id.String
	}
	if b.manually_car_id.Valid {
		a.Manually_car_id = b.manually_car_id.String
	}
	if b.license_plate.Valid {
		a.License_plate = b.license_plate.String
	}
	if b.vin.Valid {
		a.Vin = b.vin.String
	}
	if b.registration_certificate.Valid {
		a.Registration_certificate = b.registration_certificate.String
	}
	if b.place_quantity.Valid {
		a.Place_quantity = int(b.place_quantity.Int16)
	}
	if b.color.Valid {
		a.Color = b.color.String
	}
	if b.issue_year.Valid {
		a.Issue_year = int(b.issue_year.Int16)
	}
	if b.created_at.Valid {
		a.Created_at = b.created_at.Time
		a.Created_at = a.Created_at.Add(time.Duration(offset) * time.Second)
	}
	if b.updated_at.Valid {
		a.Updated_at = b.updated_at.Time
		a.Updated_at = a.Updated_at.Add(time.Duration(offset) * time.Second)
	}
	if b.added_by_integration_id.Valid {
		a.Added_by_integration_id = b.added_by_integration_id.String
	}
	if b.added_by_user_id.Valid {
		a.Added_by_user_id = b.added_by_user_id.String
	}
	if b.is_duplicate.Valid {
		a.Is_duplicate = b.is_duplicate.Bool
	}
	if b.model_id.Valid {
		a.Model_id = b.model_id.String
	}
	if b.rent_price.Valid {
		a.Rent_price = b.rent_price.Float64
	}
	if b.mapon_id.Valid {
		a.Mapon_id = int(b.mapon_id.Int32)
	}
	if b.odometer.Valid {
		a.Odometer = int(b.odometer.Int32)
	}
	if b.speed.Valid {
		a.Speed = int(b.speed.Int32)
	}
	if b.speed_updated_at.Valid {
		a.Speed_updated_at = b.speed_updated_at.Time
		a.Speed_updated_at = a.Speed_updated_at.Add(time.Duration(offset) * time.Second)
	}
	if b.high_rent_price.Valid {
		a.High_rent_price = b.high_rent_price.Float64
	}
}
