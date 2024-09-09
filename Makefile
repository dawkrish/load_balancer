.PHONY:lb be
lb: 
	go build -o lb ./cmd/load_balancer/.

be: 
	go build -o be ./cmd/backend_server/.

clean:
	rm be lb

8080:
	python3 -m http.server 8080 --directory server8080

8081:
	python3 -m http.server 8081 --directory server8081

8082:
	python3 -m http.server 8082 --directory server8082