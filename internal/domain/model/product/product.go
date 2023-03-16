package product

import (
	"errors"
	"time"

	"github.com/google/uuid"

	"github.com/snet-commerce/product/internal/domain/model/values"
	"github.com/snet-commerce/product/internal/domain/model/variant"
)

type Product struct {
	id          uuid.UUID
	number      values.ProductNumber
	name        string
	description string
	state       State
	publishedAt time.Time
	metadata    values.Metadata
	options     []values.ProductOption
	variants    []*variant.ProductVariant
}

func NewProduct(
	number values.ProductNumber,
	name string,
	description string,
) (*Product, error) {
	return &Product{
		id:          uuid.New(),
		number:      number,
		name:        name,
		description: description,
		state:       0,
		publishedAt: time.Time{},
		metadata:    nil,
		variants:    nil,
	}, nil
}

func (p *Product) Archive() error {
	if p.state != StateActive {
		return errors.New("only active product can be archived")
	}
	p.state = StateArchived
	return nil
}

func (p *Product) Restore() error {
	if p.state != StateArchived {
		return errors.New("only archived product can be reactivated")
	}
	p.state = StateActive
	return nil
}

func (p *Product) Rename(name string) error {
	if p.name == "" {
		return errors.New("name can't be initial")
	}
	p.name = name
	return nil
}

func (p *Product) ChangeDescription(description string) {
	p.description = description
}
