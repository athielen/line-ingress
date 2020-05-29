package models

import (
	"encoding/json"
	"fmt"
	"time"
)

// Line corresponds to an incoming tracking request and contains the data
// to be persisted.
type Line struct {
	Payload []byte
	Time time.Time
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
