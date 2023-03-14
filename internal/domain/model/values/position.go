package values

import "fmt"

type Position int

func NewPosition(p int) (Position, error) {
	if p <= 0 {
		return 0, fmt.Errorf("position must be greater or equal 1")
	}
	return Position(p), nil
}
