package image

import (
	"time"

	"github.com/google/uuid"
)

type ProductImage struct {
	id        uuid.UUID
	src       string
	position  int
	createdAt time.Time
	updatedAt time.Time
}
