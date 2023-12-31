package main

import (
	"log"
	"net/http"
	"webServer/router"
)

func main() {
	router := &router.PathRouter{}
	err := http.ListenAndServe(":80", router) //设置监听端口
	if err != nil {
		log.Fatal("ListenAndServer http: ", err)
	}
}
