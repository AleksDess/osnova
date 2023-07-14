package times

import (
	"fmt"
	"sort"
	"time"
)

const TNS = "02.01.2006"
const TNSF = "02.01.2006 15:04:05"
const TNSFS = "02.01.2006 15:04"
const TNS2 = "2006-01-02"
const TNSF2 = "2006-01-02 15:04:05"
const TNSF3 = "2006-01-02 15:04"

const TimeBitrix = "2006-01-02T15:04:05-07:00"

// // Получить текущее локальное время
// localTime := time.Now()
// fmt.Println("Local Time:", localTime)

// // Преобразовать локальное время в UTC
// utcTime := localTime.UTC()
// fmt.Println("UTC Time:", utcTime)

// // Создать произвольную дату и время в локальной временной зоне
// specificLocalTime := time.Date(2023, 6, 30, 12, 0, 0, 0, time.Local)
// fmt.Println("Specific Local Time:", specificLocalTime)

// // Преобразовать произвольную дату и время в UTC
// specificUTCTime := specificLocalTime.UTC()
// fmt.Println("Specific UTC Time:", specificUTCTime)

// Возврат текущего времени
func Tek_time() (tsrt string, th, tm, tn int) {
	t := time.Now()
	return t.Format("02.01.2006 15:04"), t.Hour(), t.Minute(), int(t.Weekday())
}

// отдает список дней недели
func Ned(n int) []time.Time {
	t := time.Now()
	_, _, _, tn := Tek_time()
	if tn == 0 {
		tn = 7
	}
	day := make([]time.Time, 0)
	tnach := t.Add(time.Duration((1 - tn)) * 24 * time.Hour)
	tnach = tnach.Add(time.Duration(-n*7) * 24 * time.Hour)
	for i := 0; i < 7; i++ {
		day = append(day, tnach)
		tnach = tnach.Add(24 * time.Hour)
	}
	return day
}

// отдает список дней месяца
func Mounth(n, year int) []time.Time {

	tn := time.Date(year, time.Month(n), 1, 3, 7, 7, 7, time.UTC)

	day := make([]time.Time, 0)

	for i := 0; i < 33; i++ {
		vm := int(tn.Month())
		y := tn.Year()
		if vm == n && y == year {
			day = append(day, tn)
		}
		tn = tn.Add(24 * time.Hour)
	}
	return day
}

// отдает список дней предыдущего месяца
func Previos_mounth() []time.Time {
	t := time.Now()
	y, _ := t.ISOWeek()
	tn := time.Date(y, t.Month(), 1, 3, 7, 7, 7, time.UTC)
	tn = tn.AddDate(0, -1, 0)
	tk := tn.AddDate(0, 1, -1)
	ran := tk.Day() - tn.Day()
	day := make([]time.Time, 0)
	day = append(day, tn)
	for i := 0; i < ran; i++ {
		tn = tn.Add(24 * time.Hour)
		day = append(day, tn)
	}
	return day
}

// отдает список дней текущего месяца
// c сегодняшним днем
func Mounth_until_today() []time.Time {

	t := time.Now()
	y, _ := t.ISOWeek()
	tn := time.Date(y, t.Month(), 1, 3, 7, 7, 7, time.UTC)
	day := make([]time.Time, 0)
	day = append(day, tn)
	for i := 1; i < t.Day(); i++ {
		tn = tn.Add(24 * time.Hour)
		day = append(day, tn)
	}
	return day
}

// отдает список дней текущего месяца
// до сегодняшнего дня
func Mounth_until_yestoday() []time.Time {

	t := time.Now()
	y, _ := t.ISOWeek()
	tn := time.Date(y, t.Month(), 1, 3, 7, 7, 7, time.UTC)
	day := make([]time.Time, 0)
	day = append(day, tn)
	for i := 1; i < t.Day(); i++ {
		tn = tn.Add(24 * time.Hour)
		day = append(day, tn)
	}
	return day[:len(day)-1]
}

