build: server 

dep:
	dep ensure

server: test
	docker build -f cmd/server/Dockerfile -t bosonic-code/api-mock .

test: dep 
	go test -race `go list ./... | grep -v example`
	