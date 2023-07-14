package crm

import (
	"database/sql"
)

// 1. Место хранения аккаунтов црм
// proxy_drivers_assign - прокси связь карточек водителя

// 2. Привязка аккаунтов к драйверу црм
// driver_sync_to_external_ids -  связь карточек водителей с айди службы такси

// 3. Место хранения поездок по службам
// calculated_statements - колонка integration_details

// 4 место хранения пробега авто
// car_odometer_history

type DriverAccaunt struct {
	Id               string
	Driver_id        string
	External_id      string
	Integration_type string
}

type Driver struct {
	Id              string
	Auto_park_id    string
	Full_name       string
	First_name      string
	Last_name       string
	Phone           string
	Email           string
	Inner_status    string
	Contract_number string
	Accaunts        []DriverAccaunt
}

type Drivers []Driver

func GetDriversCityMap(db *sql.DB, city string) (res map[string]Driver, err error) {
	res = make(map[string]Driver)
	r := Drivers{}
	err = r.GetDriverCity(db, city)
	if err != nil {
		return
	}
	for _, i := range r {
		res[i.Id] = i
	}
	return
}

func (driver *Driver) GetId(db *sql.DB, id string) (err error) {
	// Запрос на получение водителя по ID
	row := db.QueryRow("SELECT id, auto_park_id, full_name, first_name, last_name, phone, email, inner_status FROM drivers WHERE id = $1", id)
	err = row.Scan(&driver.Id, &driver.Auto_park_id, &driver.Full_name, &driver.First_name, &driver.Last_name, &driver.Phone, &driver.Email, &driver.Inner_status)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
	}
	return
}

func (driver *Driver) Get(db *sql.DB, contract string) (err error) {
	// Запрос на получение водителя по ID
	row := db.QueryRow("SELECT id, auto_park_id, full_name, first_name, last_name, phone, email, inner_status FROM drivers WHERE contract_number = $1", contract)
	err = row.Scan(&driver.Id, &driver.Auto_park_id, &driver.Full_name, &driver.First_name, &driver.Last_name, &driver.Phone, &driver.Email, &driver.Inner_status)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		return
	}
	driver.Contract_number = contract

	// Запрос на получение всех записей из таблицы driver_sync_to_external_ids
	rows, err := db.Query("SELECT id, driver_id, external_id, integration_type FROM driver_sync_to_external_ids WHERE driver_id = $1", driver.Id)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var driversAccount DriverAccaunt
		err := rows.Scan(&driversAccount.Id, &driversAccount.Driver_id, &driversAccount.External_id, &driversAccount.Integration_type)
		if err != nil {
			continue
		}
		driver.Accaunts = append(driver.Accaunts, driversAccount)
	}
	return
}

func (a *Drivers) GetDriverCity(db *sql.DB, city string) (err error) {
	// Запрос на получение водителя по ID
	row, err := db.Query("SELECT id, auto_park_id, full_name, first_name, last_name, phone, email, inner_status FROM drivers WHERE auto_park_id = $1", city)
	if err != nil {
		return
	}
	for row.Next() {
		driver := Driver{}
		err = row.Scan(&driver.Id, &driver.Auto_park_id, &driver.Full_name, &driver.First_name, &driver.Last_name, &driver.Phone, &driver.Email, &driver.Inner_status)
		if err != nil {
			if err == sql.ErrNoRows {
				return
			}
			return
		}
		*a = append(*a, driver)
	}
	return
}
