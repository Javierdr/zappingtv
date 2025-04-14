.PHONY: init build run

init:
	@echo "Inicializando proyecto..."
	@chmod +x scripts/init.sh
	@./scripts/init.sh

build: init
	@echo "Construyendo contenedores..."
	docker compose build

run: build
	@echo "Iniciando aplicación..."
	docker compose up

all: run
