# Etapa 1: Compilación
FROM golang:1.21.5-alpine AS builder

# Directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar los archivos go.mod y go.sum
COPY go.mod go.sum ./

# Descargar las dependencias
RUN go mod download

# Copiar el código fuente
COPY . .

# Establecer el directorio de trabajo para la compilación
WORKDIR /app/cmd

# Compilar el binario
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/dnq-backend

# Etapa 2: Imagen ligera
FROM alpine:latest

# Crear un usuario no root para mayor seguridad
RUN adduser -D myuser

# Crear un directorio para la aplicación
WORKDIR /home/myuser/app

# Copiar el binario compilado desde la etapa de compilación
COPY --from=builder /app/dnq-backend .

# Cambiar los permisos del binario
RUN chown myuser:myuser dnq-backend

# Cambiar al usuario no root
USER myuser

# Puerto en el que la aplicación escucha (ajústalo según tu aplicación)
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./dnq-backend"]
