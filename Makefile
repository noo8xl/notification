
build:
	rm -rf ./bin/ | GOFLAGS=-mod=mod GOOS=linux GOARCH=amd64 go build -o bin/notification-api main.go

# ------------- run all tests
test:
	go test ./... -v
