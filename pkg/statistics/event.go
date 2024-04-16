package statistics

import "encoding/json"

var _ Serializable = Event{}

type Event struct {
	Type       string         `json:"type,omitempty"`
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
