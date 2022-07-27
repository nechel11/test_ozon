FROM golang:1.18.3

RUN mkdir -p ozon
COPY ./ /ozon
WORKDIR /ozon

RUN go mod download
RUN go build -o ./bin/app ./cmd/apiserver/main.go 

ENTRYPOINT ["./bin/app"]
