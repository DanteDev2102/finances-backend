FROM golang:1.22

WORKDIR /usr/app

COPY . .

RUN go mod tidy
RUN go mod download

CMD ["go", "run", "cmd/main.go"]