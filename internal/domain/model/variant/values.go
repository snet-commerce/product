package variant

import (
	"errors"

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
