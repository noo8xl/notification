
# ------------ export PATH="$PATH:$(go env GOPATH)/bin"
gen:
	@protoc --go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	api/wallet.proto

build:
	rm -rf ./bin/ | GOFLAGS=-mod=mod GOOS=linux GOARCH=amd64 go build -o bin/notificator main.go

# ------------- run all tests 
test:
	go test ./... -v

migrate:
	migrate create -ext=sql -dir=/database/migrations/ -seq $(migrationName)