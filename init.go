package main

import (
	"fmt"
	"log"
	"os"
	"osnova/citys"
	"osnova/crm"
	"osnova/imap"
	"osnova/logger"
	"osnova/postgree"
	"osnova/trs"
	"time"

	"github.com/joho/godotenv"
)

var url = ""
var urlday = ""
var tgToken = ""

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// load data .env
	OsGetEnv()

	citys.RunListCity()
	logger.RunLogger()

	// без него не Run ---  аемся
	if tgToken == "" {
		logger.ErrorLog.Println("-telegrambottoken is required")
		os.Exit(1)
	}

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

	// CreateListCitys()
	// go create_sitys_List()
}

func OsGetEnv() {
	tgToken = os.Getenv("TGTOKEN")
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
	postgree.PostName = os.Getenv("PostName")
	url = os.Getenv("URL")
	urlday = os.Getenv("URLDAY")
	imap.Speed_message = os.Getenv("Speed_message")
	imap.Hourly_report = os.Getenv("Hourly_report")
	imap.Behavior_report = os.Getenv("Behavior_report")
	imap.ImapClient = os.Getenv("IMAPCLIENT")
	imap.ImapFiles = os.Getenv("IMAPFILES")
	imap.ImapFilesPass = os.Getenv("IMAPFILESPASS")
	imap.ImapSpeed = os.Getenv("IMAPSPEED")
	imap.ImapSpeedPass = os.Getenv("IMAPSPEEDPASS")

}

func PrintGetEnv() {
	log.Println(tgToken)
	log.Println(crm.SshHost)
	log.Println(crm.SshPort)
	log.Println(crm.SshUser)
	log.Println(crm.DbUser)
	log.Println(crm.DbPass)
	log.Println(crm.DbHost)
	log.Println(crm.DbName)
	log.Println(crm.PrivateKeyPath)
	log.Println(postgree.PostHost)
	log.Println(postgree.PostPort)
	log.Println(postgree.PostPass)
	log.Println(postgree.PostName)
	log.Println(url)
	log.Println(urlday)
	log.Println(imap.Speed_message)
	log.Println(imap.Hourly_report)
	log.Println(imap.Behavior_report)
	log.Println(imap.ImapClient)
	log.Println(imap.ImapFiles)
	log.Println(imap.ImapFilesPass)
	log.Println(imap.ImapSpeed)
	log.Println(imap.ImapSpeedPass)
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
