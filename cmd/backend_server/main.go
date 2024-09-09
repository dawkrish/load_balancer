package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		metadata := fmt.Sprintf("Received request from %v\n", r.Host) +
			fmt.Sprintf("%v %v %v\n", r.Method, r.URL, r.Proto) +
			fmt.Sprintf("Host: %v\n", r.Host) +
			fmt.Sprintf("User-Agent: %v\n", r.Header.Get("User-Agent")) +
			fmt.Sprintf("Accept: %v\n", r.Header.Get("Accept"))
		fmt.Println(metadata)

		fmt.Fprintf(w, "Hello From Backend Server")
	})

	log.Fatal(http.ListenAndServe(":7842", nil))
}
