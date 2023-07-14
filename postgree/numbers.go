package postgree

import (
	"strings"
	"trs"
)

type Numbers struct {
	id   int64
	data []string
}

func GetNumbersCar(mapon int64) (res []string, err error) {
	var r []byte
	row := GcarDB.QueryRow("SELECT data FROM numbers WHERE id = $1", mapon)
	err = row.Scan(&r)
	if err != nil {
		return
	}
	s := string(r)
	s = strings.Trim(s, "{}")
	s = strings.Replace(s, "\",\"", "|", 1)
	s = strings.ReplaceAll(s, "\"", "")
	rs := strings.Split(s, "|")
	for _, i := range rs {
		r := strings.Split(i, ",")
		if len(r) == 3 {
			res = append(res, trs.Trs(r[0]))
		}
	}
	return
}
