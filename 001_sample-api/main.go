package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func goodByeHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Good Bye...")
}

func helloWorldHandler(response http.ResponseWriter, request *http.Request) {
	b, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Println("Error reading body", err)

		http.Error(response, "Unable to read request body", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(response, "Hello %s", b)
}

func main() {
	http.HandleFunc("/good-bye", goodByeHandler)

	http.HandleFunc("/", helloWorldHandler)

	log.Println("Starting Server")
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}
