package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func RestClient() {
	// Reads a non registered member : error-case
	r := getMember("http://127.0.0.1:5000/membership_api/0001")
	printResponse("#1",r)

	// Creates a new registered member : non-error case
	r= postMember("http://127.0.0.1:5000/membership_api/0001", "0001=apple")
	printResponse("#2",r)


	// Reads a registered member : non-error case
	r = getMember("http://127.0.0.1:5000/membership_api/0001")
	printResponse("#3",r)

	// Creates an already registered member : error case
	r=postMember("http://127.0.0.1:5000/membership_api/0001","0001=xpple")
	printResponse("#4",r)


	// Updates a non registered member : error case
	r=putMember("http://127.0.0.1:5000/membership_api/0002","0002=xrange")
	printResponse("#5",r)

	// Updates a registered member : non-error case
	postMember("http://127.0.0.1:5000/membership_api/0002","0002=xrange")
	r=putMember("http://127.0.0.1:5000/membership_api/0002","0002=orange")
	printResponse("#6",r)

	// Delete a registered member : non-error case
	r=deleteMember("http://127.0.0.1:5000/membership_api/0001")
	printResponse("#7",r)

	// Delete a non registered member : non-error case
	r=deleteMember("http://127.0.0.1:5000/membership_api/0001")
	printResponse("#8",r)

}

func getMember(url string) *http.Response {
	r, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	return r
}

func postMember(url string,  data string) *http.Response {
	r, err := http.Post(url,"application/x-www-form-urlencoded", strings.NewReader(data))
	if err != nil {
		panic(err)
	}
	return r
}

func putMember(url string, data string) *http.Response {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, url, strings.NewReader(data))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	r, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	return r
}

func deleteMember(url string) *http.Response {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		panic(err)
	}
	r, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	return r
}

func parseResponseResult(bodyBytes []byte, key string) interface{} {
	var result map[string]interface{}
	if err := json.Unmarshal(bodyBytes, &result); err != nil {
		panic(err)
	}
	value, exists := result[key]
	if !exists {
		return nil
	}
	return value
}

func printResponse(tag string, r *http.Response) {
	bodyBytes, _ := io.ReadAll(r.Body)
	r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	urlParts := strings.Split(r.Request.URL.Path, "/")
	memberID := urlParts[len(urlParts)-1] //member id 추출
	fmt.Printf("%s Code: %d >> JSON: %s >> JSON Result: %s\n", tag, r.StatusCode, string(bodyBytes), parseResponseResult(bodyBytes, memberID))
}