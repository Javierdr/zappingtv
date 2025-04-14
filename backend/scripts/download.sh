#!/bin/sh

echo "Descargando segmentos desde Google Drive..."
cd /app/hls_test

# Intentar descargar con diferentes métodos
if ! gdown --fuzzy --no-cookies "$GOOGLE_DRIVE_URL"; then
    echo "Intento alternativo de descarga..."
    gdown --id "1exGq6BJ6r1lXezOanp88sWwxqcMbDntJ"
fi

if [ -f segments.zip ]; then
    echo "Descomprimiendo archivos..."
    unzip segments.zip
    rm segments.zip
    echo "Segmentos preparados correctamente"
else
    echo "Error: No se pudo descargar el archivo"
    exit 1
fi
