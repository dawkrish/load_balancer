package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		resp := fmt.Sprintf("Received request from %v\n", r.Host) +
			fmt.Sprintf("%v %v %v\n", r.Method, r.URL, r.Proto) +
			fmt.Sprintf("Host: %v\n", r.Host) +
			fmt.Sprintf("User-Agent: %v\n", r.Header.Get("User-Agent")) +
			fmt.Sprintf("Accept: %v\n", r.Header.Get("Accept"))

		respFromServer, err := http.Get("http://localhost:5000/")
		if err != nil {
			log.Println(err)
		}

		body, err := ioutil.ReadAll(respFromServer.Body)

		if err != nil {
			log.Fatal(err)
		}

		resp += "\n" + string(body)

		fmt.Fprintf(w, resp)

	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
