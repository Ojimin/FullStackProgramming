package http

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
)

func Main_server() {
    /* Main function. */
    serverName := "localhost"
    serverPort := 8080
    addr := fmt.Sprintf("%s:%d", serverName, serverPort)
    http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodGet {
            do_Get(w, r)
        } else if r.Method == http.MethodPost {
            do_Post(w, r)
        }
    })
    err := http.ListenAndServe(addr, nil)
    if err != nil {
		fmt.Println("Fail to start : ", err)
	}
    fmt.Printf("## HTTP server started at http://%s:%d.\n", serverName, serverPort)
}

func print_http_request_detail(request *http.Request) {
    clientURL := request.RemoteAddr
    clientAddress, port, err := net.SplitHostPort(clientURL)
    if err!= nil {
        log.Println("ERROR : Don't split host and port", err)
        clientAddress = clientURL
        port = "error"
    }
    //Print HTTP request in detail.
    fmt.Printf("::Client address : %s\n", clientAddress)
    fmt.Printf("::Client port   : %s\n", port)
    fmt.Printf("::Request command  : %s\n", request.Method)
    fmt.Printf("::Request line     : %s %s %s\n", request.Method, request.URL.RequestURI(), request.Proto)
    query := request.URL.RawQuery
    if query != "" {
        query = "?" + query
    }
    fmt.Printf("::Request path     : %s%s\n", request.URL.Path, query)
    fmt.Printf("::Request version  : %s\n", request.Proto)       
}

func send_http_response_header(writer http.ResponseWriter) {
    writer.WriteHeader(http.StatusOK)
    writer.Header().Set("Content-type", "text/html; charset=utf-8")
}

func do_Get(writer http.ResponseWriter, request *http.Request) {
    // HTTP GET request handler.
    fmt.Println("## do_GET() activated.")

    //GET response header generation
    print_http_request_detail(request)
    send_http_response_header(writer)

    if strings.Contains(request.URL.RawQuery, "=") {
        routine := request.URL.RawQuery
        parameter := parameter_retrieval(routine)
        para1, _ := strconv.Atoi(parameter[0])
        para2, _ := strconv.Atoi(parameter[1])
        result := simple_calc(para1, para2)

        writer.Write([]byte("<html>"))
        get_response := fmt.Sprintf("GET request for calculation => %d x %d = %d", para1, para2, result)
        writer.Write([]byte(get_response))
        writer.Write([]byte("</html>"))
        fmt.Printf( "## GET request for calculation => %d x %d = %d.\n", para1, para2, result)
    } else {
        writer.Write([]byte("<html>"))
        writer.Write([]byte(fmt.Sprintf("<p>HTTP Request GET for Path: %s</p>",request.URL.Path )))
        writer.Write([]byte("</html>"))
        fmt.Printf("## GET request for directory => %s.\n",request.URL.Path)
    }
}

func do_Post(writer http.ResponseWriter, request *http.Request) {
    //HTTP POST request handler.
    fmt.Println("## do_POST() activated.")

    print_http_request_detail(request)
    send_http_response_header(writer)

    content_length, _ := strconv.Atoi(request.Header.Get("Content-Length")) //data size 얻기
    post_data := make([]byte, content_length)
    request.Body.Read(post_data)
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