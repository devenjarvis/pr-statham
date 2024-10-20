package model

type Text struct {
	Type string `json:"type,omitempty"`
	Text string `json:"text,omitempty"`
}

func (t Text) ElementType() ElementType {
	return ElementTypeText
}
