package values

import (
	"errors"
)

type MetadataProperty struct {
	Key   string
	Value string
}

type Metadata map[string]string

func NewMetadata(props ...MetadataProperty) (Metadata, error) {
	meta := make(Metadata)
	for _, prop := range props {
		if prop.Key == "" || prop.Value == "" {
			return nil, errors.New("metadata property and its value can't be empty")
		}
		meta[prop.Key] = prop.Value
	}
	return meta, nil
}
