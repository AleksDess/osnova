package main

import (
	"fmt"
	"log"
	"os"
	"osnova/citys"
	"osnova/crm"
	"osnova/logger"
	"osnova/postgree"
	"osnova/trs"
	"time"

	"github.com/joho/godotenv"
)

var problem_avto string

const speed_message = "Отчет подготовлен: Оповещения Компания"
const hourly_report = "Отчет подготовлен: Пробег по часам"
const behavior_report = "Отчет подготовлен: Поведение вождения"

var url = ""
var urlday = ""

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	tgToken := os.Getenv("TGTOKEN")
	crm.SshHost = os.Getenv("SshHost")
	crm.SshPort = trs.String_to_int(os.Getenv("SshPort"))
	crm.SshUser = os.Getenv("SshUser")
	crm.DbUser = os.Getenv("DbUser")
	crm.DbPass = os.Getenv("DbPass")
	crm.DbHost = os.Getenv("DbHost")
	crm.DbName = os.Getenv("DbName")
	crm.PrivateKeyPath = os.Getenv("PrivateKeyPath")
	postgree.PostHost = os.Getenv("PostHost")
	postgree.PostPort = os.Getenv("PostPort")
	postgree.PostPass = os.Getenv("PostPass")
	url = os.Getenv("URL")
	urlday = os.Getenv("URLDAY")

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
