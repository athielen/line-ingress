package backend

import (
	"fmt"
	"line-ingress/models"
	"log"
	"os"
)

type FileBackend struct {
	Lines chan models.Line
}

func (backend *FileBackend) Consume() {

	fmt.Println("Register the worker")

	f, err := os.OpenFile("testfile", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}

	for i := range backend.Lines {
		fmt.Println("file..." + i.String() + "\n")
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
