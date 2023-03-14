package image

import (
	"github.com/google/uuid"
	"github.com/snet-commerce/product/internal/domain/model/values"
)

type ProductImage struct {
	id       uuid.UUID
	src      Source
	position values.Position
}

func New(src Source, pos values.Position) *ProductImage {
	return &ProductImage{
		id:       uuid.New(),
		src:      src,
		position: pos,
	}
}
