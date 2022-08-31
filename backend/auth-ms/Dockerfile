FROM golang:latest

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o auth-ms ./cmd/main/main.go

CMD ["./auth-ms"]

