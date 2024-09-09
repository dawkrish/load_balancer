.PHONY:lb be
lb: 
	go build -o lb ./cmd/load_balancer/.

be: 
	go build -o be ./cmd/backend_server/.

clean:
	rm be lb