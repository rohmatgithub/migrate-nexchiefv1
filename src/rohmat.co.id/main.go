package main

import (
	"archive/zip"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/OneOfOne/xxhash"
	_ "github.com/golang-migrate/migrate/database/postgres"
	//_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)
func doPrint(wg *sync.WaitGroup, message string) {
	defer wg.Done()
	fmt.Println(message)
}

func test(i int, str *string){
	if i == 50 {
		*str = "ada"
	}
}
func main() {
	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup
	var str string
	for i := 0; i < 100; i++ {
		var data = fmt.Sprintf("data %d", i)

		wg.Add(1)
		go doPrint(&wg, data)
		test(i, &str)
		//if str != "" {
		//	break
		//	//fmt.Println("ada")
		//}
	}
	wg.Wait()
	fmt.Println("start")
}

//func main() {
//	var PostgresqlUrl ="postgres://postgres:root@localhost:5432/Stagging?sslmode=disable&search_path=gorp_test"
//
//	m, err := migrate.New(
//		"file://C:/Nexchief2-dev/src/nexsoft.co.id/nexchief2/sql_migrations/sql_migrations_global",
//		PostgresqlUrl)
//	if err != nil {
//		log.Fatal(err)
//	}
//	if err = m.Up(); err != nil {
//		log.Fatal(err)
//	}
//}

func CheckSumWithMD5(content []byte) (checksum string) {
	hash := md5.New()
	hash.Write(content)
	hashInBytes := hash.Sum(nil)[:16]
	return hex.EncodeToString(hashInBytes)
}
func stringSaveDB(input string) bool {
	return input != ""
}

func CheckSumWithXXHASH(content []byte) (checksum string) {
	hash := xxhash.Checksum64(content)
	return strconv.Itoa(int(hash))
}

func tex(x *map[string]string) {
	//var temp = &x
	//fmt.Println("Address of variable x = ", &x)
	//fmt.Println("Address of variable temp = ", temp)

	if (*x)["1"] == "" {
		fmt.Println("kosong")
	}
	(*x)["1"] = "satu"
}
func IsSaveDBTime(input time.Time) bool {
	return !input.IsZero()
}

func strToTime2(input string) (temp time.Time, errorS error) {
	temp, errorS = time.Parse("15:04", input)
	return
}

const (
	MenuRoleConstanta      = "setup.role.role"
	MenuDataGroupConstanta = "setup.role.data-group"
	MenuUserConstanta      = "setup.role.user"
	LevelRootRole          = 13
	LevelSuperRole         = 14
)

func validatePermission(mustHavePermission string, permissionUser map[string][]string, levelRole int32) (isValid bool) {
	//funcName := "ValidatePermissionWithRole"
	//var isValid = false
	//var permissionAllowed string
	splitMustHavePermission := strings.Split(mustHavePermission, ":")
	menu := splitMustHavePermission[0]
	permission := splitMustHavePermission[1]

	isValid = false
	splitDotMenu := strings.Split(menu, ".")
	size := len(splitDotMenu)
	if size != 3 {
		return
	}

	if menu == MenuRoleConstanta || menu == MenuDataGroupConstanta || menu == MenuUserConstanta {
		if levelRole != LevelRootRole && levelRole != LevelSuperRole {
			return
		} else {
			temp := strings.Split(permission, "-")
			if len(temp) > 1 {
				if temp[1] == "own" {
					isValid = true
					return
				}
			} else {
				return
			}
		}
	}
	for size > 0 {
		menu = ""
		for i := 0; i < size; i++ {
			menu += splitDotMenu[i]
			if i < size-1 {
				menu += "."
			}
		}
		if permissionUser[menu] != nil {
			isValid = roleChecker(permission, permissionUser[menu])
			if isValid {
				break
			}
		}
		size--
	}
	return
}

func roleChecker(permissionNeed string, listPermission []string) bool {
	for i := 0; i < len(listPermission); i++ {
		if listPermission[i] == permissionNeed {
			return true
		}
		//if listPermission[i] == permissionNeed+"-own" {
		//	return true, listPermission[i]
		//}
	}
	return false
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "nexSOFT"
)

func GetListPermission(db *sql.DB) (result []string) {
	//funcName := "GetListPermission"
	query := "SELECT permission FROM " +
		"nexchief.permission "

	rows, errS := db.Query(query)
	if errS != nil {
		fmt.Println(errS)
		return
	}
	if rows != nil {
		defer func() {
			errS = rows.Close()
			if errS != nil {
				fmt.Println(errS)
			}
		}()
		for rows.Next() {
			var temp string
			errS = rows.Scan(
				&temp)
			if errS != nil {
				fmt.Println(errS)
				return
			}
			result = append(result, temp)
		}
	} else {
		fmt.Println(errS)
	}

	return
}
func ValidateStringContainInStringArray(listString []string, key string) bool {
	for i := 0; i < len(listString); i++ {
		if listString[i] == key {
			return true
		}
	}
	return false
}
func GenerateHashMapPermissionAndDataScope(listPermission []string, isRemoveDuplicate bool) (result map[string][]string) {
	sort.Slice(listPermission, func(i, j int) bool {
		if strings.Contains(listPermission[i], ".") && strings.Contains(listPermission[j], ".") {
			idxI := strings.Split(listPermission[i], ".")
			idxJ := strings.Split(listPermission[j], ".")
			if idxI[0] == idxJ[0] {
				return len(idxI) < len(idxJ)
			}
		} else if strings.Contains(listPermission[i], ".") {
			return false
		} else if strings.Contains(listPermission[j], ".") {
			return true
		}
		return listPermission[i] < listPermission[j]
	})

	result = make(map[string][]string)
	for i := 0; i < len(listPermission); i++ {
		var menu string
		permission := listPermission[i]
		//validationResult, _ := util.IsNexsoftPermissionStandardValid(permission)
		//if !validationResult {
		//	continue
		//}
		splitPermission := strings.Split(permission, ":")
		splitDotMenu := strings.Split(splitPermission[0], ".")
		sizeWithDot := len(splitDotMenu)
		if isRemoveDuplicate {
			var isAvailable = false
			for sizeWithDot > 0 {
				menu = ""
				for j := 0; j < sizeWithDot; j++ {
					menu += splitDotMenu[j]
					if j < sizeWithDot-1 {
						menu += "."
					}
				}
				if ValidateStringContainInStringArray(result[menu], splitPermission[1]) {
					isAvailable = true
					break
				}
				sizeWithDot--
			}

			if !isAvailable {
				result[splitPermission[0]] = append(result[splitPermission[0]], splitPermission[1])
			}
		} else {
			result[splitPermission[0]] = append(result[splitPermission[0]], splitPermission[1])
		}
	}

	return
}

func Unzip(src string, dest string) ([]string, error) {
	var filenames []string

	r, err := zip.OpenReader(src)
	if err != nil {
		return filenames, err
	}

	defer func() {
		if err = r.Close(); err != nil {
			log.Println("Failed to close file", err)
		}
	}()

	for _, f := range r.File {
		// Store filename/path for returning and using later on
		fpath := filepath.Join(dest, f.Name)

		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return filenames, fmt.Errorf("%s: illegal file path", fpath)
		}

		filenames = append(filenames, fpath)

		if f.FileInfo().IsDir() {
			// Make Folder
			err = os.MkdirAll(fpath, os.ModePerm)
			if err != nil {
				return filenames, err
			}
			continue
		}

		// Make File
		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return filenames, err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return filenames, err
		}

		rc, err := f.Open()
		if err != nil {
			return filenames, err
		}

		_, err = io.Copy(outFile, rc)

		// Close the file without defer to close before next iteration of loop
		err = outFile.Close()
		if err != nil {
			return filenames, err
		}
		err = rc.Close()
		if err != nil {
			return filenames, err
		}
	}
	return filenames, nil
}

