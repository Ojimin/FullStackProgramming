package main

import (
	"Webserver_Go/http"
	// "Webserver_Go/json_example"
	// "Webserver_Go/rest"
)

/*
go init mod
go mod tidy
각 라인 별로 주석을 해제하시면서 실행시키면 됩니다
*/
func main() {
	/*
	http_web_server & client
	*/
	http.Main_server()
	// http.Http_web_client()
	/*
	json_example
	*/
	// json_example.JsonToGo()
	// json_example.GoToJson()
	// json_example.JsonToGoPrint()
	// json_example.JsonStringToGo()
	/*
	rest server & client
	*/
	// rest.RestServer()
	// rest.RestClient()	
}