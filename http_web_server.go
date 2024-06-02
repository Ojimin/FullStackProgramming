// 핸들러 함수 : http 메서드 요청 처리 & 요청 세부 정보, 응답 헤더 등 설정
// get, post 요청 처리 함수
// 계산 처리
// 매개변수 추출 함수
// 메인 함수 - 요청 핸들링 및
// get요청에서 ? 뒤에 숫자들 받아 곱하기 해주고 post 요청에서 똑같이 calculation 해줄것
package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func main() {
    var addr string = "localhost:8080"
    http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodGet {
            do_Get(w, r)
        } else if r.Method == http.MethodPost {
            do_Post(w, r)
        }
    })
    err := http.ListenAndServe(addr, nil)
    if err != nil {
		fmt.Println("Failed to start server:", err)
	}
    fmt.Printf("## HTTP server started at http://%s.\n", addr)
    http_web_client()
}

func http_web_server() {
	http.HandleFunc("/hello", func(w http.ResponseWriter,  req *http.Request) {
        w.Write([]byte("Hello World"))
    })
    http.ListenAndServe(":8080", nil)
}
// 수정 필요
func print_http_request_detail(request *http.Request) {
    //Print HTTP request in detail.
    fmt.Printf("::Client address : %s\n", request.RemoteAddr)
    // fmt.Printf("::Client port   : %s\n", request.)
    fmt.Printf("::Request command  : %s\n")
    fmt.Printf("::Request line     : %s\n", request.RequestURI)
    fmt.Printf("::Request path     : %s\n", request.URL.Path)
    fmt.Printf("::Request version  : %s\n", request.Proto)       
}
func send_http_response_header(writer http.ResponseWriter) {
    //200도 표시
    writer.Header().Set("Content-type", "text/html")
}

func do_Get(writer http.ResponseWriter, request *http.Request) {
    fmt.Println("## do_GET() activated.")
    print_http_request_detail(request)
    send_http_response_header(writer)
    if strings.Contains(request.URL.RawQuery, "?") {
        routine := strings.Split(request.URL.RawQuery, "?")[1]
        parameter := parameter_retrieval(routine)
        para1, _ := strconv.Atoi(parameter[0])
        para2, _ := strconv.Atoi(parameter[1])
        result := simple_calc(para1, para2)

        fmt.Fprintf(writer, "<html>")
        fmt.Fprintf(writer, "GET request for calculation => %d x %d = %d", para1, para2, result)
        fmt.Fprintf(writer, "</html>")
        fmt.Printf( "## GET request for calculation => %d x %d = %d.", para1, para2, result)
    } else {
        fmt.Fprintf(writer, "<html>")
        fmt.Fprintf(writer, "<p>HTTP Request GET for Path: %s</p>",request.URL.Path)
        fmt.Fprintf(writer, "</html>")
        fmt.Printf("## GET request for directory => %s.",request.URL.Path)
    }
}

func do_Post(writer http.ResponseWriter, request *http.Request) {
    fmt.Println("## do_POST() activated.")

    print_http_request_detail(request)
    send_http_response_header(writer)

    post_data, _ := io.ReadAll(request.Body)
    defer request.Body.Close()

    parameter := parameter_retrieval(string(post_data))
    para1, _ := strconv.Atoi(parameter[0])
    para2, _ := strconv.Atoi(parameter[1])
    result := simple_calc(para1, para2)

    post_response := fmt.Sprintf("POST request for calculation => %d x %d = %d", para1, para2, result)
    writer.Write([]byte(post_response))

    fmt.Printf("## POST request data => %s\n.", string(post_data))
    fmt.Printf("## POST request for calculation => %d x %d = %d.\n", para1, para2, result)
}    


func log_message() {
    return 
}

func simple_calc(para1 , para2 int) int {
    return para1 * para2
}

func parameter_retrieval(msg string) []string{
    var fields []string = strings.Split(msg, "&")
    result := []string{
        strings.Split(fields[0],"=")[1],
        strings.Split(fields[1],"=")[1],
    }
    return result
}   