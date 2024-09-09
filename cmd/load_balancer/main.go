package main

import (
	"fmt"
	"io"
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

		body := sendReq(r, "7842")
		resp += "\n" + string(body)
		fmt.Fprint(w, resp)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}

func sendReq(r *http.Request, port string) string {
	c := http.Client{}
	newReq, err := http.NewRequest(r.Method, "http://localhost:" + port, r.Body)
	if err != nil {
		log.Println(err)
	}
	resp, err := c.Do(newReq)
	if err != nil {
		log.Println(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(body)
}
