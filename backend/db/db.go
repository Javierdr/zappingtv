package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

var DB *pgxpool.Pool

func init() {
	connStr := os.Getenv("DATABASE_URL")

	maxRetries := 5
	var err error

	for i := 0; i < maxRetries; i++ {
		DB, err = pgxpool.Connect(context.Background(), connStr)
		if err == nil {
			fmt.Println("Base de datos conectada correctamente")
			return
		}

		fmt.Printf("Intento %d: Error al conectar con la base de datos: %v\n", i+1, err)
		if i < maxRetries-1 {
			fmt.Printf("Reintentando en 5 segundos...\n")
			time.Sleep(5 * time.Second)
		}
	}

	log.Fatalf("Error al conectar con la base de datos después de %d intentos: %v", maxRetries, err)
}
