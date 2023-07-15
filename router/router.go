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

type RegistInfo struct {
	PhoneNumber string `json:"phone_number"`
	Nick        string `json:"nick"`
	Password    string `json:"password"`
}

type PathRouter struct {
}

func (p *PathRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayhelloName(w, r)
		return
	} else if r.URL.Path == "/regist" {
		registUser(w, r)
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
