APP=./bin/app
IMG_NAME = ozon
IMG_VERSION = latest
SHELL := /bin/bash

build:
	go build -o $(APP) ./cmd/apiserver/main.go
test:
	go test ./internal/handlers -cover && go test ./internal/utils/hash_func_tests -cover
run_pg:
	$(APP) -storage pg
run_im:
	$(APP) -storage cache
docker_build_im:
	docker build -t $(IMG_NAME):$(IMG_VERSION) .
docker_run_im:
	docker run -d -p 8080:8080 $(IMG_NAME):$(IMG_VERSION)
docker_build_pg:
	docker-compose build
docker_run_pg:
	docker-compose up
post: 
	curl --request POST --data '{"url" : "ozon"}' http://localhost:8080/
get: 
	curl --request GET --data '{"url" : "Li0QUvKTcT"}' http://localhost:8080/

# source ./db_init.sh $(sed -n 4p configs/pg_config)