// оедает год и номер недели по дате "02.01.2006"
func YN_date(d string) (y, n string, ok bool) {
	t, err := time.Parse(TNS, d)
	if err != nil {
		return
	}
	yy, nn := t.ISOWeek()
	y = fmt.Sprint(yy)
	n = fmt.Sprint(nn)
	ok = true
	return
}

func Date_YN(n, y int) (a, b string) {
	t := time.Date(y, 1, 1, 3, 3, 3, 0, time.UTC)
	// fmt.Println(t.Format(TNSF))
	d := 0
	for {
		d++
		yy, nn := t.ISOWeek()
		if yy == y && nn == n {
			return t.Format(TNS), t.AddDate(0, 0, 6).Format(TNS)
		}
		t = t.AddDate(0, 0, 1)
		if d == 500 {
			break
		}
	}
	return "", ""
}

// диапазон дат вв виде Unix
func Days_Unix(a []time.Time) (n, k int64) {
	t1 := a[0]
	t2 := a[len(a)-1]
	dt1 := time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, time.UTC)
	dt2 := time.Date(t2.Year(), t2.Month(), t2.Day(), 23, 59, 59, 999, time.UTC)
	return dt1.Unix(), dt2.Unix()
}

// диапазон дат вв виде Unix
func Day_Unix(t time.Time) (n, k int64) {

	dt1 := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
	dt2 := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999, time.UTC)
	return dt1.Unix(), dt2.Unix()
}

// отдает список дней между в1 и в2 включительно
func List_Day(d1, d2 string) (res []time.Time) {
	dt1, _ := time.Parse(TNS, d1)
	dt2, _ := time.Parse(TNS, d2)

	if dt1.Before(dt2) {
		res = append(res, dt1)
		for dt1.Before(dt2) {
			dt1 = dt1.Add(24 * time.Hour)
			res = append(res, dt1)
		}
	} else if dt2.Before(dt1) {
		res = append(res, dt2)
		for dt2.Before(dt1) {
			dt2 = dt2.Add(24 * time.Hour)
			res = append(res, dt2)
		}
	} else {
		res = append(res, dt1)
	}
	return
}

func Convert_Date(d string) (s string) {
	t, _ := time.Parse(TNS, d)
	s = t.Format("2006-01-02")
	return
}

func Add_Date(d string, n int) (s string) {
	t, _ := time.Parse(TNS, d)
	t = t.AddDate(0, 0, n)
	s = t.Format(TNS)
	return
}

func Convert_Date_ADD(d string, n int) (s string) {
	t, _ := time.Parse(TNS, d)
	t = t.AddDate(0, 0, n)
	s = t.Format("2006-01-02")
	return
}

// список часов в сутках
func ListHour(t time.Time) (res [25]time.Time) {
	t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 1, 0, time.UTC)
	t = t.Add(-3 * time.Hour)
	for i := 0; i < 25; i++ {
		res[i] = t
		t = t.Add(time.Hour)
	}
	return
}

// перевод даты из дд.мм.гггг
// в формат мм/дд/гггг
func Convert_Date_Reverse(a string) (b string) {
	t, err := time.Parse("2006-01-02", a)
	if err != nil {
		fmt.Println(err)
	}
	b = t.Format(TNS)
	return
}

func Period_days(n, k string) []time.Time {

	tn, _ := time.Parse(TNS, n)
	tk, _ := time.Parse(TNS, k)

	res := make([]time.Time, 0)

	for tk.After(tn.Add(-24 * time.Hour)) {
		res = append(res, tn)
		tn = tn.Add(24 * time.Hour)
	}
	return res
}

