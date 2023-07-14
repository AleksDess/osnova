package postgree

import "fmt"

func GetSummaFloat64(z string) (r float64) {
	row := GcarDB.QueryRow(z)
	row.Scan(&r)
	return
}

func GetSummaInt(z string) (r int) {
	row := GcarDB.QueryRow(z)
	err := row.Scan(&r)
	fmt.Println(err)
	return
}
