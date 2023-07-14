package crm

// func AppendData(list [][]string, name string) (err error) {

// 	for n, i := range list[21] {
// 		if i == "" {
// 			if len(list[21][n-1]) < 10 || len(list[21][n-2]) < 10 {
// 				continue
// 			}
// 			if list[21][n-1][:10] == list[21][n-2][:10] {
// 				day := make([]string, 0)
// 				for j := 0; j < 20; j++ {
// 					day = append(day, fmt.Sprintf("=%s22+1", trs.СolumnNumberToLetters(n-1+j)))
// 				}
// 				fmt.Println(day)
// 				fmt.Println()
// 				rang := fmt.Sprintf("'ГРАФИК'!%s22", trs.СolumnNumberToLetters(n+1))
// 				trs.Rec(name, rang, [][]string{day}, false, false)
// 				break
// 			}
// 		}
// 	}
// 	return nil
// }

// func getListGrafik(city Citys.City) (list [][]string, err error) {
// 	list, err = trs.Read_sheets_CASHBOX_err(city.Name_grafik, "ГРАФИК", 0)
// 	if err != nil {
// 		return
// 	}
// 	return
// }

// func CheckAppendData(city Citys.City) (err error) {

// 	list, err := getListGrafik(city)
// 	if err != nil {
// 		return
// 	}

// 	for n, i := range list[21] {
// 		if n < 20 {
// 			continue
// 		}
// 		if i == "" {
// 			err := AppendData(list, city.Name_grafik)
// 			return err
// 		}
// 		t, _ := time.Parse(times.TNS, i[:10])
// 		if t.After(time.Now().AddDate(0, 0, 3)) {
// 			return nil
// 		}
// 	}
// 	return
// }

// func SaveDayGrafik(crmDB *sql.DB, city Citys.City, day time.Time) {

// 	list, err := getListGrafik(city)
// 	if err != nil {
// 		return
// 	}
// 	fmt.Println(len(list))

// 	n1 := 0
// 	for n, i := range list[21] {
// 		if len(i) < 10 {
// 			continue
// 		}
// 		if i[:10] == day.Format(times.TNS) {
// 			n1 = n
// 		}
// 	}

// 	g := ListGrafikPlus{}
// 	g.GetCityDayV1(crmDB, city.CRM_Name, day)

// 	for n, i := range g {
// 		fmt.Println(n, i)
// 	}

// 	z := make([][]string, len(list)-22)

// 	for m := 22; m < len(list); m++ {
// 		fmt.Println(m, list[m][4], list[m][6])
// 		for _, i := range g {
// 			if list[m][4] == fmt.Sprint(i.Mapon_id) {
// 				fmt.Println(list[m][4], list[m][6], i.Mapon_id, i.Car, m+1, n1, trs.СolumnNumberToLetters(n1), day.Format(times.TNS), i.ToString())
// 				z[m-22] = i.ToString()
// 				break
// 			}
// 		}
// 	}
// 	for _, i := range z {
// 		fmt.Println(i)
// 	}
// 	trs.Rec(city.Name_grafik, fmt.Sprintf("'ГРАФИК'!%s23", trs.СolumnNumberToLetters(n1)), z, false, false)

// }
