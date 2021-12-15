package centreonapi

import "github.com/pkg/errors"

var ErrNoFound error = errors.New("File not found")

func ErrIsNotFound(err error) bool {
	return err == ErrNoFound
}
