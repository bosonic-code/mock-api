build: server 

proto: 
	protoc -I mocker mocker/*.proto --go_out=plugins=grpc:mocker

dep:
	dep ensure

server: test
	docker build -f cmd/server/Dockerfile -t bosonic/mock-api .

test: dep proto
	go test -race `go list ./... | grep -v example`
	