package variant

import (
	"time"

	"github.com/google/uuid"
)

func CreateProductVariant(
	code string,
	imageID uuid.UUID,
	images []uuid.UUID,
	position int,
	now time.Time,
) (*ProductVariant, error) {

}
