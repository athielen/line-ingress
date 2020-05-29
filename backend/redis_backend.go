package backend

import (
	"fmt"
	"athielen.com/line-ingress/pkg/models"
)

type RedisBackend struct {
	Lines chan models.Line
}

func (rb *RedisBackend) Consume() {
	fmt.Println("Register the worker")
	for i := range rb.Lines {
		fmt.Println(i)
	}
}

func (rb *RedisBackend) Input() chan<- models.Line {
	return rb.Lines
}

func (rb *RedisBackend) Output() <-chan models.Line {
	return rb.Lines
}

