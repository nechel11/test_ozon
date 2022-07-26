FROM golang:1.18.3

RUN mkdir ozon
COPY ./ ./ozon
WORKDIR /ozon


RUN apt-get update
RUN apt-get -y install postgresql-client
RUN . ./db_init.sh $(cat configs/pg_config)
RUN go mod download
RUN go build -o app ./cmd/app/main.go

CMD ["./app"]