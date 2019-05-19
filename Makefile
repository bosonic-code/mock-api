build: server 

proto: 
	protoc -I internal/proto internal/proto/*.proto --go_out=plugins=grpc:internal/proto

dep:
	dep ensure

server: test
	docker build -f cmd/server/Dockerfile -t bosonic/mock-api .

test: dep proto
	go test -race `go list ./... | grep -v example`
	