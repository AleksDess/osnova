package trs

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

func Generate_Uuid() string {
	return uuid.New().String()
}

// Замена кирилических символов на латинские
func Trs(a string) string {

	ch := []rune(a)

	if len(ch) < 8 {
		return "nocorect"
	}

	rs := ""
	res := ""

	for i := 0; i < 8; i++ {
		rs = string(ch[i])
		switch string(ch[i]) {
		case "А":
			rs = "A"
		case "a":
			rs = "A"
		case "а":
			rs = "A"
		case "В":
			rs = "B"
		case "в":
			rs = "B"
		case "С":
			rs = "C"
		case "с":
			rs = "C"
		case "О":
			rs = "O"
		case "о":
			rs = "O"
		case "І":
			rs = "I"
		case "і":
			rs = "I"
		case "|":
			rs = "I"
		case "Р":
			rs = "P"
		case "р":
			rs = "P"
		case "Т":
			rs = "T"
		case "М":
			rs = "M"
		case "м":
			rs = "M"
		case "Н":
			rs = "H"
		case "н":
			rs = "H"
		case "К":
			rs = "K"
		case "к":
			rs = "K"
		case "Е":
			rs = "E"
		case "е":
			rs = "E"
		case "Х":
			rs = "X"
		case "х":
			rs = "X"
		}
		res = res + rs
	}

	return res
}

// проверка значения в графике
// на имя водителя
func If_Name_Voditel(s string) bool {
	return strings.Contains(s, " ")
}

func OnlyNumber(s string) string {
	b := []byte(s)
	r := make([]byte, 0)
	for _, i := range b {
		if i < 47 && i < 58 {
			r = append(r, i)
		}
	}
	return string(r)
}

// Проверка номера авто на корректность
func Check_nomer(a string) string {

	b := []byte(a)
	if len(a) < 8 {
		return "Слишком короткий номер"
	}
	if b[0] == 32 {
		return "Номер начинается с пробела"
	}
	if b[0] < 65 || b[0] > 91 {
		return "Первый символ номера некорректен"
	}
	if b[1] < 65 || b[1] > 91 {
		return "Второй символ номера некорректен"
	}
	if b[6] < 65 || b[6] > 91 {
		return "Седьмой символ номера некорректен"
	}
	if b[7] < 65 || b[7] > 91 {
		return "Восьмой символ номера некорректен"
	}
	if b[2] < 48 || b[2] > 57 {
		return "Четвертый символ нмера некорректен"
	}
	if b[2] < 48 || b[2] > 57 {
		return "Пятый символ номер некорректен"
	}
	if b[2] < 48 || b[2] > 57 {
		return "Шестой символ номера некорректен"
	}

	return "OK"
}

func Parse_time_google_sheets(s string) (t time.Time, err error) {

	t, err = time.Parse("1/2/2006 15:04:05", s)
	if err == nil {
		return
	}
	t, err = time.Parse("01/02/2006 15:04:05", s)
	if err == nil {
		return
	}
	t, err = time.Parse("02.01.2006 15:04:05", s)
	if err == nil {
		return
	}
	t, err = time.Parse("2006-01-02 15:04:05", s)
	if err == nil {
		return
	}
	t, err = time.Parse("2006-01-02", s)
	if err == nil {
		return
	}
	t, err = time.Parse("1/2/2006", s)
	if err == nil {
		return
	}
	t, err = time.Parse("01/02/2006", s)
	if err == nil {
		return
	}
	t, err = time.Parse("02.01.2006", s)
	if err == nil {
		return
	}
	return
}

func Parse_num_google_sheets(s string) int {
	if strings.Contains(s, ",") && strings.Contains(s, ".") {
		s = strings.Replace(s, ".00", "", 1)
		s = strings.Replace(s, ",", "", 1)
		z, err := strconv.ParseFloat(s, 64)
		if err == nil {
			return int(z)
		} else {
			fmt.Println(s)
		}
	} else if strings.Contains(s, ",") {
		s = strings.Replace(s, ",", "", 1)
		z, err := strconv.ParseFloat(s, 64)
		if err == nil {
			return int(z)
		} else {
			fmt.Println(s)
		}
	} else if strings.Contains(s, ".") {
		z, err := strconv.ParseFloat(s, 64)
		if err == nil {
			return int(z)
		} else {
			fmt.Println(s)
		}
	} else {
		z, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println(s)
		} else {
			return z
		}
	}
	return 0
}

