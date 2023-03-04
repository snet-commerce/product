package variant

import (
	"time"

	"github.com/google/uuid"

	"github.com/snet-commerce/product/internal/domain/model/image"
)

type ProductVariant struct {
	id        uuid.UUID
	code      string
	options   []Option
	imageID   uuid.UUID
	images    []*image.ProductImage
	position  int
	createdAt time.Time
	updatedAt time.Time
}
