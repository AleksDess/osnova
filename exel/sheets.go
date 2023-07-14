package exel

import (
	"context"
	"fmt"
	"log"
	"osnova/logger"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

// преобразует лист в [][]string
// n - сколько первых строк пропустить
func Filling_standart_row(a *sheets.ValueRange, n int) [][]string {

	res := make([][]string, 0)
	leights := 0
	for _, row := range a.Values {
		if len(row) > leights {
			leights = len(row)
		}
	}
	for nn, row := range a.Values {
		dob := leights - len(row)
		if nn < n {
			continue
		}
		rs := make([]string, 0)
		for _, i := range row {
			rs = append(rs, i.(string))
		}
		for i := 0; i < dob; i++ {
			rs = append(rs, "")
		}
		res = append(res, rs)
	}
	return res
}

// запись на лист googl
// s - имя листа
// d - диапазон для чтения листа
// z - диапазон для записи листа
// modify_nul - замена "0" на ""
// append_str - добавкка 1000 пустых строк для затирания предыдущей информации

func Rec(s, z string, res [][]string, modify_nul, append_str bool) (err error) {

	if s == "" || z == "" {
		logger.ErrorLog.Println("нет данных для записи таблицы")
		return fmt.Errorf("нет данных для записи таблицы")
	}

	srv := ServisSumyCashboClient()

	spreadsheetId := s
	writeRange := z

	var vr sheets.ValueRange

	for _, i := range res {
		myval := make([]interface{}, 0)
		for _, j := range i {
			if modify_nul && (j == "0" || j == "0.0" || j == "0.00") {
				j = ""
			}
			myval = append(myval, j)
		}
		vr.Values = append(vr.Values, myval)
	}

	// затирание предыдущей информации
	if append_str {
		dob := 1000
		col := 0
		for _, i := range res {
			if len(i) > col {
				col = len(i)
			}
		}

		for i := 0; i < dob+1; i++ {
			myval := make([]interface{}, 0)
			for j := 0; j < col; j++ {
				myval = append(myval, "")
			}
			vr.Values = append(vr.Values, myval)
		}
	}

	_, err = srv.Spreadsheets.Values.Update(spreadsheetId, writeRange, &vr).ValueInputOption("USER_ENTERED").Do()
	if err != nil {
		str := "Не вышло записать таблицу " + s + " " + " диапазон " + z + "\n" + err.Error()
		logger.ErrorLog.Println(str)
	}
	return
}

func Rec_Clear(s, z string, res [][]string, modify_nul, append_str bool) (err error) {

	if s == "" || z == "" {
		return fmt.Errorf("нет имени листа или диапазона для записи таблицы")
	}

	if len(res) == 0 {
		return fmt.Errorf("нет данных для записи таблицы")
	}

	srv := ServisSumyCashboClient()

	pr, err := ReadGoogleSheets(s, z+":ZZZ")
	if err != nil {
		return
	}

	dob1 := len(pr)

	spreadsheetId := s
	writeRange := z

	var vr sheets.ValueRange

	for _, i := range res {
		myval := make([]interface{}, 0)
		for _, j := range i {
			if modify_nul && (j == "0" || j == "0.0" || j == "0.00" || j == "0,0" || j == "0,00") {
				j = ""
			}
			myval = append(myval, j)
		}
		vr.Values = append(vr.Values, myval)
	}

	// затирание предыдущей информации
	if append_str {
		dob := dob1 - len(res) + 100
		col := 0
		for _, i := range res {
			if len(i) > col {
				col = len(i)
			}
		}

		for i := 0; i < dob+1; i++ {
			myval := make([]interface{}, 0)
			for j := 0; j < col; j++ {
				myval = append(myval, "")
			}
			vr.Values = append(vr.Values, myval)
		}
	}

	_, err = srv.Spreadsheets.Values.Update(spreadsheetId, writeRange, &vr).ValueInputOption("USER_ENTERED").Do()
	if err != nil {
		return
	}
	return
}

// добавляет данные на лист
// sumy-cashbox@up-statistics.iam.gserviceaccount.com
func Rec_Append(l, r string, res [][]string) (err error) {

	srv := ServisSumyCashboClient()

	var vr sheets.ValueRange

	for _, i := range res {
		mv := make([]interface{}, 0)
		for _, j := range i {
			mv = append(mv, j)
		}
		vr.Values = append(vr.Values, mv)
	}

	_, err = srv.Spreadsheets.Values.Append(l, r, &vr).ValueInputOption("USER_ENTERED").Do()

	if err != nil {
		fmt.Println("Не вышло записать таблицу"+"\n", err)
	}
	return
}

// читает лист гугл таблицы
// отдает типизованый
// по ширине результат
// n - сколько строк пропустить вначале
// sumy-cashbox@up-statistics.iam.gserviceaccount.com
func ReadGoogleSheets(spr, r string) (res [][]string, err error) {

	if spr == "" || r == "" {
		err = fmt.Errorf("нет имени файла или диапазона")
		return
	}

	srv := ServisSumyCashboClient()

	sheet, err := srv.Spreadsheets.Values.Get(spr, r).Do()
	if err != nil {
		return
	}
	res = Filling_standart_row(sheet, 0)

	return
}

// сервис аккаунт
// Виталик безлимит
func ServisSumyCashboClient() (srv *sheets.Service) {

	//srv, err := sheets.NewService(context.Background(), option.WithCredentialsFile("C:/Reverso_Context/cashbox_key.json"))
	srv, err := sheets.NewService(context.Background(), option.WithCredentialsJSON(Create_Key_Servise_Accaunt_UMLIMITED()))
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}
	return
}

func СolumnNumberToLetters(columnNumber int) string {
	result := ""
	for columnNumber > 0 {
		columnNumber-- // Отнимаем 1, чтобы сделать это 0-индексированным
		remainder := columnNumber % 26
		result = fmt.Sprint(remainder+65) + result
		columnNumber = columnNumber / 26
	}
	return result
}
