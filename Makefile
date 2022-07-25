APP=./bin/app

build:
	go build -o $(APP) ./cmd/apiserver/main.go
test:
	go test ./internal/handlers -cover && go test ./internal/utils/hash_func_tests -cover
run_pg:
	$(APP)
run_im:
	$(APP) -storage cache
env:
	.	./db_init.sh
