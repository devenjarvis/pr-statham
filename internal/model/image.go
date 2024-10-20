package model

type Image struct {
	Type     string `json:"type,omitempty"`
	ImageUrl string `json:"image_url,omitempty"`
	AltText  string `json:"alt_text,omitempty"`
}

func (i Image) ElementType() ElementType {
	return ElementTypeImage
}
