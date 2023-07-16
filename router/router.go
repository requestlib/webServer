package router

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"webServer/log/logger"
)

// 注册接收消息体
type RegistInfo struct {
	PhoneNumber string `json:"phone_number"`
	Nick        string `json:"nick"`
	Password    string `json:"password"`
}

// 登录接收消息体
type LoginInfo struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

// 【临时】 数据库格式
const (
	PhoneNumberIdx int16 = iota
	NickIdx
	PasswordIdx
)

type PathRouter struct {
}

func (p *PathRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayhelloName(w, r)
		return
	} else if r.URL.Path == "/regist" {
		registUser(w, r)
		return
	} else if r.URL.Path == "/login" {
		loginUser(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

// 注册用户
// 暂时先写入csv文件作为数据库 TODO:改成mysql数据库
func registUser(w http.ResponseWriter, r *http.Request) {
	// 解析body
	var registInfo RegistInfo
	err := json.NewDecoder(r.Body).Decode(&registInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	logger.Writer.Info("body: %+v", registInfo)
	fmt.Fprint(w, "regist success")
	// 用户注册信息写入数据库 (暂时先用csv文件)
	csvPath := "/root/liweiran/project/webServer/data/user_info.csv"
	csvObj, err := os.OpenFile(csvPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logger.Writer.Error("%v", err)
	}
	defer csvObj.Close()
	csvWriter := csv.NewWriter(csvObj)
	record := []string{registInfo.PhoneNumber, registInfo.Nick, registInfo.Password}
	err = csvWriter.Write(record)
	if err != nil {
		logger.Writer.Error("%v", err)
	}
	csvWriter.Flush()
}

// 用户登录
// 暂时先写入csv文件作为数据库 TODO:改成mysql数据库
func loginUser(w http.ResponseWriter, r *http.Request) {
	// 解析body
	var loginInfo LoginInfo
	err := json.NewDecoder(r.Body).Decode(&loginInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	logger.Writer.Info("body: %+v", loginInfo)
	// 数据库查询是否登录成功 (暂时先用csv文件)
	csvPath := "/root/liweiran/project/webServer/data/user_info.csv"
	csvObj, err := os.Open(csvPath)
	if err != nil {
		logger.Writer.Error("open csv file: %v", err)
	}
	defer csvObj.Close()

	reader := csv.NewReader(csvObj)
	records, err := reader.ReadAll()
	if err != nil {
		logger.Writer.Error("read csv file: %v", err)
	}
	var is_match bool = false //用户是否登录成功
	for _, record := range records {
		phoneNumber, password := record[PhoneNumberIdx], record[PasswordIdx]
		if loginInfo.Password == password && loginInfo.PhoneNumber == phoneNumber {
			is_match = true
			break
		}
	}
	if is_match {
		fmt.Fprint(w, "success")
	} else {
		fmt.Fprint(w, "fail")
	}

}

// 默认
func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数
	fmt.Println(r.Form) //格式
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("value:", strings.Join(v, ""))
	}
	fmt.Fprint(w, "Hello 李夏雨 : )")
}
