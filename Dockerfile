FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .
RUN go mod tidy

RUN go build -o integration-tests .

ENTRYPOINT ["go", "test", "-v", "./tests/integration/..."]