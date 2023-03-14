package image

import (
	"errors"
	"net/url"
)

type Source string

func NewSource(rawUrl string) (Source, error) {
	if _, err := url.ParseRequestURI(rawUrl); err != nil {
		return "", errors.New("url must be absolute (example: https://example.com/resource)")
	}
	return Source(rawUrl), nil
}