type StructData struct {
	Count string `json:"count"`
	Data  []struct {
		SyncTxtID          int64  `json:"idSyncTXTLogNexchief"`
		BranchID           string `json:"branchID"`
		DateFileGenerate   int    `json:"tanggalFileGenerate"`
		PrincipalCode      string `json:"principalID"`
		Message            string `json:"messageSyncNexchief"`
		StatusFileGenerate bool   `json:"statusFileGenerate"`
		DateSyncNexchief   int    `json:"tanggalSyncNexchief"`
		DistributorCode    string `json:"distributorID"`
		CompanyCode        string `json:"companyID"`
		FileName           string `json:"namaFileGenerate"`
		StatusFileNExchief bool   `json:"statusSyncNexchief"`
	}
}

func ensureDir(dirName string) error {
	err := os.MkdirAll(dirName, 0755)
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
func GetDateContainer() (result string) {
	timeNow := time.Now()

	year, month, day := timeNow.Date()

	result += strconv.Itoa(year) + "/"
	if int(month) < 10 {
		result += "0"
	}
	result += strconv.Itoa(int(month)) + "/"

	if day < 10 {
		result += "0"
	}
	result += strconv.Itoa(day) + "/"

	return
}

func returnPointer() (s string) {
	var i int
	defer func() {
		fmt.Println("tes1")
	}()
	for i < 10 {
		s = s + strconv.Itoa(i) + ","
		i++
	}

	defer func() {
		fmt.Println("test2")
	}()
	return
}


func AESRohmat(key []byte, str string) []byte {
	text := []byte(str)
	// generate a new aes cipher using our 32 byte long key
	c, err := aes.NewCipher(key)
	// if there are any errors, handle them
	if err != nil {
		fmt.Println(err)
	}

	// gcm or Galois/Counter Mode, is a mode of operation
	// for symmetric key cryptographic block ciphers
	// - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	gcm, err := cipher.NewGCM(c)
	// if any error generating new GCM
	// handle them
	if err != nil {
		fmt.Println(err)
	}

	// creates a new byte array the size of the nonce
	// which must be passed to Seal
	nonce := make([]byte, gcm.NonceSize())
	// populates our nonce with a cryptographically secure
	// random sequence
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
	}

	// here we encrypt our text using the Seal function
	// Seal encrypts and authenticates plaintext, authenticates the
	// additional data and appends the result to dst, returning the updated
	// slice. The nonce must be NonceSize() bytes long and unique for all
	// time, for a given key.
	return gcm.Seal(nonce, nonce, text, nil)
}

