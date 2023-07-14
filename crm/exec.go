package crm

import "fmt"

func GetSummaFloat64(z string) (r float64) {
	row := CrmDB.QueryRow(z)
	row.Scan(&r)
	return
}

func GetSummaInt(z string) (r int) {
	row := CrmDB.QueryRow(z)
	err := row.Scan(&r)
	fmt.Println(err)
	return
}
