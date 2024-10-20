package model

type Message struct {
	IconEmoji string `json:"icon_emoji,omitempty"`
	// Channel   string  `json:"channel,omitempty"`
	Blocks []Block `json:"blocks,omitempty"`
}
