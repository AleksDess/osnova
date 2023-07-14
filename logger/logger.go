package logger

import (
	"fmt"
	"log"
	"os"
	"osnova/path"
)

var InfoLog *log.Logger
var ErrorLog *log.Logger
var MessLog *log.Logger

func RunLogger() {

	f_log, err := os.Create("C:/bolt_db/bolt.log")

	if err != nil {
		fmt.Println(err)
	}

	// defer f_log.Close()

	if path.Place == "server" {
		InfoLog = log.New(f_log, "INFO\t", log.Ldate|log.Ltime)
		ErrorLog = log.New(f_log, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
		MessLog = log.New(f_log, "MESSAGE\t", log.Ldate|log.Ltime|log.Lshortfile)
	} else {
		InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
		ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
		MessLog = log.New(os.Stdout, "MESSAGE\t", log.Ldate|log.Ltime|log.Lshortfile)
	}
}
