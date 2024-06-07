package http

import (
	"fmt"
	"io" //응답 본문을 읽음
	"net/http"
	"net/url" //url 인코딩 설정
)

func Http_web_client() {
	fmt.Println("## HTTP client started.")

	fmt.Println("## GET request for http://localhost:8080/temp/")
	http_request, _ := http.Get("http://localhost:8080/temp/")
	body, _ := io.ReadAll(http_request.Body)
	fmt.Println("## GET response [start]")
	fmt.Println(string(body))
	fmt.Println("## GET response [end]")
	http_request.Body.Close()

	fmt.Println("## GET request for http://localhost:8080/?var1=9&var2=9")
	http_request, _ = http.Get("http://localhost:8080/?var1=9&var2=9")
	body, _ = io.ReadAll(http_request.Body)
	fmt.Println("## GET response [start]")
	fmt.Println(string(body))
	fmt.Println("## GET response [end]")
	http_request.Body.Close()

	fmt.Println("## POST request for http://localhost:8080/ with var1 is 9 and var2 is 9")
	data := url.Values{}
	data.Set("var1", "9")
	data.Set("var2", "9")
	http_request, _ = http.PostForm("http://localhost:8080/", data)
	body, _ = io.ReadAll(http_request.Body)
	fmt.Println("## POST response [start]")
	fmt.Println(string(body))
	fmt.Println("## POST response [end]")
	http_request.Body.Close()

	fmt.Println("## HTTP client completed.")
}