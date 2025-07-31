# Usa una imagen liviana de Go
FROM golang:1.24.5-alpine

# Variables de entorno necesarias
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Instala dependencias necesarias
RUN apk update && apk add --no-cache git

# Crea carpeta para app
WORKDIR /app

# Copia los archivos
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

# Compila la app
RUN go build -o main ./cmd/main.go

# Puerto de escucha
EXPOSE 8080

# Comando default
CMD ["./main"]
