package dto

type NewProduct struct {
	Number      *string              `json:"number"`
	Name        string               `json:"name"`
	Description *string              `json:"description"`
	Metadata    []map[string]string  `json:"metadata"`
	Options     []string             `json:"options"`
	Variants    []*NewProductVariant `json:"variants"`
}

type NewProductVariant struct {
	Number   *string                    `json:"number"`
	Position int                        `json:"position"`
	Options  []*NewProductOptionMapping `json:"options"`
	Images   []*NewProductImage         `json:"images"`
	Metadata []map[string]string        `json:"metadata"`
}

type NewProductOptionMapping struct {
	Option string `json:"option"`
	Value  string `json:"value"`
}

type NewProductImage struct {
	Src       string `json:"src"`
	Position  int    `json:"position"`
	IsDefault bool   `json:"isDefault"`
}
