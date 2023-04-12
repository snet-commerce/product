package variant

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/snet-commerce/list"

	"github.com/snet-commerce/product/internal/domain/model/image"
	"github.com/snet-commerce/product/internal/domain/model/values"
)

type ProductVariant struct {
	id       uuid.UUID
	number   values.ProductNumber
	position values.Position
	options  []OptionMapping
	images   *list.List[*image.ProductImage]
	metadata values.Metadata
}

func NewVariant(
	number values.ProductNumber,
	pos values.Position,
	options []OptionMapping,
	images []*image.ProductImage,
	metadata values.Metadata,
) (*ProductVariant, error) {
	if len(options) == 0 {
		return nil, errors.New("product variant must have value defined for each product option")
	}

	variant := &ProductVariant{
		id:       uuid.New(),
		number:   number,
		position: pos,
		options:  options,
		images:   list.New[*image.ProductImage](),
		metadata: metadata,
	}

	for _, img := range images {
		if err := variant.AddImage(img); err != nil {
			return nil, err
		}
	}

	if variant.defaultImage() == nil {
		return nil, errors.New("there must be default image specified for product variant")
	}

	return variant, nil
}

func (v *ProductVariant) AddImage(newImg *image.ProductImage) error {
	elem := v.images.Find(func(img *image.ProductImage) bool {
		return newImg.Position() == img.Position() || (newImg.IsDefault() && img.IsDefault())
	})
	if elem != nil {
		if elem.Value.Position() == newImg.Position() {
			return fmt.Errorf("image with position %d is already present", newImg.Position())
		}
		return errors.New("there is already default image present in product variant")
	}
	return nil
}

func (v *ProductVariant) defaultImage() *image.ProductImage {
	elem := v.images.Find(func(img *image.ProductImage) bool { return img.IsDefault() })
	if elem != nil {
		return elem.Value
	}
	return nil
}
