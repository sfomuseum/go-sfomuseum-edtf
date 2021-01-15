GO=~/src/go/bin/go

test-from:
	$(GO) test -v -run TestEDTFStringFrom

test-to:
	$(GO) test -v -run TestToEDTFDate

cli:
	$(GO) build -mod vendor -o bin/to-edtf cmd/to-edtf/main.go
	$(GO) build -mod vendor -o bin/to-edtf-string cmd/to-edtf-string/main.go
	$(GO) build -mod vendor -o bin/server cmd/server/main.go

server:
	$(GO) build -mod vendor -o bin/server cmd/server/main.go
	bin/server

lambda:
	@make lambda-server

lambda-server:
	if test -f main; then rm -f main; fi
	if test -f server.zip; then rm -f server.zip; fi
	GOOS=linux $(GO) build -mod vendor -o main cmd/server/main.go
	zip server.zip main
	rm -f main
