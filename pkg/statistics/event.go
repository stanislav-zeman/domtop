package statistics

var _ Serializable = Event{}

type Event struct {
	Type       string         `json:"type,omitempty"`
	Parameters map[string]any `json:"parameters,omitempty"`
}

func (e Event) Serialize() (string, error) {
	panic("unimplemented")
}
