package api

import (
	"fmt"
	"hls_zappingtv/video"
	"net/http"
)

type VideoHandler struct {
	manager *video.StreamManager
}

func NewVideoHandler(manager *video.StreamManager) *VideoHandler {
	return &VideoHandler{
		manager: manager,
	}
}

func (h *VideoHandler) SegmentHandler(w http.ResponseWriter, r *http.Request) {
	segmentName := r.URL.Path[1:]
	filePath := h.manager.GetBasePath() + segmentName
	http.ServeFile(w, r, filePath)
}

func (h *VideoHandler) PlaylistHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/vnd.apple.mpegurl")
	response := h.manager.WritePlaylist()
	fmt.Fprintln(w, response)
}

func (h *VideoHandler) ResetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	h.manager.ResetSegmentList()
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, `{"message": "Lista de segmentos reiniciada"}`)
}

func (h *VideoHandler) PreviousHandler(w http.ResponseWriter, r *http.Request) {
	h.manager.PreviousSegment()
	fmt.Fprintln(w, "Segmento anterior")
}
