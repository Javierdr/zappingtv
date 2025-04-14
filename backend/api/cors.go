package api

import "net/http"

func WithCORS(handler http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
		writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Maneja las solicitudes preflight (OPTIONS)
		if request.Method == http.MethodOptions {
			writer.WriteHeader(http.StatusOK)
			return
		}

		handler(writer, request)
	}
}