func Days_Ned(y, n int) (res []time.Time) {
	// Вычисляем первый день недели
	t := time.Date(y, 0, 0, 0, 0, 0, 0, time.UTC) // .AddDate(0, 0, (14-1)*7+1)

	fl := 0
	for i := 1; i < 1000; i++ {
		y1, n1 := t.ISOWeek()
		if y == y1 && n1 == n {
			fl = 1
			res = append(res, t)
		} else {
			if fl == 1 {
				break
			}
		}
		t = t.AddDate(0, 0, 1)
	}
	return
}

// список дней прошлой недели
func Last_days(n int) []time.Time {
	if n == 0 {
		return []time.Time{time.Now()}
	}
	res := make([]time.Time, 0)
	t := time.Now().AddDate(0, 0, -n)
	for i := 0; i < n; i++ {
		res = append(res, t)
		t = t.Add(24 * time.Hour)
	}
	return res
}

// список дней будущей недели
func Next_Ned(t time.Time, n int) (res []time.Time) {
	d := int(t.Weekday())
	dn := t
	if d == 0 {
		dn = t.AddDate(0, 0, (n-1)*7+1)
	} else {
		dn = t.AddDate(0, 0, (n-1)*7+(8-d))
	}
	res = append(res, dn)
	res = append(res, dn.AddDate(0, 0, 1))
	res = append(res, dn.AddDate(0, 0, 2))
	res = append(res, dn.AddDate(0, 0, 3))
	res = append(res, dn.AddDate(0, 0, 4))
	res = append(res, dn.AddDate(0, 0, 5))
	res = append(res, dn.AddDate(0, 0, 6))
	return
}

func SliseDat(a []time.Time) (res []string) {
	for _, i := range a {
		res = append(res, i.Format(TNS))
	}
	return
}

// список дней прошлой недели
func Last_day(n int) time.Time {
	return Last_days(n)[0]
}

// список дней прошлой недели
func Day(n string) time.Time {
	t, _ := time.Parse("02.01.2006", n)
	return t
}

// список дней прошлой недели
func DayList(n string) []time.Time {
	t, _ := time.Parse("02.01.2006", n)
	return []time.Time{t}
}

// список воскресений прошлого месяца
func Last_mounth_voskr(n int) (res []time.Time) {
	for _, i := range Last_mounth_one(n) {
		if i.Weekday() == 0 {
			res = append(res, i)
		}
	}
	return
}

// список дней прошлых месяцев
func Last_mounth(n int) []time.Time {
	if n == 0 {
		return []time.Time{time.Now()}
	}
	res := make([]time.Time, 0)
	t := time.Now()
	m := t.Month()
	if int(t.Day()) != 1 {
		for int(t.Day()) != 1 {
			t = t.AddDate(0, 0, -1)
		}
	}
	t = t.AddDate(0, -n, 0)

	for {
		res = append(res, t)
		t = t.Add(24 * time.Hour)
		if t.Day() == 1 && t.Month() == m {
			break
		}
	}
	return res
}

// список дней прошлых месяцев
func Last_mounth_one(n int) []time.Time {
	t := time.Now().AddDate(0, -n, 0)
	y := t.Year()
	m := int(t.Month())

	return Mounth(m, y)
}

func List_mounth(a []time.Time) (res [][2]int) {
	sort.Slice(a, func(i, j int) (less bool) {
		return a[i].Before(a[j])
	})
	for _, i := range a {
		n := int(i.Month())
		fl := 0
		for _, j := range res {
			if j[0] == n && j[1] == i.Year() {
				fl = 1
				break
			}
		}
		if fl == 0 {
			res = append(res, [2]int{n, i.Year()})
		}
	}
	return
}

// список дней прошлых недель
func Last_ned(n int) []time.Time {
	if n == 0 {
		return []time.Time{time.Now()}
	}
	res := make([]time.Time, 0)
	t := time.Now()
	if int(t.Weekday()) != 1 {
		for int(t.Weekday()) != 1 {
			t = t.AddDate(0, 0, -1)
		}
	}
	t = t.AddDate(0, 0, -7*n)

	for i := 0; i < n*7; i++ {
		res = append(res, t)
		t = t.Add(24 * time.Hour)
	}
	return res
}

