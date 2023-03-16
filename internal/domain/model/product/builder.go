package product

import "github.com/snet-commerce/product/internal/domain/model/values"

type Builder struct {
	number      values.ProductNumber
	name        string
	description string
	options     []values.ProductOption
	props       []values.MetadataProperty
}

func NewBuilder() *Builder {
	return &Builder{
		options: make([]values.ProductOption, 0),
		props:   make([]values.MetadataProperty, 0),
	}
}

func (b *Builder) Number(number values.ProductNumber) *Builder {
	b.number = number
	return b
}

func (b *Builder) Name(name string) *Builder {
	b.name = name
	return b
}

func (b *Builder) Description(desc string) *Builder {
	b.description = desc
	return b
}

func (b *Builder) Options(opts ...values.ProductOption) *Builder {
	b.options = append(b.options, opts...)
	return b
}

func (b *Builder) Metadata(props ...values.MetadataProperty) *Builder {
	b.props = append(b.props, props...)
	return b
}

func (b *Builder) Build() (*Product, error) {

}
