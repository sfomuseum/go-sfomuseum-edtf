test-from:
	go test -v -run TestEDTFStringFrom

test-to:
	go test -v -run TestToEDTFDate

cli:
	go build -mod vendor -o bin/to-edtf cmd/to-edtf/main.go
	go build -mod vendor -o bin/to-edtf-string cmd/to-edtf-string/main.go
	go build -mod vendor -o bin/server cmd/server/main.go