func AESEncrypt(src string, key []byte, iv []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("key error1", err)
	}
	if src == "" {
		fmt.Println("plain content empty")
	}
	ecb := cipher.NewCBCEncrypter(block, iv)
	content := []byte(src)
	content = PKCS5Padding(content, block.BlockSize())
	crypted := make([]byte, len(content))
	ecb.CryptBlocks(crypted, content)

	return crypted
}

func AESDecrypt(crypt []byte, key []byte, iv []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("key error1", err)
	}
	if len(crypt) == 0 {
		fmt.Println("plain content empty")
	}
	ecb := cipher.NewCBCDecrypter(block, iv)
	decrypted := make([]byte, len(crypt))
	ecb.CryptBlocks(decrypted, crypt)

	//return PKCS5Trimming(decrypted) //option 1
	return PKCS5UnPadding(decrypted) //option2

}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}

func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	if length%64 == 0 {
		return src
	}
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}

func CheckScope(listScope map[string]interface{}, checkedScope []string) map[string]interface{} {
	result := make(map[string]interface{})
	numberOfAddedQuery := 0
	for key := range listScope {
		switch listScope[key].(type) {
		case []interface{}:
			for i := 0; i < len(checkedScope); i++ {
				scopeQuery, isExist := checkIsScopeContains(key, checkedScope[i])
				if isExist {
					numberOfAddedQuery++
					result = appendQuery(result, checkedScope[i], listScope[scopeQuery])
				}
			}
		case []map[string][]string:
			data := listScope[key].([]map[string][]string)
			isExist := isScopeExistInListHashmap(data, checkedScope)
			if isExist {
				numberOfAddedQuery++
				keySplit := strings.Split(key, ".")
				result = appendQuery(result, keySplit[len(keySplit)-1], listScope[key].([]map[string][]string))
			}
		}
	}
	if numberOfAddedQuery >= len(checkedScope) {
		return result
	} else {
		return nil
	}
}

