package image

import (
	"errors"
	"net/url"
)

type Source string

func NewSource(uri string) (Source, error) {
	if _, err := url.ParseRequestURI(uri); err != nil {
		return "", errors.New("url must be absolute (e.g. https://example.com/resource)")
	}
	return Source(uri), nil
}
