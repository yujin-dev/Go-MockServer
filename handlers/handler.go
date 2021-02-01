package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	_ "github.com/lib/pq"
)

func checkConn(db_info string) error {

	db, err := sql.Open("postgres", db_info)

	if err != nil {
		// panic(err)
		return err
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		// panic(err)
		return err
	}
	return nil
}

func getMain() bool {
	result := checkConn(MainServer)
	if result != nil {
		fmt.Println("Main Server Connection Failed")
		return false
	} else {
		fmt.Println("Main Server Successfully connected")
		return true
	}
}

func getSub() bool {
	result := checkConn(SubServer)
	if result != nil {
		fmt.Println("Sub Server Connection Failed")
		return false
	} else {
		fmt.Println("Sub Server Successfully connected")
		return true
	}
}

type serverStatus struct {
	Main bool
	Sub  bool
}

func GetStatus() serverStatus {
	conn1 := getMain()
	conn2 := getSub()
	status := serverStatus{conn1, conn2}
	// jsonBytes, err := json.Marshal(status)
	// if err != nil {
	// 	panic(err)
	// }
	return status
}

func GetConfig() map[string]interface{} {
	// json 파일 읽기
	jsonFile, err := os.Open("handlers/config.json")
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()
	// 디코딩: JSON 패키지가 디코딩 데이터를 저장할 수 있는 변수를 선언해야 함
	var info map[string]interface{}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	mErr := json.Unmarshal([]byte(byteValue), &info)
	if err != nil {
		panic(mErr)
	}
	return info
}
