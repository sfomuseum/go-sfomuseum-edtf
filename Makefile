test-from:
	go test -v -run TestEDTFStringFrom

test-to:
	go test -v -run TestToEDTFDate

cli:
	go build -mod vendor -o bin/to-edtf cmd/to-edtf/main.go
	go build -mod vendor -o bin/to-edtf-string cmd/to-edtf-string/main.go
	go build -mod vendor -o bin/server cmd/server/main.go

server:
	go build -mod vendor -o bin/server cmd/server/main.go
	bin/server

lambda:
	@make lambda-server

lambda-server:
	if test -f main; then rm -f main; fi
	if test -f server.zip; then rm -f server.zip; fi
	GOOS=linux go build -mod vendor -o main cmd/server/main.go
	zip server.zip main
	rm -f main
