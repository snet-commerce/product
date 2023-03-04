package product

type NewProductDto struct {
	Name        string
	Description *string
	IsActive    bool
	Metadata    map[string]string
}

type VariantOptionDto struct {
	Name  string
	Value string
}

type NewProductImageDto struct {
}

type NewProductVariantDto struct {
	Code    string
	Options []VariantOptionDto
}