// попытка уйти от двойных пробелов в именах
func Clear_32(s string) string {
	s = strings.ReplaceAll(s, "     ", " ")
	s = strings.ReplaceAll(s, "    ", " ")
	s = strings.ReplaceAll(s, "   ", " ")
	s = strings.ReplaceAll(s, "  ", " ")
	s = strings.TrimSpace(s)
	return s
}

// попытка уйти от двойных пробелов в именах
// uber uklon
func Clear_name_driver(s string) string {
	s = strings.ReplaceAll(s, "     ", " ")
	s = strings.ReplaceAll(s, "    ", " ")
	s = strings.ReplaceAll(s, "   ", " ")
	s = strings.ReplaceAll(s, "  ", " ")
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, "(UK) ", "")
	s = strings.ReplaceAll(s, "(UB) ", "")
	s = strings.ReplaceAll(s, "(M) ", "")
	return s
}

func Reverse_name_driver(s string) string {
	r := strings.Split(s, " ")
	//fmt.Println(s, r, len(r))
	if len(r) == 2 {
		return r[1] + " " + r[0]
	} else if len(r) == 2 {
		fmt.Println(s, r[1]+" "+r[0]+" "+r[2])
		return r[1] + " " + r[0] + " " + r[2]
	} else {
		return s
	}
}

func Float_to_string(a float64) string {
	return strings.Replace(fmt.Sprintf("%.1f", a), ".", ",", 1)
}
func Float_to_string_2(a float64) string {
	return strings.Replace(fmt.Sprintf("%.2f", a), ".", ",", 1)
}
func Float_to_string_4(a float64) string {
	return strings.Replace(fmt.Sprintf("%.4f", a), ".", ",", 1)
}

func String_to_int(s string) int {
	s = strings.Replace(s, ",", ".", 1)
	r, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	} else {
		return int(r)
	}
}

func String_to_int64(s string) int64 {
	s = strings.Replace(s, ",", ".", 1)
	r, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}
	return int64(int(r))
}

func String_TGID_to_int64(s string) int64 {
	r, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}
	return r
}

func String_to_float(s string) float64 {
	s = strings.Replace(s, ",", ".", 1)
	r, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	} else {
		return r
	}
}

func String_to_float_Sheets(s string) float64 {
	if strings.Count(s, ".") == 2 {
		s = strings.Replace(s, ".", "", 1)
	}
	s = strings.Replace(s, ",", "", 1)
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ReplaceAll(s, string(rune(160)), "")

	r, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Println(s, err)
		return 0
	} else {
		return r
	}
}

func String_to_float_SheetsA(s string) float64 {
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ReplaceAll(s, ",", ".")
	s = strings.ReplaceAll(s, string(rune(160)), "")

	r, err := strconv.ParseFloat(s, 64)
	if err != nil {
		//fmt.Println(s, err)
		return 0
	} else {
		return r
	}
}

func Parse_ID_BOT(s string) int64 {
	r, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return int64(0)
	}
	return int64(r)
}

func Compaire_int64(a int64, b []int64) bool {
	for _, i := range b {
		if a == i {
			return true
		}
	}
	return false
}

func Compaire_String(a string, b []string) bool {
	for _, i := range b {
		if a == i {
			return true
		}
	}
	return false
}

