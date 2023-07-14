package citys

import (
	"fmt"
	"osnova/exel"
	"osnova/trs"
	"strings"
	"time"
)

var CityList ListCity

func RunListCity() {
	CityList = Create_Citys_List()
}

func Create_Citys_List() (res []City) {

	list, err := exel.ReadGoogleSheets("10JOrpQU63VyK6VL7EEdj1iiyk6sNzbDaGzW7Fqg2eMs", "'Citys'")
	if err != nil {
		return
	}

	for _, i := range list {
		if i[2] != "YES" {
			continue
		}
		cit := City{}
		cit.Create(i)
		res = append(res, cit)
	}

	return res
}

type B_q_get struct {
	Name_dataset string
	Name_table   string
}

type City struct {
	Status          []string
	Imployment      B_q_get
	Bolt_trip       B_q_get
	Uclon_trip      B_q_get
	Uclon_cash      B_q_get
	Uclon_today     B_q_get
	Grafik          B_q_get
	Vod             B_q_get
	Kassa           B_q_get
	Save_kassa_BGQ  bool
	Ident           string
	Id_BGQ          string
	Name            string
	Name_vod        string
	Name_online     string
	Name_grafik     string
	Sity_report     string
	Online_diapazon string
	Nov_vod         string
	Mapon_bolt      string
	Mapon_uclon     string
	Mapon_bolt_id   string
	Mapon_uclon_id  string
	Id              int
	Procent_bolt    int
	Procent_uclon   int
	Procent_vozvrat int
	Procent_vod     int
	Cena_gas        int
	Rasxod_avto     int
	ProfitV2        bool
	TygnZvtBGQ      bool
	ViewKassaBGQ    bool
	Flag            bool
	CRM             bool
	CRM_Name        string
	CRM_DATE        time.Time
}

type ListCity []City

func (a *City) IsCRM() bool {
	return a.CRM_DATE.Before(time.Now())
}

func (a *ListCity) City(n int) (res ListCity) {
	return ListCity{(*a)[n]}
}

func (a *ListCity) CityId(s string) City {
	for _, i := range *a {
		if i.Id_BGQ == s {
			return i
		}
	}
	return City{}
}

func (a *ListCity) CityName(s string) City {
	for _, i := range *a {
		if i.Name == s {
			return i
		}
	}
	return City{}
}

// Печать элементаCity
func (a *City) Print() {
	fmt.Println()
	fmt.Println("City                :")
	fmt.Println("Status              :", a.Status)
	fmt.Println("Imployment          :", a.Imployment)
	fmt.Println("Bolt_trip           :", a.Bolt_trip)
	fmt.Println("Uclon_trip          :", a.Uclon_trip)
	fmt.Println("Uclon_cash          :", a.Uclon_cash)
	fmt.Println("Uclon_today         :", a.Uclon_today)
	fmt.Println("Grafik              :", a.Grafik)
	fmt.Println("Vod                 :", a.Vod)
	fmt.Println("Kassa               :", a.Kassa)
	fmt.Println("Save_kassa_BGQ      :", a.Save_kassa_BGQ)
	fmt.Println("Ident               :", a.Ident)
	fmt.Println("Id_BGQ              :", a.Id_BGQ)
	fmt.Println("Name                :", a.Name)
	fmt.Println("Name_vod            :", a.Name_vod)
	fmt.Println("Name_online         :", a.Name_online)
	fmt.Println("Name_grafik         :", a.Name_grafik)
	fmt.Println("Sity_report         :", a.Sity_report)
	fmt.Println("Online_diapazon     :", a.Online_diapazon)
	fmt.Println("Nov_vod             :", a.Nov_vod)
	fmt.Println("Mapon_bolt          :", a.Mapon_bolt)
	fmt.Println("Mapon_uclon         :", a.Mapon_uclon)
	fmt.Println("Mapon_bolt_id       :", a.Mapon_bolt_id)
	fmt.Println("Mapon_uclon_id      :", a.Mapon_uclon_id)
	fmt.Println("Id                  :", a.Id)
	fmt.Println("Procent_bolt        :", a.Procent_bolt)
	fmt.Println("Procent_uclon       :", a.Procent_uclon)
	fmt.Println("Procent_vozvrat     :", a.Procent_vozvrat)
	fmt.Println("Procent_vod         :", a.Procent_vod)
	fmt.Println("Cena_gas            :", a.Cena_gas)
	fmt.Println("Rasxod_avto         :", a.Rasxod_avto)
	fmt.Println("ProfitV2            :", a.ProfitV2)
	fmt.Println("TygnZvtBGQ          :", a.TygnZvtBGQ)
	fmt.Println("ViewKassaBGQ        :", a.ViewKassaBGQ)
	fmt.Println("Flag                :", a.Flag)
	fmt.Println("CRM                 :", a.CRM)
	fmt.Println("CRM_Name            :", a.CRM_Name)
	fmt.Println("CRM_DATE            :", a.CRM_DATE)
}

