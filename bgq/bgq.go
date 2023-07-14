package BGQ

import (
	"context"
	"fmt"
	"os"
	"osnova/trs"
	"strconv"
	"strings"
	"time"

	"cloud.google.com/go/bigquery"
	"cloud.google.com/go/civil"
	"google.golang.org/api/iterator"
)

// создает клиент BigQuery
func Get_BG_client() (client *bigquery.Client, ctx context.Context) {
	//projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
	projectID := "up-statistics"
	if projectID == "" {
		fmt.Println("GOOGLE_CLOUD_PROJECT environment variable must be set.")
	}
	ctx = context.Background()
	var err error
	for {
		client, err = bigquery.NewClient(ctx, projectID)
		if err != nil {
			fmt.Println("Невозможно создать клиент стр 27")
			fmt.Println(err)
			time.Sleep(10 * time.Second)
			continue
		} else {
			break
		}
	}

	defer client.Close()
	return
}

// read BigQuery
func Read_BGQ(sc string) (res [][]bigquery.Value, err error) {

	client, ctx := Get_BG_client()

	querys := client.Query(sc)

	iters, err := querys.Read(ctx)

	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		var row []bigquery.Value
		err = iters.Next(&row)

		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println(3, err)
			return
		}

		res = append(res, row)

	}
	return res, nil
}

type Kassa_BGQ struct {
	Created_Time   time.Time
	User           string
	Contractor     string
	Coment         string
	Driver_Name    string
	Car_Reg_Number string
	Amount         float64
	Basis          string
	Inv            string
	Week_number    int64
	Box_Type       string
	Uuid           string `bigquery:"uuid"`
	City           string
	Value_date     civil.Date `bigquery:"value_date"`
}

type List_Kassa_BGQ []Kassa_BGQ

func (a *List_Kassa_BGQ) Amount() (r float64) {
	for _, i := range *a {
		if i.Inv == "1.2 - Каса водія готівка" {
			continue
		}
		r += i.Amount
	}
	return
}

func (a *List_Kassa_BGQ) Print() {
	for _, i := range *a {
		i.Print()
	}
}

func (a *Kassa_BGQ) CreateStringSlise() []string {
	return []string{a.Created_Time.Format("02.01.2006 15:04:05"), a.User, a.Contractor, a.Coment, a.Driver_Name, a.Car_Reg_Number, strings.Replace(fmt.Sprint(a.Amount), ".", ",", 1), a.Basis, a.Inv, fmt.Sprint(a.Week_number), a.Box_Type, a.City}
}

func (a *Kassa_BGQ) Add(b *Kassa_BGQ) {
	a.Amount += b.Amount
}

func (a *Kassa_BGQ) CreateSlise(pp []string, city string) (err error) {
	t, err := time.Parse("1/2/2006 15:04:05", pp[0])
	if err != nil {
		return
	}

	a.Created_Time = t
	a.User = "DriverCSS"
	a.Coment = pp[4]
	a.Driver_Name = pp[5]
	a.Car_Reg_Number = trs.Trs(pp[6])
	a.Amount, err = strconv.ParseFloat(pp[7], 64)
	if err != nil {
		return
	}
	a.Basis = pp[8]
	a.Inv = pp[9]
	week, err := strconv.ParseFloat(pp[10], 64)
	if err != nil {
		return
	}
	a.Week_number = int64(week)

	a.Box_Type = "Готівка"
	a.Uuid = trs.Generate_Uuid()
	a.City = city
	a.Value_date = civil.Date{Year: t.Year(), Month: t.Month(), Day: t.Day()}
	return
}

func (a *Kassa_BGQ) Print() {
	fmt.Println()
	fmt.Println("type Kassa_BGQ struct   :")
	fmt.Println("Created_Time            :", a.Created_Time)
	fmt.Println("User                    :", a.User)
	fmt.Println("Contractor              :", a.Contractor)
	fmt.Println("Coment                  :", a.Coment)
	fmt.Println("Driver_Name             :", a.Driver_Name)
	fmt.Println("Car_Reg_Number          :", a.Car_Reg_Number)
	fmt.Println("Amount                  :", a.Amount)
	fmt.Println("Basis                   :", a.Basis)
	fmt.Println("Inv                     :", a.Inv)
	fmt.Println("Week_number             :", a.Week_number)
	fmt.Println("Box_Type                :", a.Box_Type)
	fmt.Println("uuid                    :", a.Uuid)
	fmt.Println("City                    :", a.City)
	fmt.Println("Value_date              :", a.Value_date)

}

