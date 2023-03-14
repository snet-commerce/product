package variant

import (
	"time"

	"github.com/google/uuid"
	"github.com/snet-commerce/list"

	"github.com/snet-commerce/product/internal/domain/model/image"
)

type ProductVariant struct {
	id        uuid.UUID
	number    string
	options   []Option
	imageID   uuid.UUID
	images    *list.List[*image.ProductImage]
	position  int
	createdAt time.Time
	updatedAt time.Time
}
