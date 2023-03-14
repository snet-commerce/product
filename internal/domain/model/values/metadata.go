package values

import (
	"errors"
)

type Metadata map[string]string

func MetadataFromMap(m map[string]string) (Metadata, error) {
	meta := make(Metadata)
	for k, v := range m {
		if k == "" || v == "" {
			return nil, errors.New("metadata property and its value can't be empty")
		}
		meta[k] = v
	}
	return meta, nil
}
