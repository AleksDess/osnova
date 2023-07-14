package gcardb

import (
	"fmt"
	"osnova/logger"
	"osnova/postgree"
	"osnova/times"
	"time"
	"trs"
)

func ViewNumbersCarDays(n int) {
	day := times.Last_days(n)
	type work struct {
		city string
		day  time.Time
		car  int
		nom  int
	}
	rs := make([]work, 0)
	cit := make([]string, 0)
	dt := make([]string, 0)
	for _, d := range day {
		dt = append(dt, d.Format(times.TNS))
		z := fmt.Sprintf("SELECT city, COUNT(car), Count(nomer) FROM shedule WHERE day = '%s' GROUP BY city", d.Format(times.TNS))
		rows, err := postgree.GcarDB.Query(z)
		if err != nil {
			logger.InfoLog.Println(err)
			return
		}
		for rows.Next() {
			r := work{}
			rows.Scan(&r.city, &r.car, &r.nom)
			if !trs.Comp_string(r.city, cit) {
				cit = append(cit, r.city)
			}
			r.day = d
			rs = append(rs, r)
		}
	}

	for _, c := range cit {
		fmt.Print(c)
		for _, d := range dt {
			for _, i := range rs {
				if i.city == c && i.day.Format(times.TNS) == d {
					fmt.Print(" ", i.car)
					break
				}
			}
		}
		fmt.Println()
	}

}
