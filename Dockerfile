FROM golang:1.24

WORKDIR /go/stc/app

COPY . .

EXPOSE 8000

RUN go build -o  ./cmd/main.go

CMD ["./main"]