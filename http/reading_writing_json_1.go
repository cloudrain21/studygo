package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type helloWorldResponse struct {
	// 출력 필드를 "msg" 로 바꾼다.
	Message string `json:"msg""`
	// 이 필드를 출력하지 않는다.
	Author string `json:"-"`
	// 값이 비어있으면 필드를 출력하지 않는다.
	Date string `json:", omitempty"`
	// 출력을 문자열로 변환하고 이름을 "id"로 바꾼다.
	Id int `json:"id, string"`
}

func main() {
	port := 8080

	http.HandleFunc("/helloworld", helloWorldHandler)

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	response := helloWorldResponse{Message: "HelloWorld", Author: "Rain", Date: "", Id: 100}
	data, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		panic("Ooops")
	}
	fmt.Fprintf(w, string(data))
}
