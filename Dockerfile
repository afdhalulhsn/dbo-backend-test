FROM golang:1.18-alpine3.16

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN make build

EXPOSE 8830
#RUN make rest
ENTRYPOINT ["./main"]
