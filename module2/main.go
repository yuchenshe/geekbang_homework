/*
(本人功底太差，借鉴了老师和网上的代码，尝试读懂中...)
作业:编写一个 HTTP 服务器
1.接收客户端 request，并将 request 中带的 header 写入 response header
2.读取当前系统的环境变量中的 VERSION 配置，并写入 response header
3.Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
4.当访问 {url}/healthz 时，应返回200
*/
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	HttpServerStart(8080)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func HttpServerStart(port int) {
	http.HandleFunc("/", httpAccessFunc)
	http.HandleFunc("healthz", healthzFunc)

	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		log.Fatal(err)
	}
}

func healthzFunc(w http.ResponseWriter, r *http.Request) {
	HealthzCode := "200"
	w.Write([]byte(HealthzCode))
}

func httpAccessFunc(w http.ResponseWriter, r *http.Request) {
	if len(r.Header) > 0 {
		for k, v := range r.Header {
			log.Printf("%s=%s", k, v[0])

			//request header写入response header
			w.Header().Set(k, v[0])
		}
	}

	r.ParseForm() //解析所有请求数据，否则无法获取数据
	if len(r.Form) > 0 {
		for k, v := range r.Form {
			log.Printf("%s=%s", k, v[0])
		}
	}

	os.Setenv("VERSION", "v1.0")
	name := os.Getenv("VERSION")
	fmt.Println("VERSION: ", name)

	//获取Client IP
	clientip := getCurrentIP(r)
	log.Printf("Success! Response code: %d", 200)
	log.Printf("Success! clientip: %s", clientip)

	fmt.Printf("http Status Code ===>>%v\n", http.StatusOK)
	log.Println(http.StatusOK)

	//response响应
	w.WriteHeader(http.StatusOK)

	w.Write([]byte("Server Access,Success!"))
}

func getCurrentIP(r *http.Request) string {
	ip := r.Header.Get("X-Real-IP")
	if ip == "" {
		ip = strings.Split(r.RemoteAddr, ":")[0]
	}
	return ip
}
