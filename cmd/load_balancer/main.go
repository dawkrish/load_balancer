package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type ServerConfig struct {
	requestID       int
	availableServer string
	servers         map[string][]int
	numOfServers    int
}

func NewServerConfig() *ServerConfig {
	return &ServerConfig{
		requestID:       0,
		availableServer: "",
		servers: map[string][]int{
			"8080": {},
			"8081": {},
			"8082": {},
		},
		numOfServers: 3,
	}
}

func main() {
	srvCfg := NewServerConfig()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		resp := fmt.Sprintf("Received request from %v\n", r.Host) +
			fmt.Sprintf("%v %v %v\n", r.Method, r.URL, r.Proto) +
			fmt.Sprintf("Host: %v\n", r.Host) +
			fmt.Sprintf("User-Agent: %v\n", r.Header.Get("User-Agent")) +
			fmt.Sprintf("Accept: %v\n", r.Header.Get("Accept"))

		srvCfg.roundRobin()
		cSrv := srvCfg.availableServer

		body := sendReq(r, cSrv)
		resp += "\n" + string(body)
		fmt.Fprint(w, resp)
	})

	log.Fatal(http.ListenAndServe(":7811", nil))
}

func sendReq(r *http.Request, port string) string {
	c := http.Client{}
	newReq, err := http.NewRequest(r.Method, "http://localhost:"+port, r.Body)
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

func (c *ServerConfig) roundRobin() {
	c.requestID++
	minK, minV := "", (c.requestID-1)/c.numOfServers
	for k, v := range c.servers {
		if len(v) <= minV {
			minK = k
			break
		}
	}
	c.servers[minK] = append(c.servers[minK], c.requestID)
	c.availableServer = minK
}
