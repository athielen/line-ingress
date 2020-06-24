package backend

import (
	log "github.com/sirupsen/logrus"
	"line-ingress/models"
	"os"
)

type FileBackend struct {
	Lines chan models.Line
}

func (backend *FileBackend) Consume() {
	log.Debug("Register the worker for file backend channel")

	f, err := os.OpenFile("testfile", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}

	for i := range backend.Lines {
		f.WriteString(i.String() + "\n")
	}
	f.Close()

}

func (backend *FileBackend) Input() chan<- models.Line {
	return backend.Lines
}

func (backend *FileBackend) Output() <-chan models.Line {
	return backend.Lines
}