func (cit *City) Create(r []string) {
	cit.Name = r[0]
	cit.Id_BGQ = r[1]

	cit.Imployment.Name_dataset = r[3]
	cit.Imployment.Name_table = r[4]

	cit.Bolt_trip.Name_dataset = r[5]
	cit.Bolt_trip.Name_table = r[6]

	cit.Uclon_trip.Name_dataset = r[7]
	cit.Uclon_trip.Name_table = r[8]

	cit.Uclon_cash.Name_dataset = r[9]
	cit.Uclon_cash.Name_table = r[10]

	cit.Vod.Name_dataset = r[11]
	cit.Vod.Name_table = r[12]

	cit.Kassa.Name_dataset = r[13]
	cit.Kassa.Name_table = r[14]

	if r[15] == "TRUE" {
		cit.Save_kassa_BGQ = true
	}
	cit.Ident = r[16]
	cit.Id_BGQ = r[17]
	cit.Name = r[18]
	cit.Name_vod = r[19]
	cit.Name_online = r[20]
	cit.Name_grafik = r[21]
	cit.Sity_report = r[22]
	cit.Online_diapazon = r[23]
	cit.Nov_vod = r[26]
	cit.Mapon_bolt = r[27]
	cit.Mapon_uclon = r[28]
	cit.Mapon_bolt_id = r[29]
	cit.Mapon_uclon_id = r[30]

	cit.Id = trs.String_to_int(r[39])
	cit.Procent_bolt = trs.String_to_int(r[40])
	cit.Procent_uclon = trs.String_to_int(r[41])
	cit.Procent_vozvrat = trs.String_to_int(r[42])
	cit.Procent_vod = trs.String_to_int(r[43])
	cit.Cena_gas = trs.String_to_int(r[44])
	cit.Rasxod_avto = trs.String_to_int(r[45])

	cit.Grafik.Name_dataset = r[49]
	cit.Grafik.Name_table = r[50]

	cit.Uclon_today.Name_dataset = r[51]
	cit.Uclon_today.Name_table = r[52]
	if r[53] == "YES" {
		cit.ProfitV2 = true
	}
	if r[54] == "Y" {
		cit.TygnZvtBGQ = true
	}
	if r[55] == "Y" {
		cit.ViewKassaBGQ = true
	}
	if r[2] == "YES" {
		cit.Flag = true
	}
	if r[56] == "Y" {
		cit.CRM = true
	}
	cit.CRM_Name = strings.ReplaceAll(r[57], "\"", "")
	if r[58] != "" {
		cit.CRM_DATE, _ = time.Parse("02.01.2006", r[58])
	}
}

