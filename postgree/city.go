package postgree

import (
	"Citys"
	"database/sql"
)

func CreateListCitys(db *sql.DB) (r Citys.ListCity, err error) {
	err = r.ReadListCity(db)
	return
}
