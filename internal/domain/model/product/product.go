package product

import (
	"errors"
	"time"

	"github.com/google/uuid"

	"github.com/snet-commerce/product/internal/domain/model/values"
	"github.com/snet-commerce/product/internal/domain/model/variant"
)

// TODO: think of pointers and error handling + think of reactive validation

type Product struct {
	id          uuid.UUID
	number      string
	name        string
	description string
	state       State
	publishedAt time.Time
	metadata    values.Metadata
	variants    []*variant.ProductVariant
}

func (p *Product) Activate() error {
	if p.state != StateDraft {
		return errors.New("only draft product can be activated")
	}
	p.state = StateActive
	return nil
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