var Name_CL = map[string]string{
	"Київ":                "Kyiv",
	"Kyiv":                "Kyiv",
	"Дніпро":              "Dnipro",
	"Dnipro":              "Dnipro",
	"Запоріжжя":           "Zaporizhia",
	"Zaporizhia":          "Zaporizhia",
	"Одеса":               "Odesa",
	"Odesa":               "Odesa",
	"Івано Франківськ":    "Ivano",
	"Ivano":               "Ivano",
	"Суми":                "Sumy",
	"Sumy":                "Sumy",
	"Полтава":             "Poltava",
	"Poltava":             "Poltava",
	"Луцьк":               "Lutsk",
	"Lutsk":               "Lutsk",
	"Ужгород":             "Uzhhorod",
	"Uzhhorod":            "Uzhhorod",
	"Черкаси":             "Cherkasy",
	"Cherkasy":            "Cherkasy",
	"Чернігів":            "Chernihiv",
	"Chernihiv":           "Chernihiv",
	"Чернівці":            "Chernivtsi",
	"Chernivtsi":          "Chernivtsi",
	"Львів":               "Lviv",
	"Lviv":                "Lviv",
	"Кривий Ріг":          "Kryvyi",
	"Kryvyi":              "Kryvyi",
	"Рівне":               "Rivne",
	"Rivne":               "Rivne",
	"Вінниця":             "Vinnytsia",
	"Vinnytsia":           "Vinnytsia",
	"Житомир":             "Zhytomyr",
	"Zhytomyr":            "Zhytomyr",
	"Хмельницький":        "Khmelnytskyi",
	"Khmelnytskyi":        "Khmelnytskyi",
	"Тернопіль":           "Ternopil",
	"Ternopil":            "Ternopil",
	"Мукачево":            "Mukachevo",
	"Mukachevo":           "Mukachevo",
	"Камянец":             "Kamianets_Podilskyi",
	"Kamianets_Podilskyi": "Kamianets_Podilskyi",
	"Харків":              "Kharkiv",
	"Kharkiv":             "Kharkiv",
	"Херсон":              "Kherson",
	"Kherson":             "Kherson",
	"Варшава":             "Warsaw",
	"Warsaw":              "Warsaw",
}
var Name_LC = map[string]string{
	"Київ":                "Київ",
	"Kyiv":                "Київ",
	"Дніпро":              "Дніпро",
	"Dnipro":              "Дніпро",
	"Запоріжжя":           "Запоріжжя",
	"Zaporizhia":          "Запоріжжя",
	"Одеса":               "Одеса",
	"Odesa":               "Одеса",
	"Івано Франківськ":    "Івано Франківськ",
	"Ivano":               "Івано Франківськ",
	"Суми":                "Суми",
	"Sumy":                "Суми",
	"Полтава":             "Полтава",
	"Poltava":             "Полтава",
	"Луцьк":               "Луцьк",
	"Lutsk":               "Луцьк",
	"Ужгород":             "Ужгород",
	"Uzhhorod":            "Ужгород",
	"Черкаси":             "Черкаси",
	"Cherkasy":            "Черкаси",
	"Чернігів":            "Чернігів",
	"Chernihiv":           "Чернігів",
	"Чернівці":            "Чернівці",
	"Chernivtsi":          "Чернівці",
	"Львів":               "Львів",
	"Lviv":                "Львів",
	"Кривий Ріг":          "Кривий Ріг",
	"Kryvyi":              "Кривий Ріг",
	"Рівне":               "Рівне",
	"Rivne":               "Рівне",
	"Вінниця":             "Вінниця",
	"Vinnytsia":           "Вінниця",
	"Житомир":             "Житомир",
	"Zhytomyr":            "Житомир",
	"Хмельницький":        "Хмельницький",
	"Khmelnytskyi":        "Хмельницький",
	"Тернопіль":           "Тернопіль",
	"Ternopil":            "Тернопіль",
	"Мукачево":            "Мукачево",
	"Mukachevo":           "Мукачево",
	"Камянец":             "Камянец",
	"Kamianets_Podilskyi": "Камянец",
	"Харків":              "Харків",
	"Kharkiv":             "Харків",
	"Херсон":              "Херсон",
	"Kherson":             "Херсон",
	"Варшава":             "Варшава",
	"Warsaw":              "Варшава",
}