type FiveString struct {
	B string
	C string
	D string
	E string
	F string
}

func (a *FiveString) Create(i []bigquery.Value) {
	if i[0] != nil {
		a.B = i[0].(time.Time).Format("02.01.2006")
	}
	if i[5] != nil {
		a.C = i[5].(string)
	}
	if i[6] != nil {
		a.D = fmt.Sprintf("%.2f", i[6].(float64))
	}
	if i[8] != nil {
		a.E = i[8].(string)
	}
	if i[9] != nil {
		a.F = fmt.Sprint(i[9].(int64))
	}
}

func (a *FiveString) Print() {
	fmt.Println(a.B, a.C, a.D, a.E, a.F)
}

func (a *Kassa_BGQ) Create(i []bigquery.Value) {
	if i[0] != nil {
		a.Created_Time = i[0].(time.Time)
	}
	if i[1] != nil {
		a.User = i[1].(string)
	}
	if i[2] != nil {
		a.Contractor = i[2].(string)
	}
	if i[3] != nil {
		a.Coment = i[3].(string)
	}
	if i[4] != nil {
		a.Driver_Name = i[4].(string)
	}
	if i[5] != nil {
		a.Car_Reg_Number = i[5].(string)
	}
	if i[6] != nil {
		a.Amount = i[6].(float64)
	}
	if i[7] != nil {
		a.Basis = i[7].(string)
	}
	if i[8] != nil {
		a.Inv = i[8].(string)
	}
	if i[9] != nil {
		a.Week_number = i[9].(int64)
	}
	if i[10] != nil {
		a.Box_Type = i[10].(string)
	}
	if i[11] != nil {
		a.Uuid = i[11].(string)
	}
	if i[12] != nil {
		a.City = i[12].(string)
	}
}

// запись в кассы бигквери
func Save_kassa_BGQ(dataset, table string, myval []*Kassa_BGQ) error {

	client, ctx := Get_BG_client()

	datas := client.Dataset(dataset)
	tabl := datas.Table(table)

	u := tabl.Uploader()
	err := u.Put(ctx, myval)

	return err
}

type Drivers struct {
	Uuid               string
	Created_time       time.Time
	Name               string
	Mail               string
	Phone              string
	Taxi_id            string
	ContractId         string
	Status             string
	City               string
	Taxi_type          string
	Fired_date         time.Time
	Bolt_online_status string
	Tariff             string
	Junior_bot         string
}

func (a *Drivers) Create(row []bigquery.Value) {
	if row[0] != nil {
		a.Uuid = row[0].(string)
	}
	if row[1] != nil {
		a.Created_time = row[1].(time.Time)
	}
	if row[2] != nil {
		a.Name = row[2].(string)
	}
	if row[3] != nil {
		a.Mail = row[3].(string)
	}
	if row[4] != nil {
		a.Phone = row[4].(string)
	}
	if row[5] != nil {
		a.Taxi_id = row[5].(string)
	}
	if row[6] != nil {
		a.ContractId = row[6].(string)
	}
	if row[7] != nil {
		a.Status = row[7].(string)
	}
	if row[8] != nil {
		a.City = row[8].(string)
	}
	if row[9] != nil {
		a.Taxi_type = row[9].(string)
	}
	if row[10] != nil {
		t := row[10].(civil.Date)
		a.Fired_date = time.Date(t.Year, t.Month, t.Day, 7, 7, 7, 0, time.UTC)
	}
	if row[11] != nil {
		a.Bolt_online_status = row[11].(string)
	}
	if row[12] != nil {
		a.Tariff = row[12].(string)
	}
	if row[13] != nil {
		a.Junior_bot = row[13].(string)
	}
}

