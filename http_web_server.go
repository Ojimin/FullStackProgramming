// get요청에서 ? 뒤에 숫자들 받아 곱하기 해주고 post 요청에서 똑같이 calculation 해줄것
package main

import (
	"net/http"
)

func http_web_server() {
	http.HandleFunc("/hello", func(w http.ResponseWriter,  req *http.Request) {
        w.Write([]byte("Hello World"))
    })
 
    http.ListenAndServe(":8080", nil)
}