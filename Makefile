GOMOD=$(shell test -f "go.work" && echo "readonly" || echo "vendor")

test-from:
	go test -v -run TestEDTFStringFrom

test-to:
	go test -v -run TestToEDTFDate

cli:
	go build -mod $(GOMOD) -ldflags="-s -w" -o bin/to-edtf cmd/to-edtf/main.go
	go build -mod $(GOMOD) -ldflags="-s -w" -o bin/to-edtf-string cmd/to-edtf-string/main.go
	go build -mod $(GOMOD) -ldflags="-s -w" -o bin/server cmd/server/main.go

server:
	go build -mod $(GOMOD) -ldflags="-s -w" -o bin/server cmd/server/main.go
	bin/server

lambda:
	@make lambda-server

lambda-server:
	if test -f bootstrap; then rm -f bootstrap; fi
	if test -f server.zip; then rm -f server.zip; fi
	GOARCH=arm64 GOOS=linux go build -mod $(GOMOD) -ldflags="-s -w" -tags lambda.norpc -o bootstrap cmd/server/main.go
	zip server.zip bootstrap
	rm -f bootstrap
