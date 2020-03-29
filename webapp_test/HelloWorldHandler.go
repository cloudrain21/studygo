package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type hellowWorldRequest struct {
	Name string `json: "name"`
}

type hellowWorldResponse struct {
	Message string `json: "message"`
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	var request hellowWorldRequest
	err = json.Unmarshal(body, &request)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	//fmt.Println(request)

	response := hellowWorldResponse{Message: "Hello " + request.Name}
	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}

// This function is faster than the function above. (about 33% faster)
func hellowWorldHandlerDecoder(w http.ResponseWriter, r *http.Request) {
	var request hellowWorldRequest
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	response := hellowWorldResponse{Message: "Hello Decoder " + request.Name}
	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}

func main() {
	http.HandleFunc("/marshal", helloWorldHandler)
	http.HandleFunc("/decoder", hellowWorldHandlerDecoder)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
