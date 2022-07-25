APP=./bin/app

test:
	go test ./internal/handlers -cover && go test ./internal/utils/hash_func_tests -cover
build:
	go build -o $(APP) ./cmd/appiserver/main.go
env:
	. db_init.sh
