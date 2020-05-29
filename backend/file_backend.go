package backend

import (
	"fmt"
	"athielen.com/line-ingress/pkg/models"
)

type FileBackend struct {
	Lines chan models.Line
}

func (backend *FileBackend) Consume() {
	fmt.Println("Register the worker")
	for i := range backend.Lines {
		fmt.Println(i)
	}
}

func (backend *FileBackend) Input() chan<- models.Line {
	return backend.Lines
}

func (backend *FileBackend) Output() <-chan models.Line {
	return backend.Lines
}