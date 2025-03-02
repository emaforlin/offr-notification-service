FROM golang:1.24-alpine AS builder

WORKDIR /src

RUN go install github.com/bufbuild/buf/cmd/buf@latest

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN buf generate

RUN CGO_ENABLED=0 GOOS=linux go build -o Whisper ./cmd/whisper

FROM alpine:3.21.3

WORKDIR /app

COPY --from=builder /src/Whisper /app

CMD [ "/app/Whisper" ]





