package statistics

import (
	"encoding/json"
	"time"
)

var _ Serializable = Event{}

type EventType string

const (
	LifecycleEvenType = "lifecycle"
	GraphicsEventType = "graphics"
	RebootEventType   = "reboot"
)

type Event struct {
	Type       EventType      `json:"type,omitempty"`
	Time       time.Time      `json:"time,omitempty"`
	Parameters map[string]any `json:"parameters,omitempty"`
}

func (e Event) Serialize() (data string, err error) {
	bytes, err := json.Marshal(e)
	if err != nil {
		return
	}

	data = string(bytes)
	return
}
