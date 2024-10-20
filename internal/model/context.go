package model

type ElementType string

const (
	ElementTypeText  ElementType = "text"
	ElementTypeImage ElementType = "image"
)

type Element interface {
	ElementType() ElementType
}

type Context struct {
	Type     string    `json:"type,omitempty"`
	Elements []Element `json:"elements,omitempty"`
}

func (c Context) BlockType() MessageBlockType {
	return MBTContext
}