var Status = map[string]string{
	"ДТП":          "ДТП",
	"NW-ДТП":       "ДТП",
	"Ремонт":       "Ремонт",
	"NW-Ремонт":    "Ремонт",
	"ТО-Сервіс":    "Сервіс",
	"NW-ТО-Сервіс": "Сервіс",
	"Выходной":     "Вихідний",
	"Вихідний":     "Вихідний",
	"ВЫХОДНОЙ":     "Вихідний",
	"ЕВАКУАЦІЯ - Відправка": "ЕВАКУАЦІЯ - Відправка",
	"ЕВАКУАЦІЯ-Відправка":   "ЕВАКУАЦІЯ - Відправка",
	"ЕВАКУАЦІЯ - Прийом":    "ЕВАКУАЦІЯ - Прийом",
	"ЕВАКУАЦІЯ-Прийом":      "ЕВАКУАЦІЯ - Прийом",
	"ВИКОНАНО":              "ВИКОНАНО",
	"Штрафмайданчик":        "Штрафмайданчик",
	"Штрафмайданчик ":       "Штрафмайданчик",
	"NW-Штрафплощадка":      "Штрафмайданчик",
	"Документи":             "Документи",
	"NW-Документы":          "Документи",
	"Оренда":                "Оренда",
	"NW-Аренда":             "Оренда",
	"ПРОДАЖА":               "Оренда",
	"Викуп":                 "Оренда",
	"Угон":                  "Угон",
	"NW-Угон":               "Угон",
	"Ключі":                 "Ключі",
	"NW-Ключи":              "Ключі",
	"Волонтер":              "Волонтер",
	"NW-Волонтер":           "Волонтер",
	"Списання":              "Списання",
	"СПИСАННЯ":              "Списання",
	"NW-Списання":           "Списання",
	"Лікарняний":            "Лікарняний",
	"Больничный":            "Лікарняний",
	"NW-Лікарняний":         "Лікарняний",
}

// проверяет вхождение time.Time в слайс
func Comp_time(s time.Time, ss []time.Time) bool {
	for _, i := range ss {
		if i.Equal(s) {
			return true
		}
	}
	return false
}

func Check_Usi_Vodii(s string) bool {
	if s == "Усі водії" || s == "Всі водії" || s == " Усі водії" || s == " Всі водії" {
		return true
	}
	return false
}

// проверяет вхождение string в слайс
func Comp_string(s string, ss []string) bool {
	for _, i := range ss {
		if i == s {
			return true
		}
	}
	return false
}

// проверяет вхождение string в слайс
func Comp_string_Append(s string, ss *[]string) bool {
	for _, i := range *ss {
		if i == s {
			return true
		}
	}
	*ss = append(*ss, s)
	return false
}

// проверяет вхождение string в слайс
func Comp_string_2(s string, ss [][2]string) bool {
	for _, i := range ss {
		if i[0] == s {
			return true
		}
	}
	return false
}

// проверяет вхождение string в слайс
func Comp_slise_string(s, ss []string) bool {
	if len(s) != len(ss) {
		return false
	}
	for i := 0; i < len(s); i++ {
		if s[i] != ss[i] {
			return false
		}
	}
	return true
}

// просчет количества вхождений
// элемента из а в b
func Comp_strings_contains(a, b []string) (r int) {
	for _, i := range a {
		for _, j := range b {
			if i == j {
				r++
			}
		}
	}
	return
}

// проверяет вхождение int64 в слайс
func Comp_int64(s int64, ss []int64) bool {
	for _, i := range ss {
		if i == s {
			return true
		}
	}
	return false
}

// проверяет вхождение int64 в слайс
func Comp_int(s int, ss []int) bool {
	for _, i := range ss {
		if i == s {
			return true
		}
	}
	return false
}

// сравнивает имена
// для двух аккаунтов
func Comp_Name(ss1, ss2 string) bool {

	clear := func(s string) string {
		b := []byte(s)
		r := make([]byte, 0)
		for _, i := range b {
			if i == 32 || i == 40 || i == 41 || i == 85 || i == 75 {
				continue
			}
			r = append(r, i)
		}
		return string(r)
	}

	s1 := clear(ss1)
	s2 := clear(ss2)

	run1 := []rune(s1)
	run2 := []rune(s2)

	if len(run1) != len(run2) {
		return false
	}

	n := 0
	for nn, i := range run1 {
		if i != run2[nn] {
			n++
		}
	}

	if n < 2 {
		return true
	} else {
		return false
	}

}
