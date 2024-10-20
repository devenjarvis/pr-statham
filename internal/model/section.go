package model

type Section struct {
	Type string `json:"type,omitempty"`
	Text Text   `json:"text,omitempty"`
}

func (s Section) BlockType() MessageBlockType {
	return MBTSection
}
