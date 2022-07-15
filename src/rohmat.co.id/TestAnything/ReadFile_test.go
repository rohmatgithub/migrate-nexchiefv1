package TestAnything

import (
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"math"
	"net/smtp"
	"os"
	"reflect"
	"rohmat.co.id/model"
	"rohmat.co.id/service"
	"rohmat.co.id/util"
	"strconv"
	"strings"
	"testing"
	"time"
	"unicode/utf8"
)

func TestFloat(t *testing.T) {
	f, err := os.Open("C:\\cdc-tools\\data sql\\keyaccount1.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		fmt.Println(err)
		return
	}
	for i, line := range lines {
		fmt.Println(i, ":", line, len(line))
	}
}
func nextRune(b []byte) rune {
	r, _ := utf8.DecodeRune(b)
	return r
}
func TestByte(t *testing.T) {
	str := "\"id\",\"principalID\",\"keyaccount1Code\",\"name\",\"\\\"createdDate\\\"\",\"modifiedDate\",\"modifiedBy\""

	b := []byte(str)
	fmt.Println(string(b))
	b = b[len(`"`):]
	fmt.Println(len(`"`))
	fmt.Println(string(b))
	//r := nextRune(b)
	//fmt.Println(r)
}
func TestCopyArrat(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	arr = arr[0:2]
	fmt.Println(arr)
}
func TestValidasiFileCsv(t *testing.T) {
	var (
		//s time.Time // current time
		//file *os.File
		fileByte []byte
		//fileStat os.FileInfo
		err error
		//lastLineSize int64
	)
	//s = time.Now()

	if fileByte, err = ioutil.ReadFile("C:\\cdc-tools\\data sql\\keyaccount1.csv"); err != nil {
		fmt.Println(err)
		return
	}
	everyPrint := (len(fileByte) / 100) + 1
	if everyPrint == 0 {
		everyPrint++
	}
	//fmt.Println(len(fileByte), "every print :", everyPrint)
	var cek, countLine int
	countLine++
	for i := 0; i < len(fileByte); i++ {
		if fileByte[i] == []byte("\n")[0] {
			if countLine == 1 && !cekDelimiter(fileByte[0:i]) {
				fmt.Println("invalid delimiter")
				return
			}
			countLine++
		}
		if i%everyPrint == 0 || i == len(fileByte)-1 {
			cek++
			presents := float64(i) / float64(len(fileByte)) * 100.0
			fmt.Println(cek, ":", math.Round(presents))
		}
	}
}

func cekDelimiter(line []byte) bool {
	temp := string(line)
	arr := strings.Split(temp, ",")
	if len(arr) < 2 {
		arr = strings.Split(temp, ";")
	}
	if len(arr) < 2 {
		arr = strings.Split(temp, "|")
	}
	return len(arr) >= 2
}

func TestReadFileCsv(t *testing.T) {
	//delimiter := ","
	//f, err := os.Open("C:\\cdc-tools\\data sql\\keyaccount1.csv")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//var totalColumn int
	//scanner := bufio.NewScanner(f)
	//scanner.Split(bufio.ScanLines)
	//for scanner.Scan() {
	//	dataLine := scanner.Text()
	//	totalColumn = len(strings.Split(dataLine, delimiter))
	//	fmt.Println(dataLine)
	//}
}

func TestTagName(t *testing.T) {
	type structTest struct {
		Name sql.NullString `json:"name" index:"0"`
		Age  sql.NullInt32  `json:"age"`
		B    sql.NullBool   `json:"b"`
	}
	//str := `{"name":"rohmat","age":10}`
	//m := make(map[string]interface{})

	//var a interface{}
	var o structTest
	o.B.Bool = false
	o.Name.String = "Y"
	o.Age.Int32 = 2
	//r, _ := json.Marshal(o)
	//_ = json.Unmarshal([]byte(str), &m)
	//fmt.Println(m["name"])

	test(o)
	//fmt.Println(o.Name)

}

