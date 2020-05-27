package backend

import (
	"athielen.com/line-ingress/pkg/models"
	"time"
)

type BackendManager struct {
	File FileBackend
}

func BuildBackendManager() BackendManager {
	fb := FileBackend{make(chan models.Line)}
	return BackendManager{fb}
}

func (bman BackendManager) StartBackends() {
	go bman.File.Consume()
}

func (bman BackendManager) SendLineToChannels(b []byte) {

	// create line model from supplied json
	line := models.Line{b, time.Now(), "web_analytics"}

	// pass line model to applicable channels
	bman.File.Lines <- line
}