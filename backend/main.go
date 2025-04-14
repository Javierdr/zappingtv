package main

import (
	"context"
	"fmt"
	"hls_zappingtv/api"
	"hls_zappingtv/auth"
	"hls_zappingtv/config"
	"hls_zappingtv/video"
	"net/http"
)

func main() {
	// Crear el stream manager
	streamManager := video.NewStreamManager(&config.AppConfig)

	// Iniciar el stream con un nuevo contexto
	ctx, cancel := context.WithCancel(context.Background())
	streamManager.SetStreamCancel(cancel)
	go streamManager.StartStream(ctx)

	// Crear los handlers
	videoHandler := api.NewVideoHandler(streamManager)

	// Configurar rutas
	http.HandleFunc("/register", api.WithCORS(auth.RegisterHandler))
	http.HandleFunc("/login", api.WithCORS(auth.LoginHandler))

	http.HandleFunc("/playlist.m3u8", api.WithCORS(videoHandler.PlaylistHandler))
	http.HandleFunc("/previous", api.WithCORS(videoHandler.PreviousHandler))
	http.HandleFunc("/reset", api.WithCORS(videoHandler.ResetHandler))
	http.HandleFunc("/", api.WithCORS(videoHandler.SegmentHandler))

	fmt.Printf("Server HLS corriendo en el puerto %s\n", config.AppConfig.Port)
	if err := http.ListenAndServe(config.AppConfig.Port, nil); err != nil {
		fmt.Printf("Error al iniciar el servidor: %v\n", err)
	}
}
