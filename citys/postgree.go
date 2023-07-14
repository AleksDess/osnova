package citys

import (
	"database/sql"
	"encoding/json"
)

func (a *ListCity) RecCitys(db *sql.DB) (err error) {
	for _, cit := range *a {
		err = cit.UpCity(db)
		if err != nil {
			return
		}
	}
	return nil
}

func (a *City) RecCity(db *sql.DB) (err error) {

	jsonData, err := json.Marshal(a)
	if err != nil {
		return
	}

	_, err = db.Exec("INSERT INTO city (id, data) VALUES ($1, $2)", a.Name, jsonData)
	if err != nil {
		return
	}
	return nil
}

func (a *City) UpCity(db *sql.DB) (err error) {

	jsonData, err := json.Marshal(a)
	if err != nil {
		return
	}

	_, err = db.Exec("UPDATE city SET data = $2 WHERE id = $1", a.Name, jsonData)
	if err != nil {
		return
	}
	return nil
}

func (a *ListCity) ReadListCity(db *sql.DB) (err error) {
	z := "SELECT * FROM public.city"

	rows, err := db.Query(z)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id string
		var data []byte
		err = rows.Scan(&id, &data)
		if err != nil {
			return
		}

		cit := City{}
		err = json.Unmarshal(data, &cit)
		if err != nil {
			return
		}

		*a = append(*a, cit)
	}
	return nil
}

func (a *City) ReadCity(db *sql.DB, cit string) (err error) {

	row := db.QueryRow("SELECT data FROM public.city WHERE id = $1", cit)

	var data []byte
	err = row.Scan(&data)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &a)
	if err != nil {
		return
	}
	return nil
}
