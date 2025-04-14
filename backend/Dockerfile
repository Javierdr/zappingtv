FROM golang:1.20-alpine

# Instalar herramientas necesarias
RUN apk add --no-cache python3 py3-pip unzip

WORKDIR /app

# Instalar gdown con soporte de autenticación
RUN pip install --break-system-packages --no-cache-dir gdown

# Copiar archivos de Go
COPY go.mod go.sum ./
RUN go mod download

# Copiar el código
COPY . .

# Crear directorio para los segmentos y darle permisos al script
RUN mkdir -p /app/hls_test && \
    chmod +x /app/scripts/download.sh

# Descargar y preparar segmentos
RUN /app/scripts/download.sh

# Construir la aplicación
RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
