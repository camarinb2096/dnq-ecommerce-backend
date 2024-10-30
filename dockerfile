# Etapa de compilación
FROM golang:1.21.5-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o dnq-backend cmd/main.go

# Etapa final

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/dnq-backend /app/dnq-backend

EXPOSE 8080

CMD ["./dnq-backend"]

COPY --from=builder /app/dnq-backend .

CMD ["./dnq-backend"]

EXPOSE 8080