func checkIsScopeContains(listScopeKey string, scope string) (string, bool) {
	splitMustHavePermission := strings.Split(scope, ":")
	menu := splitMustHavePermission[0]
	splitDotMenu := strings.Split(scope, ".")
	size := len(splitDotMenu)
	for size > 0 {
		menu = ""
		for i := 0; i < size; i++ {
			menu += splitDotMenu[i]
			if i < size-1 {
				menu += "."
			}
		}
		if menu == listScopeKey {
			return menu, true
		}
		size--
	}
	return "", false
}

func appendQuery(data map[string]interface{}, field string, value interface{}) map[string]interface{} {
	temp := data[field]
	switch value.(type) {
	case []map[string][]string:
		data = appendListScope(data, field, value.([]map[string][]string))
	case []interface{}:
		if temp == nil {
			if len(data) == 0 {
				data[field] = value
			} else {
				isFound := false
				for key := range data {
					switch data[key].(type) {
					case []map[string][]string:
						temp := data[key].([]map[string][]string)
						for i := 0; i < len(temp); i++ {
							for keyOnHash := range temp[i] {
								if keyOnHash == field {
									isFound = true
									break
								}
							}
						}
					}
				}
				if !isFound {
					data[field] = value
				}
			}
		}
	}
	return data
}

func isScopeExistInListHashmap(listScope []map[string][]string, listNeedScope []string) (isAllExist bool) {
	for i := 0; i < len(listScope); i++ {
		for key := range listScope[i] {
			isAllExist = false
			for j := 0; j < len(listNeedScope); j++ {
				_, isExist := checkIsScopeContains(key, listNeedScope[j])
				isAllExist = isExist || isAllExist
			}
			if !isAllExist {
				return false
			}
		}
	}
	return true
}

func appendListScope(data map[string]interface{}, field string, value []map[string][]string) map[string]interface{} {
	var fieldValue []map[string][]string

	if data[field] == nil {
		fieldValue = append(fieldValue, value...)
	} else {
		switch data[field].(type) {
		case []string:
			fieldValue = append(fieldValue, value...)
		case []map[string][]string:
			fieldValue = data[field].([]map[string][]string)
			fieldValue = append(fieldValue, value...)
		}
	}
	data[field] = fieldValue

	for i := 0; i < len(fieldValue); i++ {
		dataSlice := fieldValue[i]
		if dataSlice != nil {
			for key := range dataSlice {
				if data[key] != nil {
					delete(data, key)
				}
			}
		}
	}
	return data
}

func testPointer(str *string) {
	fmt.Println(str)
}
func TimeToString(time time.Time) string {
	return time.Format("15:04:05Z")
}
func caseUser(emailMdb string, emailAuth string, length int32) error {
	if length > 0 {
		//if resultMasterData[0].AuthUserID > 0 && resultMasterData[0].AuthUserID != resultAuth.UserID {
		//	return errors.New("err")
		//}

		if emailAuth != "" && emailMdb != emailAuth {
			return errors.New("err")
		}
	}
	return nil
}
func tesScope(scopeGeo interface{}) {
	//scopeGeo := authenticationModel.Data.Scope[constanta.ScopeGeoTree]
	if scopeGeo != nil {
		scopeAccess := scopeGeo.([]interface{})
		for i := 0; i < len(scopeAccess); i++ {
			if scopeAccess[i] == "ALL" {
				return
			}
			for j := i + 1; j < len(scopeAccess); j++ {
				if scopeAccess[i].(string) > scopeAccess[j].(string) {
					temp := scopeAccess[i]
					scopeAccess[i] = scopeAccess[j]
					scopeAccess[j] = temp
				}
			}
		}
		fmt.Println(scopeAccess)
	}
}

type tes struct {
	Str string
	Id int
}

func tesPointer(t *[]tes) {
	for i, t2 := range *t {
		t2.Id = i + 1
	}
}

//
//const (
//	name = "file"
//	path = "C:\\Nexchief2_config\\kumpulan_file"
//	fileName =	"NCO_ASW_B0005_20200722_001.txt"
//	usage = "file read"
//)
