package models

import (
	"encoding/json"
	"fmt"
	"time"
)

// Line corresponds to an incoming tracking request and contains the data
// to be persisted.
type Line struct {
	// Values are the query parameters encoded in JSON.
	Payload []byte

	// Time is the flavor-aware time the request was received.
	Time time.Time

	// Flavor specifies the instance to which the request corresponds to
	// (e.g. scrooge.co.uk).
	//Flavor *Flavor   <- replace with API KEY
	SourceType string
}

func (line Line) String() string {
	s, err := json.Marshal(line)
	if err == nil {
		return fmt.Sprintf("%s", string(s))
	} else {
		return "error"
	}
}