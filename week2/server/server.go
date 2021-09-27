package server

import (
	"fmt"
	"net/http"
	"os"
)

func healthz(w http.ResponseWriter, r *http.Request) {
	// 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
	w.Header().Add("VERSION", os.Getenv("VERSION"))

	// 接收客户端 request，并将 request 中带的 header 写入 response header
	for key, values := range r.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	// Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
	fmt.Println(r.RemoteAddr, r.Header)

	switch r.URL.String() {
	case "/healthz": // 当访问 localhost/healthz 时，应返回 200
		fmt.Fprintf(w, "200")
	default:
		fmt.Fprintf(w, "server")
	}

}

func Server(port string) {
	http.HandleFunc("/", healthz)
	http.ListenAndServe(port, nil)
}
