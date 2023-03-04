package product

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func NewProduct(
	name string,
	description string,
	metadata map[string]string,
	now time.Time,
) (*Product, error) {
	if dto.Name == "" {
		return nil, fmt.Errorf("")
	}

	var description string
	if dto.Description != nil {
		description = *dto.Description
	}

	state := StateDraft
	if dto.IsActive {
		state = StateActive
	}

	return &Product{
		id:          uuid.New(),
		name:        dto.Name,
		description: description,
		state:       state,
		createdAt:   now,
		updatedAt:   now,
		publishedAt: time.Time{},
		metadata:    dto.Metadata,
		variants:    nil,
	}, nil
}
