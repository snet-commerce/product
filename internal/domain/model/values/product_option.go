package values

import "errors"

type ProductOption string

func NewProductOption(opt string) (ProductOption, error) {
	if opt == "" {
		return "", errors.New("product option can't be initial")
	}
	return ProductOption(opt), nil
}
