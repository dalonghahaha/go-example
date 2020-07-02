package main

import (
	"fmt"
	"net/http"
)

func main() {
	//注册处理函数，用户连接自动调用指定的处理函数
	http.HandleFunc("/get", HandConnGet)
	http.HandleFunc("/post", HandConnPost)
	//监听http端口
	http.ListenAndServe(":8000", nil)
}

func HandConnGet(w http.ResponseWriter, r *http.Request) {
	//用户请求地址
	fmt.Println(r.URL)
	//请求头
	fmt.Println(r.Header)
	//返回数据
	w.Write([]byte("hello get"))
}

func HandConnPost(w http.ResponseWriter, r *http.Request) {
	//请求请求内容
	fmt.Println(r.PostFormValue("id"))
	fmt.Println(r.PostFormValue("name"))
	//返回数据
	w.Write([]byte("hello post"))
}