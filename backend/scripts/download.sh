#!/bin/sh

echo "Descargando segmentos desde Google Drive..."
cd /app/hls_test

# Intentar descargar con diferentes métodos
if ! gdown --fuzzy --no-cookies "https://drive.google.com/file/d/1exGq6BJ6r1lXezOanp88sWwxqcMbDntJ/view?usp=sharing"; then
    echo "Intento alternativo de descarga..."
    gdown --id "1exGq6BJ6r1lXezOanp88sWwxqcMbDntJ"
fi

if [ -f "hls test.zip" ]; then
    echo "Descomprimiendo archivos..."
    unzip "hls test.zip"
    rm hls test.zip
    echo "Segmentos preparados correctamente"
else
    echo "Error: No se pudo descargar el archivo"
    exit 1
fi
