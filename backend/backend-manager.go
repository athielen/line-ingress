package backend

import (
	log "github.com/sirupsen/logrus"
	"line-ingress/models"
	"time"
)

type BackendManager struct {
	File FileBackend
	Redis RedisBackend
}

func BuildBackendManager() BackendManager {
	log.Debug("Initializing backend manager...")

	log.Debug("Initializing file backend channel...")
	fb := FileBackend{make(chan models.Line)}

	log.Debug("Initializing redis backend channel...")
	rc := RClient()
	rb := RedisBackend{rc,make(chan models.Line)}

	return BackendManager{fb, rb}
}

func (bman BackendManager) StartBackends() {
	log.Debug("Starting backend channels to consume lines...")
	go bman.File.Consume()
	go bman.Redis.Consume()
}

func (bman BackendManager) SendLineToChannels(b []byte) {

	// create line model from supplied json
	line := models.Line{b, string(b), time.Now(), "web_analytics"}
	log.Debug("Sending line to channels: " + line.String())

	// pass line model to applicable channels
	bman.File.Lines <- line
	bman.Redis.Lines <- line
}
