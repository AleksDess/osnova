package postgree

import (
	"fmt"
	"osnova/times"
	"time"
)

func AverageCarMounth(days []time.Time, city string) (first int, res float64) {
	cars := 0
	for n, i := range days {

		z := fmt.Sprintf("SELECT COUNT(*) FROM shedule WHERE city = '%s' AND day = '%s'", city, i.Format(times.TNS))
		fmt.Println(z)
		car := GetSummaInt(z)
		fmt.Println(city, i.Format(times.TNS), car)
		cars += car
		if n == 0 {
			first = car
		}
	}
	res = float64(cars) / float64(len(days))
	return
}