func List_ned(a []time.Time) (res []int) {
	sort.Slice(a, func(i, j int) (less bool) {
		return a[i].Before(a[j])
	})
	for _, i := range a {
		_, n := i.ISOWeek()
		if !comp_int(n, res) {
			res = append(res, n)
		}
	}
	return
}

// список дней месяца
func Day_mounth(m, y int) []time.Time {
	if m == 0 || y == 0 {
		return []time.Time{time.Now()}
	}
	res := make([]time.Time, 0)
	t := time.Date(y, time.Month(m), 1, 2, 2, 2, 0, time.UTC)
	for i := 0; i < 35; i++ {
		mm := int(t.Month())
		yy := t.Year()
		if mm == m && yy == y {
			res = append(res, t)
		}
		t = t.Add(24 * time.Hour)
	}
	return res
}

// список дней месяца
func Day_mounth_String(m, y int) []string {
	if m == 0 || y == 0 {
		return []string{""}
	}
	res := make([]string, 0)
	t := time.Date(y, time.Month(m), 1, 2, 2, 2, 0, time.UTC)
	for i := 0; i < 35; i++ {
		mm := int(t.Month())
		yy := t.Year()
		if mm == m && yy == y {
			res = append(res, t.Format(TNS))
		}
		t = t.Add(24 * time.Hour)
	}
	return res
}

// список дней всех!!! недель в месяце
func Day_ned_mounth(m, y int) []time.Time {
	if m == 0 || y == 0 {
		return []time.Time{time.Now()}
	}
	res := make([]time.Time, 0)

	mes := Day_mounth(m, y)
	ned := make([]int, 0)

	for _, i := range mes {
		_, n := i.ISOWeek()
		if !comp_int(n, ned) {
			ned = append(ned, n)
		}
	}

	for _, i := range ned {
		res = append(res, Nom_ned_days(i, y)...)
	}

	return res
}

// список дней прошлой недели
func Last_days_string(n int) (res []string) {
	t := time.Now().AddDate(0, 0, -n)
	for i := 0; i < n; i++ {
		res = append(res, t.Format(TNS))
		t = t.Add(24 * time.Hour)
	}
	return res
}

// список дней текущей недели
func Teck_days() []time.Time {
	res := make([]time.Time, 0)
	wd := int(time.Now().Weekday())
	if wd == 0 {
		wd = 7
	}
	t := time.Now().AddDate(0, 0, -(wd - 1))
	for i := 0; i < 7; i++ {
		res = append(res, t)
		t = t.Add(24 * time.Hour)
	}
	return res
}

// список дней по номеру месяца
// в текущем году
func Mounth_days(n int) []time.Time {
	res := make([]time.Time, 0)
	tn := time.Now()
	mounth := int(tn.Month())
	day := tn.Day()

	t := time.Now().AddDate(0, n-mounth, -day+1)
	res = append(res, t)
	for i := 0; i < 32; i++ {
		t = t.Add(24 * time.Hour)
		if int(t.Month()) == n {
			res = append(res, t)
		}
	}
	return res
}

// список дней по номеру недели
func Nom_ned_days(n int, y int) []time.Time {
	t := time.Now()
	res := make([]time.Time, 0)
	for i := 0; i < 1000; i++ {
		ye, nd := t.ISOWeek()
		if ye == y && nd == n {
			t = t.Add(-24 * 6 * time.Hour)
			for j := 0; j < 7; j++ {
				res = append(res, t)
				t = t.Add(24 * time.Hour)
			}
			return res
		}
		t = t.Add(-24 * time.Hour)
	}
	return res
}

// проверяет вхождение int64 в слайс
func comp_int(s int, ss []int) bool {
	for _, i := range ss {
		if i == s {
			return true
		}
	}
	return false
}
