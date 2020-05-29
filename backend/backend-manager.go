package backend

import (
	"line-ingress/models"
	"time"
)

type BackendManager struct {
	File FileBackend
	Redis RedisBackend
}

func BuildBackendManager() BackendManager {
	fb := FileBackend{make(chan models.Line)}

	rc := RClient()
	rb := RedisBackend{rc,make(chan models.Line)}

	return BackendManager{fb, rb}
}

func (bman BackendManager) StartBackends() {
	go bman.File.Consume()
	go bman.Redis.Consume()
}

func (bman BackendManager) SendLineToChannels(b []byte) {

	// create line model from supplied json
	line := models.Line{b, time.Now(), "web_analytics"}

	// pass line model to applicable channels
	bman.File.Lines <- line
	bman.Redis.Lines <- line
}
