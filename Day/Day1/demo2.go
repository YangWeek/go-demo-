package day1

import (
	"fmt"
	"net/http"
	"text/template"
)

func D1() {
	var ma map[string]int
	// 引用类型都要初实话空间
	ma["yang"] = 1
	fmt.Println(ma)
}

type UserInfo struct {
	Name   string
	Gender string
	Age    int
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 解析指定文件生成模板对象
	tmpl, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}

	// 利用给定数据渲染模板，并将结果写入w
	user := UserInfo{
		Name:   "yangdacun",
		Gender: "男",
		Age:    18,
	}
	tmpl.Execute(w, user)
	// 渲染后的结果写入到 http.ResponseWriter 对象
}

func D2() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil) // 监听和启动服务器
	if err != nil {
		fmt.Println("HTTP server failed,err:", err)
		return
	}
}
