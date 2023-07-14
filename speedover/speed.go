package speedover

import (
	"database/sql"
	"fmt"
	"osnova/imap"
	"osnova/logger"
	"strings"
	"time"
	"trs"
)

// обновление в базе обработаных адресов
func WorkCheckListAdressSpeedOver(db *sql.DB) (err error) {

	list, err := trs.Read_sheets_CASHBOX_err("1qGlTouDcYquzBEZTCO7WK_QbQWG1SnjLki6AEC6ocGY", "'adress'!A2:F", 0)
	if err != nil {
		return
	}

	for _, i := range list {
		if i[5] == "y" {
			r := AdressSpeedOver{}
			r.Adress = i[1]
			r.Link = i[2]
			if i[3] == "y" {
				r.City = true
			}
			if i[4] == "y" {
				r.Overcity = true
			}
			r.In_work = true
			r.Print()
			r.DelDb(db)
			r.RecDB(db)
		}
	}
	return
}

// запись на лист необработанных адресов
// и верифицированных адресов
func SaveCheckListAdressSpeedOver() (err error) {

	dbase := sql.DB{}
	db := &dbase

	err = WorkCheckListAdressSpeedOver(db)
	if err != nil {
		return
	}

	listadress := ListAdressSpeedOver{}
	listadress.ReadDb(db)

	res := make([][]string, 0)
	ver := make([][]string, 0)
	for n, i := range listadress {
		if !strings.Contains(i.Link, "https") {
			r := i
			r.Link = "h" + i.Link
			i.DelDb(db)
			r.RecDB(db)
			i = r
		}

		if i.IsCity() {
			if !i.In_work {
				res = append(res, []string{fmt.Sprint(n + 1), i.Adress, i.Link, "n", "n", "n"})
			} else {
				var s1, s2 string = "n", "n"
				if i.City {
					s1 = "y"
				}
				if i.Overcity {
					s2 = "y"
				}
				ver = append(ver, []string{fmt.Sprint(n + 1), i.Adress, i.Link, s1, s2, "y"})
			}
		}
	}

	trs.Rec_Clear("1qGlTouDcYquzBEZTCO7WK_QbQWG1SnjLki6AEC6ocGY", "'adress'!A2", res, true, true)
	trs.Rec_Clear("1qGlTouDcYquzBEZTCO7WK_QbQWG1SnjLki6AEC6ocGY", "'verified'!A2", ver, true, true)
	return
}

// чтение писем и обработка сообщений
func ProcessingOfExceedanceLetters() {

	dbase := sql.DB{}
	db := &dbase

	listadress := ListAdressSpeedOver{}
	listadress.ReadDb(db)

	for {
		mbox, c, err := imap.GetImapBox(imap.ImapSpeed, imap.ImapSpeedPass, "INBOX")
		if err != nil {
			logger.ErrorLog.Println(err)
			<-time.After(time.Second)
			continue
		}
		n, speed := imap.Boltmessage(mbox, c, true)
		for _, i := range speed {
			r := AdressSpeedOver{}
			r.CreateSpeed(&i)
			if !r.IsCity() {
				continue
			}
			r.Adress = strings.ReplaceAll(r.Adress, "'", "")

			if !listadress.СheckФvailability(r.Adress) {
				r.RecDB(db)
				listadress = append(listadress, r)
			}
		}
		if n > 1000 {
			<-time.After(time.Second)
		} else if n > 100 {
			<-time.After(3 * time.Second)
		} else if n > 10 {
			<-time.After(30 * time.Second)
		} else {
			<-time.After(60 * time.Second)
		}
		c.Close()
	}
}
