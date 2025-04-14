package video

import (
	"context"
	"fmt"
	"hls_zappingtv/config"
	"strings"
	"sync"
	"time"
)

type StreamManager struct {
	mu             sync.Mutex
	baseSequence   int
	segmentList    []string
	currentSegment int
	streamCancel   context.CancelFunc
	config         *config.Config
}

func NewStreamManager(cfg *config.Config) *StreamManager {
	sm := &StreamManager{
		config: cfg,
	}
	sm.segmentList = sm.initializeSegments(cfg.MaxSegments)
	sm.currentSegment = cfg.MaxSegments
	return sm
}

func (sm *StreamManager) GetBasePath() string {
	return sm.config.BasePath
}

func (sm *StreamManager) initializeSegments(numberOfSegments int) []string {
	segmentList := make([]string, numberOfSegments)
	for i := 0; i < numberOfSegments; i++ {
		segmentList[i] = sm.formatSegment(i)
	}
	return segmentList
}

func (sm *StreamManager) StartStream(ctx context.Context) {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Stream detenido")
			return
		case <-ticker.C:
			sm.nextSegmentList()
		}
	}
}

func (sm *StreamManager) formatSegment(i int) string {
	return fmt.Sprintf("segment%d.ts", i)
}

func (sm *StreamManager) nextSegmentList() {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	if sm.currentSegment <= sm.config.TotalSegments-1 {
		sm.segmentList = append(sm.segmentList[1:], sm.formatSegment(sm.currentSegment))
		sm.baseSequence++
		sm.currentSegment++
	}
}

func (sm *StreamManager) ResetSegmentList() {
	// Preparar el nuevo contexto antes de adquirir el mutex
	ctx, cancel := context.WithCancel(context.Background())

	sm.mu.Lock()
	sm.segmentList = sm.initializeSegments(sm.config.MaxSegments)
	sm.baseSequence++
	sm.currentSegment = sm.config.MaxSegments
	sm.mu.Unlock()
	sm.SetStreamCancel(cancel)
	go sm.StartStream(ctx)
}

func (sm *StreamManager) PreviousSegment() {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	if sm.currentSegment <= sm.config.TotalSegments-1 && sm.currentSegment > sm.config.MaxSegments {
		sm.segmentList = append([]string{sm.formatSegment(sm.currentSegment - sm.config.MaxSegments - 1)}, sm.segmentList[:len(sm.segmentList)-1]...)
		sm.baseSequence++
		sm.currentSegment--
	}
}

func (sm *StreamManager) SetStreamCancel(cancel context.CancelFunc) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	if sm.streamCancel != nil {
		sm.streamCancel()
	}
	sm.streamCancel = cancel
}

func (sm *StreamManager) writeHeader(builder *strings.Builder) {
	fmt.Fprintln(builder, "#EXTM3U")
	fmt.Fprintln(builder, "#EXT-X-VERSION:3")
	fmt.Fprintf(builder, "#EXT-X-TARGETDURATION:%g\n", sm.config.SegmentDuration)
	fmt.Fprintf(builder, "#EXT-X-MEDIA-SEQUENCE:%d\n", sm.baseSequence)
}

func (sm *StreamManager) writeSegment(builder *strings.Builder, segment string, duration float32) {
	fmt.Fprintf(builder, "#EXTINF:%f,\n", duration)
	fmt.Fprintln(builder, segment)
}

func (sm *StreamManager) writeEndList(builder *strings.Builder) {
	fmt.Fprintln(builder, "#EXT-X-ENDLIST")
}

func (sm *StreamManager) WritePlaylist() string {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	var builder strings.Builder

	sm.writeHeader(&builder)

	for i, segment := range sm.segmentList {
		duration := sm.config.SegmentDuration

		if i == len(sm.segmentList)-1 && segment == sm.formatSegment(sm.config.TotalSegments-1) {
			duration = sm.config.LastSegmentDuration
		}

		sm.writeSegment(&builder, segment, duration)

		if i == len(sm.segmentList)-1 && sm.segmentList[len(sm.segmentList)-1] == sm.formatSegment(sm.config.TotalSegments-1) {
			sm.writeEndList(&builder)
		}
	}

	return builder.String()
}
