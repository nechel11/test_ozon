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
docker_build:
	docker build -t $(IMG_NAME):$(IMG_VERSION) .
docker_run_im:
	docker run -d -p 8080:8080 $(IMG_NAME):$(IMG_VERSION)
post: 
	curl --request POST --data '{"url" : "ozon"}' http://localhost:8080/
get: 
	curl --request GET --data '{"url" : "Li0QUvKTcT"}' http://localhost:8080/