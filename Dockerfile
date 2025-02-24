FROM golang:1.24-alpine AS builder

WORKDIR /src

RUN go install github.com/bufbuild/buf/cmd/buf@latest

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN buf generate

RUN CGO_ENABLED=0 GOOS=linux go build -o Gjallarhorn ./cmd/gjallarhorn

FROM golang:1.24-alpine

WORKDIR /app

COPY --from=builder /src/Gjallarhorn /app

CMD [ "/app/Gjallarhorn" ]





