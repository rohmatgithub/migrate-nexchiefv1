package main
//
//import (
//	"archive/zip"
//	"bufio"
//	"bytes"
//	"crypto/md5"
//	"database/sql"
//	"encoding/json"
//	"fmt"
//	"github.com/360EntSecGroup-Skylar/excelize"
//	_ "github.com/lib/pq"
//	"github.com/robfig/cron"
//	"github.com/stretchr/testify/assert"
//	"io"
//	"io/ioutil"
//	"os"
//	"path/filepath"
//	"reflect"
//	"strconv"
//	"strings"
//	"testing"
//	"time"
//	"unicode/utf8"
//	"unsafe"
//)
//
//type tes1 struct {
//	fileName string
//	str      []string
//}
//
//func (in tes1) insert1() {
//	in.str = append(in.str, "1")
//	in.fileName = "Product_test.go"
//	in.insert2()
//}
//func (in tes1) insert2() {
//	fmt.Println(in.fileName)
//}
//func (in tes1) insert3() {
//	in.str = append(in.str, "3")
//}
//
//func TestStruct1(t *testing.T) {
//	s := tes1{}
//	s.insert1()
//}
//func TestCase(t *testing.T) {
//	tesCase("dua")
//}
//func tesCase(str string) {
//	switch str {
//	case "satu":
//		str = "satu"
//	case "dua":
//		str = "like"
//	}
//	fmt.Println(str)
//}
//func TestMap(t *testing.T) {
//	maps := make(map[string][]string)
//	maps["tes1"] = []string{"1", "2"}
//	maps["tes2"] = []string{}
//
//	for s, i := range maps {
//		if maps[s] != nil {
//			fmt.Println(i)
//		}
//	}
//}
//func TestRemoveElementFromSlice(t *testing.T) {
//	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
//	for i := 0; i < len(slice); i++ {
//		if slice[i] == 10 {
//			fmt.Println(i)
//			slice = append(slice[:i], slice[i+1:]...)
//		}
//	}
//	fmt.Println(slice)
//}
//func TestString(t *testing.T) {
//	str := "insert-own"
//	result := strings.TrimSuffix(str, "-own")
//	fmt.Println(result)
//	//result := strings.HasSuffix(str, "-own")
//	//assert.Equal(t, true, result)
//}
//func CheckIsFileZip(src string) (isZip bool, err error) {
//	zipFile, errS := zip.OpenReader(src)
//	if errS != nil {
//		if errS == zip.ErrFormat {
//			return
//		}
//		err = errS
//		return
//	}
//	for i := 0; i < len(zipFile.File); i++ {
//
//	}
//	isZip = true
//	_ = zipFile.Close()
//	return
//}
//func UnzipFromByte(data []byte) (output []*zip.File, err error) {
//	//fmt.Println(len(data))
//	unziped, err := zip.NewReader(bytes.NewReader(data), int64(len(data)))
//	if err != nil {
//		return
//	}
//	output = unziped.File
//	return
//}
//
//func TestViewSizeMemory(t *testing.T) {
//	var s = "table_name"
//	query :=
//		"INSERT INTO  " +
//			s +
//			" ( company_profile_id, company_title, name,  " + // 1, 2, 3
//			"	npwp, address_1, address_2, " + // 4, 5, 6
//			"   address_3, hamlet, neighbour, " + // 7,8,9
//			"   country, district, sub_district, " + // 10,11,12
//			"   urban_village, island, postal_code, " + // 13,14,15
//			"	phone, fax, email, " + // 16,17,18
//			"	logo, created_client, created_at, " + // 19, 20, 21
//			"	created_by, updated_client, updated_at, " + // 22, 23, 24
//			"	updated_by ) " + // 25
//			"VALUES " +
//			" (	$1, $2, $3, " +
//			"  	$4, $5, $6, " +
//			"  	$7, $8, $9, " +
//			"  	$10, $11, $12, " +
//			"  	$13, $14, $15, " +
//			"  	$16, $17, $18," +
//			"  	$19, $20, $21," +
//			"  	$22, $23, $24, " +
//			"	$25 )"
//
//	query += " test woy"
//
//	fmt.Println(unsafe.Sizeof(query))
//	fmt.Println(len(query))
//	fmt.Println(utf8.RuneCountInString(query))
//}
//func BenchmarkPrintString1(b *testing.B) {
//	var s = "table_name"
//	query :=
//		`INSERT INTO  ` +
//			s +
//			` ( company_profile_id, company_title, name,
//				npwp, address_1, address_2,
//			   	address_3, hamlet, neighbour,
//			   	country, district, sub_district,
//			   	urban_village, island, postal_code,
//				phone, fax, email,
//				logo, created_client, created_at,
//				created_by, updated_client, updated_at,
//				updated_by )
//			VALUES
//			 (	$1, $2, $3,
//			  	$4, $5, $6,
//			  	$7, $8, $9,
//			  	$10, $11, $12,
//			  	$13, $14, $15,
//			  	$16, $17, $18,
//			  	$19, $20, $21,
//			  	$22, $23, $24,
//				$25 ) `
//	for i := 0; i < b.N; i++ {
//		fmt.Println(query)
//	}
//}
//func BenchmarkPrintString(b *testing.B) {
//	var s = "table_name"
//	query :=
//		"INSERT INTO  " +
//			s +
//			" ( company_profile_id, company_title, name,  " + // 1, 2, 3
//			"	npwp, address_1, address_2, " + // 4, 5, 6
//			"   address_3, hamlet, neighbour, " + // 7,8,9
//			"   country, district, sub_district, " + // 10,11,12
//			"   urban_village, island, postal_code, " + // 13,14,15
//			"	phone, fax, email, " + // 16,17,18
//			"	logo, created_client, created_at, " + // 19, 20, 21
//			"	created_by, updated_client, updated_at, " + // 22, 23, 24
//			"	updated_by ) " + // 25
//			"VALUES " +
//			" (	$1, $2, $3, " +
//			"  	$4, $5, $6, " +
//			"  	$7, $8, $9, " +
//			"  	$10, $11, $12, " +
//			"  	$13, $14, $15, " +
//			"  	$16, $17, $18," +
//			"  	$19, $20, $21," +
//			"  	$22, $23, $24, " +
//			"	$25 )"
//	for i := 0; i < b.N; i++ {
//		fmt.Println(query)
//	}
//}
//func TestCheckExtFile(t *testing.T) {
//	src := "C:\\nco_file\\test\\orderd.zip"
//
//	var (
//		temp      []byte
//		fileUnzip *zip.File
//		zips      []*zip.File
//		result    *bufio.Scanner
//	)
//	zipFile, errS := zip.OpenReader(src)
//	if errS != nil {
//		fmt.Println("1 :", errS)
//		return
//	}
//	fileUnzip = zipFile.File[0]
//	openFile, errS := fileUnzip.Open()
//	if errS != nil {
//		fmt.Println("2", errS)
//		return
//	}
//	defer openFile.Close()
//	fmt.Println("before :", openFile)
//	if filepath.Ext(fileUnzip.Name) == ".zip" {
//		temp, errS = ioutil.ReadAll(openFile)
//		if errS != nil {
//			fmt.Println("3", errS)
//			return
//		}
//		zips, errS = UnzipFromByte(temp)
//		if errS != nil && errS != zip.ErrFormat {
//			fmt.Println("4 :", errS)
//			return
//		}
//
//		if zips != nil {
//			fmt.Println("zip not nil")
//			openFile, errS = zips[0].Open()
//			if errS != nil {
//				fmt.Println("5", errS)
//				return
//			}
//			defer openFile.Close()
//		}
//	}
//	result = bufio.NewScanner(openFile)
//
//	for result.Scan() {
//		fmt.Println(result.Text())
//	}
//	_ = zipFile.Close()
//	return
//	//isZip = true
//	//isZip, err := CheckIsFileZip("C:\\nco_file\\test\\NCO_NDI_B0006_20200722_002")
//	//if err != nil{
//	//	fmt.Println("err :", err.Error())
//	//}
//	//if isZip{
//	//	fmt.Println("fileUnzip is zip")
//	//}else {
//	//	fmt.Println("file not zip")
//	//}
//	//pth, _ := os.Open("C:\\nco_file\\test\\NCO_NDI_B0006_20200722_001.zip")
//	//fd := bufio.NewScanner(pth)
//	//for fd.Scan() {
//	//	fmt.Println(fd.Text())
//	//}
//}
//func TestPermissionPath(t *testing.T) {
//
//	path := "C:\\nco_file\\inbound"
//	info, errs := os.Stat(path)
//	if errs != nil {
//		t.Error(errs)
//	}
//	t.Log("before info : ", info.Mode())
//	errs = os.Chmod(path, 0660)
//	if errs != nil {
//		t.Error(errs)
//	}
//	info, errs = os.Stat(path)
//	if errs != nil {
//		t.Error(errs)
//	}
//	t.Log("after info : ", info.Mode())
//}
//func TestEqual(t *testing.T) {
//	fmt.Println(27 <= 26)
//}
//func TestTrim(t *testing.T) {
//	str := "NCO_NDI_B0006_20200722_001.txt"
//	str2 := strings.TrimSuffix(str, ".txt")
//	fmt.Println(str2)
//}
//
//type Photo struct {
//	ID   int64  `json:"id"`
//	Host string `json:"host"`
//	Url  string `json:"url"`
//}
//
//func StructToJSON(input interface{}) (output string) {
//	b, err := json.Marshal(input)
//	if err != nil {
//		fmt.Println(err)
//		output = ""
//		return
//	}
//	output = string(b)
//	return
//}
//func TestJSONEmail(t *testing.T) {
//	msg := "Dear {{.FIRST_NAME_USER}},\n\n" +
//		"Proses registrasi anda pada sistem Nexchief telah berhasil, dengan detail sebagai berikut:\n" +
//		"Your registration process in Nexchief system has been success, with detail below:\n\n" +
//		"Username\t:\t{{.USERNAME}}\n" +
//		"Email\t\t:\t{{.EMAIL}}\n" +
//		"Phone\t\t:\t{{.PHONE}}\n\n" +
//		"Untuk menyelesaikan proses registrasi anda, perlu dilakukan aktivasi email, silahkan akses link berikut untuk melakukan aktivasi.\n" +
//		"To finish your registration process, you need to activate your email with accessing this link activation.\n\n" +
//		"{{.ACTIVATION_LINK}}\n" +
//		"Atas perhatiannya, kami ucapkan terima kasih.\n" +
//		"Thank you for your attention.\n\n" +
//		"Hormat Kami,\n" +
//		"Regards,\n\n\n" +
//		"Nexsoft System"
//	fmt.Println(msg)
//}
//func TestPhotoList(t *testing.T) {
//	var photoList []Photo
//	out := StructToJSON("null")
//	t.Log(out)
//	t.Log(photoList)
//}
//func TestCronScheduler(t *testing.T) {
//	c := cron.New()
//	_, errs := c.AddFunc("0 0/1 * 1/1 *", func() {})
//	if errs != nil {
//		assert.FailNow(t, errs.Error(), "error format cron")
//	}
//}
//
//func TestList(t *testing.T) {
//	time.Sleep(time.Second * 5)
//	t.Log("test 1")
//}
//
//func TestListDua(t *testing.T) {
//	_, err := ChecksumFileWithMD5()
//	if err != nil {
//		assert.Error(t, err, err.Error())
//	}
//}
//
//func ChecksumFileWithMD5() (strChecksum string, err error) {
//	input := strings.NewReader("C:\\nco_file\\NCO_NDI_B0006_20200722_001.zip")
//	//input, _ := os.Open("C:\\nco_file\\NCO_NDI_B0006_20200722_001.zip")
//	hash := md5.New()
//	if _, err = io.Copy(hash, input); err != nil {
//		return "", err
//	}
//	sum := hash.Sum(nil)
//	strChecksum = fmt.Sprintf("%x\n", sum)
//	fmt.Println(strChecksum)
//	return
//}
//
//func TestNexsellerSyncTxtLog(t *testing.T) {
//	db := connectDB()
//	id, err := insertToDB(db, NexsellerSyncTxtLogNexchief{
//		NexchiefAccountId:      sql.NullInt64{Int64: 1},
//		MappingNexsellerId:     sql.NullInt64{Int64: 1},
//		TanggalFileGenerate:    sql.NullTime{},
//		NamaFileGenerate:       sql.NullString{String: "test_name"},
//		StatusFileGenerate:     sql.NullBool{},
//		TanggalSyncNexchief:    sql.NullTime{},
//		StatusSyncNexchief:     sql.NullBool{},
//		MessageSyncNexchief:    sql.NullString{},
//		TanggalNcoStartProcess: sql.NullTime{},
//		TanggalNcoEndProcess:   sql.NullTime{},
//		StatusProcessNexchief:  sql.NullBool{},
//		MessageProcessNexchief: sql.NullString{},
//		ProcessedBy:            sql.NullString{},
//		JobID:                  sql.NullString{},
//		ChecksumNco:            sql.NullString{},
//		ChecksumNcoLog:         sql.NullString{},
//		CreatedBy:              sql.NullInt64{},
//		CreatedAt:              sql.NullTime{},
//		CreatedClient:          sql.NullString{},
//		UpdatedBy:              sql.NullInt64{},
//		UpdatedAt:              sql.NullTime{},
//		UpdatedClient:          sql.NullString{},
//		Deleted:                sql.NullBool{},
//	})
//	if err != nil {
//		assert.FailNow(t, err.Error(), "error insert")
//	}
//	assert.Equal(t, true, id > 0, "id more than 0")
//}
//
//type NexsellerSyncTxtLogNexchief struct {
//	Id                     sql.NullInt64
//	UuidKey                sql.NullString
//	NexchiefAccountId      sql.NullInt64
//	MappingNexsellerId     sql.NullInt64
//	TanggalFileGenerate    sql.NullTime
//	NamaFileGenerate       sql.NullString
//	StatusFileGenerate     sql.NullBool
//	TanggalSyncNexchief    sql.NullTime
//	StatusSyncNexchief     sql.NullBool
//	MessageSyncNexchief    sql.NullString
//	TanggalNcoStartProcess sql.NullTime
//	TanggalNcoEndProcess   sql.NullTime
//	StatusProcessNexchief  sql.NullBool
//	MessageProcessNexchief sql.NullString
//	ProcessedBy            sql.NullString
//	JobID                  sql.NullString
//	ChecksumNco            sql.NullString
//	ChecksumNcoLog         sql.NullString
//	CreatedBy              sql.NullInt64
//	CreatedAt              sql.NullTime
//	CreatedClient          sql.NullString
//	UpdatedBy              sql.NullInt64
//	UpdatedAt              sql.NullTime
//	UpdatedClient          sql.NullString
//	Deleted                sql.NullBool
//}
//
//func TestUpdateNil(t *testing.T) {
//	db := connectDB()
//	id, err := insertToDB(db, NexsellerSyncTxtLogNexchief{
//		NexchiefAccountId:  sql.NullInt64{Int64: 15},
//		MappingNexsellerId: sql.NullInt64{Int64: 7, Valid: true},
//		Id:                 sql.NullInt64{Int64: 46},
//	})
//	if err != nil {
//		assert.FailNow(t, err.Error())
//	}
//	fmt.Println(id)
//}
//func insertToDB(db *sql.DB, userParam NexsellerSyncTxtLogNexchief) (id int64, err error) {
//	query :=
//		"SELECT id FROM nexchief.nexsoft_client_role_scope " +
//			"WHERE role_id = $1 AND  group_id = $2 AND  id = $3 "
//
//	param := []interface{}{
//		userParam.NexchiefAccountId.Int64, userParam.MappingNexsellerId, userParam.Id,
//	}
//
//	row := db.QueryRow(query, param...)
//
//	errs := row.Scan(&id)
//	if errs != nil && errs.Error() != "sql: no rows in result set" {
//		return
//	}
//	return
//}
//
//func TestGetDataBase(t *testing.T) {
//	db := connectDB()
//	var (
//		id       sql.NullInt64
//		roleName sql.NullString
//	)
//	query := "SELECT id, role_id FROM nexsoft_role WHERE id = 1 "
//	rows, errs := db.Query(query)
//	if errs != nil && errs != sql.ErrNoRows {
//		t.Error(errs)
//	}
//	for rows.Next() {
//
//	}
//	fmt.Printf("id : %d, name : %s\n", id.Int64, roleName.String)
//}
//
//func getTest(str string) func(t *testing.T) {
//	fmt.Println(str)
//	return func(t *testing.T) {
//		defer func() {
//			t.Log("defer")
//		}()
//
//		assert.FailNow(t, "true")
//	}
//}
//
//func TestPrint(t *testing.T) {
//	printTest("satu", "dua", "tiga")
//}
//func printTest(t ...string) {
//	var s string
//	for i := 0; i < len(t); i++ {
//		s += t[i]
//	}
//	fmt.Println(s)
//}
//
//type botol struct {
//	ID   int64
//	Name string
//}
//
//type laptop struct {
//	ID    int64
//	Mouse string
//}
//
//func getReflect(value interface{}) {
//	fmt.Println(util.StructToJSON(value))
//	//id := reflect.ValueOf(value)
//	//var ids int64
//	//temp := id.FieldByName("Id")
//	//if temp.Kind() == reflect.Int64 {
//	//	ids = temp.Int()
//	//	fmt.Println(ids)
//	//
//	//}
//	//fmt.Println(id.FieldByName("Name"))
//
//}
//
//type structReflects struct {
//	ID               int64  `json:"id"`
//	Name             string `json:"name"`
//	Address          string `json:"address"`
//	MappingNexseller string `json:"mapping_nexseller"`
//}
//
//func TestReflect(t *testing.T) {
//	val := reflect.ValueOf(structReflects{})
//	for i := 0; i < val.Type().NumField(); i++ {
//		fmt.Println(val.Type().Field(i).Tag.Get("json"))
//
//	}
//}
//
//type Struct1 struct {
//	Id   int64
//	Name string
//}
//
//func (s *Struct1) func1() {
//	fmt.Println("2 > ", &s.Name)
//	s.func2()
//}
//func (s *Struct1) func2() {
//	fmt.Println("3 > ", &s.Name)
//	s.func3()
//}
//func (s *Struct1) func3() {
//	fmt.Println("4 > ", &s.Name)
//}
//
//func TestPrints(t *testing.T) {
//	var s Struct1
//	//fmt.Println("0 > ", )
//	fmt.Println("1 > ", &s.Name)
//	s.func1()
//}
//
//type StructJsonTest struct {
//	Name    string `json:"name"`
//	Usia    int    `json:"usia"`
//	Address string `json:"address"`
//}
//type RepoJsonTest struct {
//	ID   sql.NullInt64
//	Info sql.NullString
//}
//
//func TestJsonTypeInsert(t *testing.T) {
//	db := connectDB()
//
//	js := StructJsonTest{
//		Name:    "Aji",
//		Usia:    23,
//		Address: "Jakarta",
//	}
//	repo := RepoJsonTest{
//		ID:   sql.NullInt64{Int64: 4},
//		Info: sql.NullString{String: StructToJSON(js)},
//	}
//	query :=
//		"INSERT INTO test_json " +
//			"(id, column_json) " +
//			"VALUES " +
//			"($1, $2) "
//
//	param := []interface{}{
//		repo.ID.Int64, repo.Info.String,
//	}
//	errs := db.QueryRow(query, param...).Err()
//	if errs != nil && errs.Error() != "sql: no rows in result set" {
//		fmt.Println("err >", errs)
//		return
//	}
//	return
//}
//
//func TestJsonTypeSelect(t *testing.T) {
//	db := connectDB()
//	var str sql.NullString
//	//var tmp StructJsonTest
//	query :=
//		"SELECT column_json ->> 'address' " +
//			"from test_json " +
//			"WHERE column_json ->> 'name' = $1"
//
//	param := []interface{}{
//		"Aji",
//	}
//	errs := db.QueryRow(query, param...).Scan(&str)
//	if errs != nil && errs.Error() != "sql: no rows in result set" {
//		fmt.Println("err >", errs)
//		return
//	}
//	fmt.Println(str.String)
//	//_ = json.Unmarshal([]byte(str.String), &tmp)
//	//fmt.Println("name > ", tmp.Name, "; usia >", tmp.Usia, "; address > ", tmp.Address)
//	return
//}
//
//type structTest1 struct {
//	ID    int64  `json:"id"`
//	Times string `json:"times"`
//}
//
//func TestParseDate(t *testing.T) {
//	s := StructToJSON(structTest1{
//		ID:    1,
//		Times: time.Now().Format("2006-01-02T15:04:05Z"),
//	})
//	fmt.Println(s)
//	//temp, err := time.Parse("02/01/2006", "09/02/2022")
//	//if err != nil {
//	//	t.Error(err)
//	//}
//	//fmt.Println(s)
//}
//
//func TestReplaceAll(t *testing.T) {
//	s := "50"
//	fl, errS := strconv.ParseFloat(s, 64)
//	if errS != nil {
//		t.Error(errS)
//	}
//	fmt.Println(fl)
//	str := "%200%"
//	str = strings.ReplaceAll(str, "%", "")
//	fmt.Println(str)
//}
//
//func validateFloat64(funcName string, maps map[string][]interface{}) (err errorModel.ErrorModel) {
//	//for i := 0; i < len(maps); i++ {
//	for fieldName, temp := range maps {
//		if len(temp) < 2 {
//			continue
//		}
//		strValue := temp[0].(string)
//		result := temp[1].(*float64)
//
//		if strValue == "" {
//			strValue = "0.0"
//		}
//		tempResult, errS := strconv.ParseFloat(strValue, 64)
//		if errS != nil {
//			err = errorModel.GenerateFormatFieldError("ProcessFileNCOService.go", funcName, fieldName)
//			return
//		}
//		*result = tempResult
//
//	}
//	//}
//	return
//}
//func TestValidateFloat64(t *testing.T) {
//	var f1, f2, f3 float64
//	willCheck := make(map[string][]interface{})
//	willCheck["f1"] = []interface{}{"5.0", &f1}
//	willCheck["f2"] = []interface{}{"3.0", &f2}
//	willCheck["f3"] = []interface{}{"f", &f3}
//	err := validateFloat64("Test", willCheck)
//	if err.Error != nil {
//		t.Error(err.Error, ",", err.FuncName)
//	}
//	fmt.Println(f1, f2, f3)
//}
//
//func TestValidateCasting(t *testing.T) {
//	//var s int64
//	//var d int
//	//s = 9223372036854775807
//	//d = 1234567890123456789
//	//fmt.Println(d)
//	//fmt.Println(int(s))
//	format := "2006-01-02T15:04:05Z"
//	strTime := "2022-03-02T15:06:07.425Z"
//	result, err := time.Parse(format, strTime)
//	if err != nil {
//		assert.FailNow(t, err.Error())
//	}
//	fmt.Println(result)
//}
//
//func TestTimeUnix(t *testing.T) {
//	//var i int64
//	i := 0
//	temp := strconv.Itoa(i)
//	if strings.HasSuffix(temp, "000") {
//		temp = strings.TrimSuffix(temp, "000")
//		i, _ = strconv.Atoi(temp)
//		fmt.Println(i)
//	}
//	times := time.Unix(int64(i), 0)
//	fmt.Println(times)
//
//}
//
//func TestReadExcel (t *testing.T){
//	type M map[string]interface{}
//
//	var data = []M{
//		{"Name": "Noval", "Gender": "male", "Age": 18},
//		{"Name": "Nabila", "Gender": "female", "Age": 12},
//		{"Name": "Yasa", "Gender": "male", "Age": 11},
//	}
//	xlsx := excelize.NewFile()
//
//	sheet1Name := "Sheet One"
//	//xlsx.NewSheet(sheet1Name)
//	xlsx.SetSheetName(xlsx.GetSheetName(0), sheet1Name)
//	err := xlsx.SetCellValue(sheet1Name, "A1", "Name")
//	err = xlsx.SetCellValue(sheet1Name, "B1", "Gender")
//	err = xlsx.SetCellValue(sheet1Name, "C1", "Age")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	err = xlsx.AutoFilter(sheet1Name, "A1", "C1", "")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	for i, each := range data {
//		_ = xlsx.SetCellValue(sheet1Name, fmt.Sprintf("A%d", i+2), each["Name"])
//		_ = xlsx.SetCellValue(sheet1Name, fmt.Sprintf("B%d", i+2), each["Gender"])
//		_ = xlsx.SetCellValue(sheet1Name, fmt.Sprintf("C%d", i+2), each["Age"])
//	}
//	err = xlsx.SaveAs("./file1.xlsx")
//	if err != nil {
//		fmt.Println(err)
//	}
//	_ = xlsx.Close()
//}
//
//func TestToJson(t *testing.T){
//	//type s struct {
//	//	Id int64 `json:"id"`
//	//	Name string `json:"name"`
//	//}
//	//var i interface{}
//	//i = s{
//	//	Id:   0,
//	//	Name: "Aji",
//	//}
//	//p := reflect.ValueOf(i).Elem()
//	//p.Set(reflect.Zero(p.Type()))
//	s := "{\"schema\":{\"type\":\"struct\",\"fields\":[{\"type\":\"string\",\"optional\":false,\"field\":\"principalID\"},{\"type\":\"string\",\"optional\":false,\"field\":\"customerGroupID\"},{\"type\":\"string\",\"optional\":true,\"default\":\"\",\"field\":\"customerGroupName\"},{\"type\":\"string\",\"optional\":true,\"default\":\"\",\"field\":\"customerGroupStdCode\"},{\"type\":\"string\",\"optional\":true,\"default\":\"\",\"field\":\"priceIDList\"},{\"type\":\"int64\",\"optional\":true,\"name\":\"io.debezium.time.Timestamp\",\"version\":1,\"field\":\"customerGroupCreated\"},{\"type\":\"int64\",\"optional\":true,\"name\":\"io.debezium.time.Timestamp\",\"version\":1,\"field\":\"customerGroupModified\"},{\"type\":\"string\",\"optional\":true,\"default\":\"\",\"field\":\"customerGroupModifiedBy\"}],\"optional\":false,\"name\":\"servermysql.debezium.customergroup.Value\"},\"payload\":{\"principalID\":\"PVM\",\"customerGroupID\":\"CSTGRP\",\"customerGroupName\":\"Group test po\",\"customerGroupStdCode\":\"\",\"priceIDList\":\"\",\"customerGroupCreated\":null,\"customerGroupModified\":null,\"customerGroupModifiedBy\":\"\"}}"
//	s = strings.TrimSpace(s)
//	fmt.Println(s)
//}