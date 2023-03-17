package image

import (
	"github.com/google/uuid"

	"github.com/snet-commerce/product/internal/domain/model/values"
)

type ProductImage struct {
	id        uuid.UUID
	src       Source
	position  values.Position
	isDefault bool
}

func NewProductImage(src Source, pos values.Position, isDefault bool) (*ProductImage, error) {
	return &ProductImage{
		id:        uuid.New(),
		src:       src,
		position:  pos,
		isDefault: isDefault,
	}, nil
}

func (i *ProductImage) ID() uuid.UUID {
	return i.id
}

func (i *ProductImage) Source() Source {
	return i.src
}

func (i *ProductImage) Position() values.Position {
	return i.position
}

func (i *ProductImage) IsDefault() bool {
	return i.isDefault
}
