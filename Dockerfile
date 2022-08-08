FROM golang:1.18.3

RUN mkdir -p ozon
COPY ./ /ozon
WORKDIR /ozon

RUN apt-get update
RUN apt-get -y install postgresql-client

RUN go mod download
RUN go build -o ./bin/app ./cmd/apiserver/main.go 


ENTRYPOINT ["./bin/app", "-storage", "cache"]
