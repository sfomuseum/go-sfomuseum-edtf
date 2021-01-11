GO=~/src/go/bin/go

test-from:
	$(GO) test -v -run TestEDTFStringFrom

test-to:
	$(GO) test -v -run TestToEDTFDate

cli:
	$(GO) build -mod vendor -o bin/to-edtf cmd/to-edtf/main.go
	$(GO) build -mod vendor -o bin/to-edtf-string cmd/to-edtf-string/main.go
	$(GO) build -mod vendor -o bin/server cmd/server/main.go
