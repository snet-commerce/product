package variant

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/snet-commerce/list"

	"github.com/snet-commerce/product/internal/domain/model/values"
)

type ProductVariant struct {
	id       uuid.UUID
	number   values.ProductNumber
	position values.Position
	options  []OptionMapping
	metadata values.Metadata
	images   *list.List[ProductImage]
}

func NewVariant(
	number values.ProductNumber,
	position values.Position,
	options []OptionMapping,
) (*ProductVariant, error) {
	if len(options) == 0 {

	}
}

func (v *ProductVariant) AddImage(newImg ProductImage) error {
	elem := v.images.Find(func(img ProductImage) bool {
		return newImg.Position() == img.Position() || (newImg.IsDefault() && img.IsDefault())
	})
	if elem != nil {
		if elem.Value.Position() == newImg.Position() {
			return fmt.Errorf("image with position %d is already present", newImg.Position())
		}
		return errors.New("there is already default image present in variant")
	}
	return nil
}
