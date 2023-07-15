package router

import (
	"fmt"
	"net/http"
	"strings"
	"webServer/log/logger"
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
	}
	http.NotFound(w, r)
	return
}

// 注册用户
// 暂时先写入csv文件作为数据库 TODO:改成mysql数据库
func registUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	logger.Writer.Info("%v", r)
	fmt.Fprint(w, "注册成功！")
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
