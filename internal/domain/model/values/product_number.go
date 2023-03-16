package values

import (
	"fmt"
)

const productNumberMaxLength = 148

type ProductNumber string

func NewProductNumber(number string) (ProductNumber, error) {
	if len(number) > productNumberMaxLength {
		return "", fmt.Errorf("product number can't be more than %d", productNumberMaxLength)
	}
	return ProductNumber(number), nil
}
