package semantic

import "errors"

var (
	ErrSymbolExists        error = errors.New("symbol already exists")
	ErrSymbolNotFound      error = errors.New("symbol not found")
	ErrSymbolNotAccessible error = errors.New("symbol not accessible from this context")
)
