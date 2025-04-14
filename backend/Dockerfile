FROM golang:1.20-alpine

WORKDIR /app

# Copiar archivos de Go
COPY go.mod go.sum ./
RUN go mod download

# Copiar el código y los segmentos
COPY . .

# Asegurarse que el directorio hls_test existe
RUN mkdir -p /app/hls_test

# Construir la aplicación
RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
