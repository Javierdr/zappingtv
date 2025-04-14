#!/bin/bash

# Crear directorio para los segmentos
mkdir -p backend/hls_test

# Descargar el archivo usando gdown (lo instalaremos con pip)
pip install gdown
gdown https://drive.google.com/uc?id=1exGq6BJ6r1lXezOanp88sWwxqcMbDntJ -O backend/hls_test/segments.zip

# Descomprimir el archivo
cd backend/hls_test
unzip -o segments.zip
rm segments.zip

echo "Segmentos de video preparados correctamente"
