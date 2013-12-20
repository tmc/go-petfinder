package petfinder

import (
	"errors"
)

var (
        // ErrMissingAPIKey is returned if there is an attempt to create an api instance without a key
	ErrMissingAPIKey = errors.New("Missing API Key")
)