// псевдонимы городов
var Psevdonimy = map[string]string{
	"Сумы":                "Суми",
	"Sumy":                "Суми",
	"Суми":                "Суми",
	"Львов":               "Львів",
	"Lviv":                "Львів",
	"Львів":               "Львів",
	"Ивано  Франковск":    "Івано Франківськ",
	"Frankivsk":           "Івано Франківськ",
	"Ivano":               "Івано Франківськ",
	"Ivano-Frankivsk":     "Івано Франківськ",
	"Івано Франківськ":    "Івано Франківськ",
	"Івано-Франківськ":    "Івано Франківськ",
	"Ivano - Frankivsk":   "Івано Франківськ",
	"Винница":             "Вінниця",
	"Vinnytsia":           "Вінниця",
	"Вінниця":             "Вінниця",
	"Хмельницкий":         "Хмельницький",
	"Khmelnytskyi":        "Хмельницький",
	"Хмельницький":        "Хмельницький",
	"Черновцы":            "Чернівці",
	"Chernivtsi":          "Чернівці",
	"Чернівці":            "Чернівці",
	"Луцк":                "Луцьк",
	"Lutsk":               "Луцьк",
	"Луцьк":               "Луцьк",
	"Ровно":               "Рівне",
	"Рівне":               "Рівне",
	"Rivne":               "Рівне",
	"Ужгород":             "Ужгород",
	"Uzhhorod":            "Ужгород",
	"Житомир":             "Житомир",
	"Zhytomyr":            "Житомир",
	"Днепр":               "Дніпро",
	"Dnipro":              "Дніпро",
	"Dnipropetrovsk":      "Дніпро",
	"Дніпро":              "Дніпро",
	"Черкассы":            "Черкаси",
	"Черкаси":             "Черкаси",
	"Cherkasy":            "Черкаси",
	"Полтава":             "Полтава",
	"Poltava":             "Полтава",
	"Одесса":              "Одеса",
	"Odesa":               "Одеса",
	"Одеса":               "Одеса",
	"Кривий Ріг":          "Кривий Ріг",
	"Кривой Рог":          "Кривий Ріг",
	"Kr":                  "Кривий Ріг",
	"Kryvyi Rih":          "Кривий Ріг",
	"Kryvyi":              "Кривий Ріг",
	"Запорожье":           "Запоріжжя",
	"Zaporizhia":          "Запоріжжя",
	"Zaporizhzhia":        "Запоріжжя",
	"Запоріжжя":           "Запоріжжя",
	"Киев":                "Київ",
	"Kyiv":                "Київ",
	"Київ":                "Київ",
	"Тернополь":           "Тернопіль",
	"Ternopil":            "Тернопіль",
	"Тернопіль":           "Тернопіль",
	"Чернигов":            "Чернігів",
	"Chernihiv":           "Чернігів",
	"Чернігів":            "Чернігів",
	"Харків":              "Харків",
	"Харьков":             "Харків",
	"Kharkiv":             "Харків",
	"Херсон":              "Херсон",
	"Kherson":             "Херсон",
	"Варшава":             "Варшава",
	"Мукачево":            "Мукачево",
	"Камянец":             "Камянец",
	"Kamianets_Podilskyi": "Камянец",
	"Kamianets-Podilskyi": "Камянец",
	"Кам'янець-Подільський": "Камянец",
	"Kam Podil": "Камянец",
	"Warshawa":  "Варшава",
	"Mukachevo": "Мукачево",
	"Total":     "Total",
	"TOTAL":     "TOTAL",
}

var City_Name = map[string]string{
	"Київ":             "Kyiv",
	"Дніпро":           "Dnipro",
	"Запоріжжя":        "Zaporizhia",
	"Одеса":            "Odesa",
	"Івано Франківськ": "Ivano",
	"Суми":             "Sumy",
	"Полтава":          "Poltava",
	"Луцьк":            "Lutsk",
	"Ужгород":          "Uzhhorod",
	"Черкаси":          "Cherkasy",
	"Чернігів":         "Chernihiv",
	"Чернівці":         "Chernivtsi",
	"Львів":            "Lviv",
	"Кривий Ріг":       "Kryvyi",
	"Рівне":            "Rivne",
	"Вінниця":          "Vinnytsia",
	"Житомир":          "Zhytomyr",
	"Хмельницький":     "Khmelnytskyi",
	"Тернопіль":        "Ternopil",
	"Мукачево":         "Mukachevo",
	"Камянец":          "Kamianets_Podilskyi",
	"Харків":           "Kharkiv",
	"Херсон":           "Kherson",
	"Варшава":          "Warsaw",
}

var City_Name_Reverse = map[string]string{
	"Kyiv":                "Київ",
	"Dnipro":              "Дніпро",
	"Zaporizhia":          "Запоріжжя",
	"Odesa":               "Одеса",
	"Ivano":               "Івано Франківськ",
	"Sumy":                "Суми",
	"Poltava":             "Полтава",
	"Lutsk":               "Луцьк",
	"Uzhhorod":            "Ужгород",
	"Cherkasy":            "Черкаси",
	"Chernihiv":           "Чернігів",
	"Chernivtsi":          "Чернівці",
	"Lviv":                "Львів",
	"Kryvyi":              "Кривий Ріг",
	"Rivne":               "Рівне",
	"Vinnytsia":           "Вінниця",
	"Zhytomyr":            "Житомир",
	"Khmelnytskyi":        "Хмельницький",
	"Ternopil":            "Тернопіль",
	"Mukachevo":           "Мукачево",
	"Kamianets_Podilskyi": "Камянец",
	"Kharkiv":             "Харків",
	"Kherson":             "Херсон",
	"Warsaw":              "Варшава",
}
