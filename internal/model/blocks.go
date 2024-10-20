package model

type MessageBlockType string

const (
	MBTSection MessageBlockType = "section"
	MBTContext MessageBlockType = "context"
)

// Block defines an interface all block types should implement
// to ensure consistency between blocks.
type Block interface {
	BlockType() MessageBlockType
}