func test(o interface{}) {
	v := reflect.ValueOf(o)
	t := reflect.TypeOf(o)
	//fmt.Println(len(s.MapKeys()) == 0)
	//
	//for _, value := range s.MapKeys() {
	//	//v := s.MapIndex(value)
	//	fmt.Println(value.Convert(s.Type().Key()))
	//	//if v.Interface() == "Y" || v.Interface() == true{
	//	//	fmt.Println("true")
	//	//}
	//}
	//d := reflect.ValueOf(o).Elem()
	//
	//c := reflect.TypeOf(o).Elem()
	for i := 0; i < t.NumField(); i++ {
		//fmt.Println(c.Field(i).Tag.Get("index"))

		types := v.Field(i).Interface()
		switch types.(type){
		case sql.NullBool:
			c := types.(sql.NullBool)
			fmt.Println(c.Bool)
		}
		//types := t.Field(i).Type
		//fmt.Println(types == reflect.TypeOf(sql.NullBool{}))
		//fmt.Println(types.(sql.NullBool))

		//switch c.Field(i).Type.Kind() {
		//
		//}
		//if c.Field(i).Type.Kind() == reflect.String {
		//	temp := "tets"
		//tempValue := reflect.ValueOf(&temp).Elem()
		//d.Field(i).Set(tempValue)
	}

}

func TestMarshal(t *testing.T) {
	type j struct {
		Level1 int64 `json:"level1"`
		Level2 int64 `json:"level2"`
	}
	geoTreeString, _ := json.Marshal(j{
		Level1: 1,
		Level2: 2,
	})
	fmt.Println(string(geoTreeString))
}

func TestCreateFile(t *testing.T) {
	path := "C:\\cdc-tools\\data sql\\test\\test.txt"
	err := EnsureDir("C:\\cdc-tools\\data sql\\test")
	if err != nil {
		assert.FailNow(t, err.Error())
	}
	b, err := os.Create(path)
	if err != nil {
		assert.FailNow(t, err.Error())
	}
	fmt.Println(b)
}

func EnsureDir(dirName string) error {
	err := os.MkdirAll(dirName, 0777)
	if err == nil {
		return nil
	}
	if os.IsExist(err) {
		// check that the existing path is a directory
		info, err := os.Stat(dirName)
		if err != nil {
			return err
		}
		if !info.IsDir() {
			return errors.New("path exists but is not a directory")
		}
		return nil
	}
	return err
}

func TestReadJson(t *testing.T) {
	stream := service.NewJSONStream()
	go func() {
		for data := range stream.Watch() {
			if data.Error != nil {
				log.Println(data.Error)
			}
			var customer model.NexsellerCustomer
			_ = json.Unmarshal(data.Data, &customer)
			log.Println(">>", customer.Code)
		}
	}()
	stream.Start("C:\\cdc-tools\\data sql\\test.json")
}

type EntryData struct {
	Data    string
	Process bool
}

type StreamData struct {
	Stream chan EntryData
}

func (s StreamData) Watch() <-chan EntryData {
	return s.Stream
}
func (s StreamData) load() {
	for i := 0; i < 10; i++ {
		s.Stream <- EntryData{Data: "data ke -" + strconv.Itoa(i)}
	}
}

func TestWatch(t *testing.T) {
	stream := StreamData{Stream: make(chan EntryData)}
	go func() {
		for data := range stream.Watch() {
			fmt.Println(data.Data)
			time.Sleep(2 * time.Second)
			fmt.Println("tes")
		}
	}()
	stream.load()
}


func SendEmail(to []string, subject string, message string, mime string) error {
	defer func() {
		if r := recover(); r != nil {
			//TopRecoverLog("SendEmail.go", "SendEmail", r)
		}
	}()

	body :=
		"From: rohmatullah@nexsoft.co.id\n" +
			"To: " + strings.Join(to, ",") + "\n" +
			"Subject: " + subject + "\n" +
			mime +
			message

	auth := smtp.PlainAuth("Authentication Server", "rohmatullah@nexsoft.co.id", "5zGKQ82VKH", "smtp3.nexcloud.id")
	smtpAddr := fmt.Sprintf("%s:%d", "smtp3.nexcloud.id", 587)

	err := smtp.SendMail(smtpAddr, auth, "rohmatullah@nexsoft.co.id", to, []byte(body))
	if err != nil {
		return err
	}
	return nil
}

func TestSendEmail(t *testing.T){
	err := SendEmail([]string{"rohmattullah990@gmail.com"}, "TEST WOY", "test email", "")
	if err != nil {
		assert.FailNow(t, err.Error())
	}
}

func TestConnection(t *testing.T){
	db := util.ConnectDB()
	err := db.Ping()
	if err != nil {
		assert.FailNow(t, err.Error())
	}
}