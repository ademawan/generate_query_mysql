package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
	"time"
	"unicode"
)

type UserRequestCreate struct {
	Email     string  `json:"email" test:"-"`
	Name      string  `json:"name" test:"-"`
	Username  string  `json:"username" test:"-"`
	Phone     string  `json:"phone" test:"-"`
	Password  string  `json:"password" `
	Photo     string  `json:"photo" `
	Latitude  float64 `json:"latitude" `
	Longitude float64 `json:"longitude" `
}
type Payment struct {
	Id             string    `json:"id"`
	OrderId        string    `json:"order_id"`
	TransactionId  string    `json:"transaction_id"`
	PaymentType    string    `json:"payment_type"`
	PaymentMethod  string    `json:"payment_method"`
	PaymentStatus  string    `json:"payment_status"`
	Amount         int       `json:"amount"`
	AdminFee       int       `json:"admin_fee"`
	DeviceID       string    `json:"device_id"`
	MsisdnSender   string    `json:"msisdn_sender"`
	MsisdnReceiver string    `json:"msisdn_receiver"`
	ProductName    string    `json:"product_name"`
	BillingNumber  string    `json:"billing_number"`
	PurchaseMode   string    `json:"purchase_mode"`
	CreatedAt      time.Time `json:"created_at" test:"-"`
}
type Base struct {
	id   int
	name string
}

type Extended struct {
	Base
	Email    string
	Password string
}

// type ResponseSubscriberProfile struct {
// 	Profiles Profiles `json:"profile"`
// }

func main() {

	e := Extended{}
	e.Email = "me@mail.com"
	e.Password = "secret"

	for i := 0; i < reflect.TypeOf(e).NumField(); i++ {
		if reflect.ValueOf(e).Field(i).Kind() != reflect.Struct {
			fmt.Println(reflect.ValueOf(e).Field(i))
		}
	}

	fmt.Println("hallp")
	// getPropertyInfo(&UserRequestCreate{})
	// generateQueryInsert(&UserRequestCreate{})
	// generateQueryInsert(&UserRequestCreate{}, "query_insert", "user")
	generateQueryInsert(&Payment{}, "query_insert", "payment")

	var data2 interface{}
	GetFile("", "query_insert", "go", &data2)
	fmt.Println(data2)
	te := strings.Split(data2.(string), "//")
	fmt.Println(te[0])
}
func getPropertyInfo(s *UserRequestCreate) {
	var reflectValue = reflect.ValueOf(s)

	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}

	var reflectType = reflectValue.Type()

	for i := 0; i < reflectValue.NumField(); i++ {
		f, _ := reflectType.FieldByName("Email")
		fmt.Println(f.Tag)
		fmt.Println(f.Tag.Lookup("test"))
		fmt.Println("nama      :", reflectType.Field(i).Name)
		fmt.Println("tipe data :", reflectType.Field(i).Type)
		fmt.Println("nilai     :", reflectValue.Field(i).Interface())
		fmt.Println("")
	}
}

func generateQueryInsert(data interface{}, nameOfFile, nameOfTable string) {
	var reflectValue = reflect.ValueOf(data)

	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}

	var reflectType = reflectValue.Type()
	query := `INSERT INTO `
	query += nameOfTable + `(`
	var queryValues = make([]string, 0)
	var queryField = make([]string, 0)
	goFieldPointerReference := make([]string, 0)
	for i := 0; i < reflectValue.NumField(); i++ {
		field := reflectType.Field(i).Name
		f, _ := reflectType.FieldByName(field)
		val, _ := f.Tag.Lookup("test")

		//get pointer of field
		if val != "-" {
			goFieldPointerReference = append(goFieldPointerReference, nameOfTable+`.`+reflectType.Field(i).Name)
		}
		//
		newField := ""
		if val != "-" {
			for j := 0; j < len(field); j++ {

				res_1 := unicode.IsUpper(rune(field[j]))
				if res_1 {
					if res_1 && j != 0 {
						newField += "_"
					}
					newField += strings.ToLower(string(field[j]))
				} else {
					newField += string(field[j])
				}

			}
			queryField = append(queryField, newField)
			queryValues = append(queryValues, "?")
		}
	}
	query += strings.Join(queryField, ",")
	query += `)VALUES(`
	query += strings.Join(queryValues, ",")
	query += `)`
	fmt.Println(query)
	_ = ioutil.WriteFile(nameOfFile+".sql", []byte(query), 0644)

	goFieldPointerReferenceString := `
	package main
	//
	func Create(){
	query := ` + `"` + query + `"
	`
	goFieldPointerReferenceString += `stmt, err := r.db.Prepare(query)
	if err != nil {
		return err, r.log

	}
	_, err = stmt.Exec(`
	goFieldPointerReferenceString += strings.Join(goFieldPointerReference, ",")
	goFieldPointerReferenceString += `)
	`
	goFieldPointerReferenceString += `if err != nil {
		r.log.Message += "|Exec|" + err.Error()
		return err, r.log
	}
	
	}`
	_ = ioutil.WriteFile(nameOfFile+".go", []byte(goFieldPointerReferenceString), 0644)
	// fmt.Println(goFieldPointerReferenceString)
}

