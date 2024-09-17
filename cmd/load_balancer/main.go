package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type ServerConfig struct {
	requestID       int
	availableServer string
	servers         map[string]Server
	numOfServers    int
}

type Server struct {
	requests []int
	healthy  bool
}

func NewServerConfig() *ServerConfig {
	return &ServerConfig{
		requestID:       0,
		availableServer: "",
		servers: map[string]Server{
			"8080": {healthy: true},
			"8081": {healthy: true},
			"8082": {healthy: true},
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

		srvCfg.healthCheck()
		srvCfg.roundRobin()
		fmt.Println(srvCfg)

		cSrv := srvCfg.availableServer
		body := sendReq(r, cSrv)
		resp += "\n" + string(body)
		fmt.Fprint(w, resp)
	})

	log.Fatal(http.ListenAndServe(":7811", nil))
}

// Only healthy servers can come to this!
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

// Apply for healthy servers only
func (cfg *ServerConfig) roundRobin() {
	if cfg.numOfServers <= 0 {
		log.Println("no server is healthy:/")
		os.Exit(1)
	}

	cfg.requestID++
	minK, minV := "", (cfg.requestID-1)/cfg.numOfServers

	for k, v := range cfg.servers {
		if v.healthy == false {
			continue
		}
		if len(v.requests) <= minV {
			minK = k
			break
		}
	}

	reqs := append(cfg.servers[minK].requests, cfg.requestID)
	s := Server{requests: reqs, healthy: true}
	cfg.servers[minK] = s

	cfg.availableServer = minK
}

func (cfg *ServerConfig) healthCheck() {
	client := http.Client{}
	var countHealthy = 0

	for port, srv := range cfg.servers {
		resp, err := client.Get("http://localhost:" + port)
		log.Println(resp, err)

		if err != nil || resp.StatusCode != 200 {
			srv.healthy = false
			cfg.servers[port] = srv
		} else {
			srv.healthy = true
			cfg.servers[port] = srv
			countHealthy++
		}
	}
	cfg.numOfServers = countHealthy
}
