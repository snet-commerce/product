package image

import (
	"github.com/google/uuid"
	"time"
)

type ProductImage struct {
	id        uuid.UUID
	src       string
	position  int
	createdAt time.Time
	updatedAt time.Time
}