func (a *Drivers) Print() {
	fmt.Println()
	fmt.Println("a *Drivers            :")
	fmt.Println("a.uuid                :", a.Uuid)
	fmt.Println("a.created_time        :", a.Created_time)
	fmt.Println("a.name                :", a.Name)
	fmt.Println("a.mail                :", a.Mail)
	fmt.Println("a.phone               :", a.Phone)
	fmt.Println("a.taxi_id             :", a.Taxi_id)
	fmt.Println("a.contractId          :", a.ContractId)
	fmt.Println("a.status              :", a.Status)
	fmt.Println("a.city                :", a.City)
	fmt.Println("a.taxi_type           :", a.Taxi_type)
	fmt.Println("a.fired_date          :", a.Fired_date)
	fmt.Println("a.bolt_online_status  :", a.Bolt_online_status)
	fmt.Println("a.tariff              :", a.Tariff)
	fmt.Println("a.Junior_bot          :", a.Junior_bot)
}

func Get_Drivers() (res []Drivers) {
	list, _ := Read_BGQ("SELECT * FROM DB_External.drivers")
	for _, row := range list {
		r := Drivers{}
		r.Create(row)
		res = append(res, r)
	}
	return
}

func GetDriversCity(c string) (res []Drivers) {
	list, _ := Read_BGQ(fmt.Sprintf("SELECT * FROM DB_External.drivers WHERE city ='%s'", c))
	for _, row := range list {
		r := Drivers{}
		r.Create(row)
		res = append(res, r)
	}
	return
}

func GetDriversName(c string) (res []Drivers) {
	list, _ := Read_BGQ(fmt.Sprintf("SELECT * FROM DB_External.drivers WHERE name ='%s'", c))
	for _, row := range list {
		r := Drivers{}
		r.Create(row)
		res = append(res, r)
	}
	return
}

// строит схему всех доступных датасетов
func Schema_dataset() {

	client, ctx := Get_BG_client()

	it := client.Datasets(ctx)
	s := ""
	ss := ""
	for {
		dataset, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		datasetID := dataset.DatasetID
		s = s + "\n*************************************\n"
		s = s + dataset.DatasetID
		s = s + "\n*************************************\n"

		ss = ss + "\n*************************************\n"
		ss = ss + dataset.DatasetID
		ss = ss + "\n*************************************\n"
		fmt.Println(dataset.DatasetID)

		meta, err := client.Dataset(datasetID).Metadata(ctx)
		if err != nil {
			fmt.Println(err)
		}

		s = s + fmt.Sprintf("Dataset ID: %s\n", datasetID) + "\n"
		ss = ss + fmt.Sprintf("Dataset ID: %s\n", datasetID) + "\n"

		for k, v := range meta.Labels {
			fmt.Printf("\t%s: %s", k, v)
		}
		s = s + "Tables:" + "\n"
		ss = ss + "Tables:" + "\n"
		it := client.Dataset(datasetID).Tables(ctx)

		cnt := 0
		for {
			t, err := it.Next()
			if err == iterator.Done {
				break
			}
			cnt++
			s = s + fmt.Sprintf("\t%s\n", t.TableID)
			ss = ss + fmt.Sprintf("\t%s\n", t.TableID)
			tableRef := client.Dataset(datasetID).Table(t.TableID)
			meta, err := tableRef.Metadata(ctx)
			if err != nil {
				fmt.Println(err)
				continue
			}
			s = s + "Schema:\n"
			for _, i := range meta.Schema {
				s = s + fmt.Sprintln(i.Name, i.Type)
			}
			s = s + "-----------------------" + "\n\n"
		}
		if cnt == 0 {
			s = s + ("\tThis dataset does not contain any tables.")
		}
	}

	data := []byte(s)
	file, err := os.Create("datasetschema.txt")
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()
	file.Write(data)

	datas := []byte(ss)
	files, errs := os.Create("dataset.txt")
	if errs != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer files.Close()
	files.Write(datas)

	fmt.Println("Done.")
}
