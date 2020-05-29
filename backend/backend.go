package backend

import "line-ingress/models"

// Backend consumes Line values and persists them to some kind of storage.
type Backend interface {

	// Consume starts the processing loop of the backend
	Consume()

	// Input returns the channel on which the backend receives Line values
	Input() chan<- models.Line

	// Output returns the channel on which the backend processes Line values
	Output() <-chan models.Line
}
