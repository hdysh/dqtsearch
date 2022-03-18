package formaterror

import (
	"errors"
)

func FormatError(err string) error {
	return errors.New("incorrect details")
}