func generateQueryInsertFileGo(data interface{}, nameOfFile, nameOfTable string) {
	var reflectValue = reflect.ValueOf(data)

	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}

	var reflectType = reflectValue.Type()
	query := `INSERT INTO `
	query += nameOfTable + `(`
	var queryValues = make([]string, 0)
	var queryField = make([]string, 0)

	for i := 0; i < reflectValue.NumField(); i++ {
		field := reflectType.Field(i).Name
		f, _ := reflectType.FieldByName(field)
		val, _ := f.Tag.Lookup("test")

		newField := ""
		if val != "-" {
			for i := 0; i < len(field); i++ {
				res_1 := unicode.IsUpper(rune(field[i]))
				if res_1 {
					if res_1 && i != 0 {
						newField += "_"
					}
					newField += strings.ToLower(string(field[i]))
				} else {
					newField += string(field[i])
				}

			}

			queryField = append(queryField, newField)
			queryValues = append(queryValues, "?")
		}
	}
	query += strings.Join(queryField, ",")
	query += `)VALUES(`
	query += strings.Join(queryValues, ",")
	query += `)`
	fmt.Println(query)
	_ = ioutil.WriteFile(nameOfFile+".sql", []byte(query), 0644)

}

// func generateQueryInsert(data interface{}, nameOfFile, nameOfTable string) {
// 	var reflectValue = reflect.ValueOf(data)

// 	if reflectValue.Kind() == reflect.Ptr {
// 		reflectValue = reflectValue.Elem()
// 	}

// 	var reflectType = reflectValue.Type()
// 	query := `INSERT INTO `
// 	query += nameOfTable + `(`
// 	var queryValues = make([]string, 0)
// 	var queryField = make([]string, 0)
// 	goFieldPointerReference := make([]string, 0)
// 	for i := 0; i < reflectValue.NumField(); i++ {
// 		field := reflectType.Field(i).Name
// 		f, _ := reflectType.FieldByName(field)
// 		val, _ := f.Tag.Lookup("test")

// 		//get pointer of field
// 		goFieldPointerReference = append(goFieldPointerReference, `&`+nameOfTable+`.`+reflectType.Field(i).Name)
// 		//
// 		newField := ""
// 		if val != "-" {
// 			for j := 0; j < len(field); j++ {

// 				res_1 := unicode.IsUpper(rune(field[j]))
// 				if res_1 {
// 					if res_1 && j != 0 {
// 						newField += "_"
// 					}
// 					newField += strings.ToLower(string(field[j]))
// 				} else {
// 					newField += string(field[j])
// 				}

// 			}
// 			queryField = append(queryField, newField)
// 			queryValues = append(queryValues, "?")
// 		}
// 	}
// 	query += strings.Join(queryField, ",")
// 	query += `)VALUES(`
// 	query += strings.Join(queryValues, ",")
// 	query += `)`
// 	fmt.Println(query)
// 	_ = ioutil.WriteFile(nameOfFile+".sql", []byte(query), 0644)

// 	goFieldPointerReferenceString := strings.Join(goFieldPointerReference, ",")
// 	_ = ioutil.WriteFile(nameOfFile+".go", []byte(goFieldPointerReferenceString), 0644)

// }
func GetFile(baseDir, fileName, ext string, data *interface{}) {
	jsonFile, err := os.Open(baseDir + fileName + "." + ext)
	if err != nil {

		fmt.Println(err.Error())
	}

	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err.Error())
	}

	// json.Unmarshal(byteValue, &data)
	fmt.Println("hallo")
	*data = string(byteValue)

}

func checkTypeOf(x interface{}) {
	var reflectValue = reflect.ValueOf(x)

	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}

	var reflectType = reflectValue.Type()

	for i := 0; i < reflectValue.NumField(); i++ {
		f, _ := reflectType.FieldByName("Email")
		fmt.Println(f.Tag)
		fmt.Println(f.Tag.Lookup("test"))
		fmt.Println("nama      :", reflectType.Field(i).Name)
		fmt.Println("tipe data :", reflectType.Field(i).Type)
		fmt.Println("nilai     :", reflectValue.Field(i).Interface())
		fmt.Println("")
	}
}
