package main

import (
	"fmt"
	"log"
	"os"
	"osnova/citys"
	"osnova/crm"
	"osnova/logger"
	"osnova/postgree"
	"time"

	"github.com/joho/godotenv"
)

var problem_avto string

const speed_message = "Отчет подготовлен: Оповещения Компания"
const hourly_report = "Отчет подготовлен: Пробег по часам"
const behavior_report = "Отчет подготовлен: Поведение вождения"

var bot_token = ""
var bot_token_osnova = "2096779449:AAFfd9HsY0OFIQ9JAxWssXDVzs4Tp_Rn4H4"
var bot_token_test = "5419530014:AAGKNw9CcQa2jxKQiou4wPO7Eaq_g-nUfv0"

var Fesenko = "Не пройшло перевірку у контрольного ДМ"
var Fesenko_1 = "Часткова відміна штрафу"

var url = "http://92.119.231.174:8081/driver/"
var urlday = "http://92.119.231.174:8081/day/"

// Файлы для ехел обработки
// "mail/trip.xlsx"
// "mail/opov.xlsx"

func init() {

}

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	port := os.Getenv("PORT")
	tgToken := os.Getenv("TELEGRAMBOTTOKEN")

	log.Println(port)
	log.Println(tgToken)
	// без него не Run ---  аемся
	if tgToken == "" {
		logger.ErrorLog.Println("-telegrambottoken is required")
		os.Exit(1)
	}

	citys.RunListCity()
	logger.RunLogger()

	err = postgree.RunDB()
	if err != nil {
		logger.ErrorLog.Println(err)
		os.Exit(1)
	}

	err = crm.RunDB()
	if err != nil {
		logger.ErrorLog.Println(err)
		os.Exit(1)
	}

	CreateListCitys()
	// go create_sitys_List()
}

func PrintListSity() {
	for n, i := range citys.CityList {
		fmt.Println(n, i.Name, i.Id_BGQ)
	}
}

func UpgradeListCitys() {
	for {
		time.After(time.Hour)
		citys.CityList = citys.Create_Citys_List()
	}
}

func CreateListCitys() {
	citys.CityList = citys.Create_Citys_List()
	go UpgradeListCitys()
}
