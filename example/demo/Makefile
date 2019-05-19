build: api 

dep:
	dep ensure

api: test
	docker build -f cmd/api/Dockerfile -t bosonic-code/demo-mock-api .

test: dep 
	go test -race `go list ./... | grep -v integration-tests`
	
itest: build
	cd integration-tests && \
	./run.sh	

