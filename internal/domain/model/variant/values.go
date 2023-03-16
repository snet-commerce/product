package variant

import (
	"errors"
	"net/url"

	"github.com/snet-commerce/product/internal/domain/model/values"
)

type OptionMapping struct {
	option values.ProductOption
	value  string
}

func NewOptionMapping(option values.ProductOption, value string) (OptionMapping, error) {
	if value == "" {
		return OptionMapping{}, errors.New("product option must have value to define mapping")
	}

	return OptionMapping{
		option: option,
		value:  value,
	}, nil
}

func (m OptionMapping) Option() values.ProductOption {
	return m.option
}

func (m OptionMapping) Value() string {
	return m.value
}

type ProductImage struct {
	src       string
	position  values.Position
	isDefault bool
}

func NewProductImage(uri string, pos values.Position, isDefault bool) (ProductImage, error) {
	if _, err := url.ParseRequestURI(uri); err != nil {
		return ProductImage{}, errors.New("url must be absolute (e.g. https://example.com/resource)")
	}

	return ProductImage{
		src:       uri,
		position:  pos,
		isDefault: isDefault,
	}, nil
}

func (i ProductImage) Source() string {
	return i.src
}

func (i ProductImage) Position() values.Position {
	return i.position
}

func (i ProductImage) IsDefault() bool {
	return i.isDefault
}
