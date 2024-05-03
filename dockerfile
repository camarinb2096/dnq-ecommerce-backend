FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o dnq-backend

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/dnq-backend .

CMD ["./dnq-backend"]

EXPOSE 8080


