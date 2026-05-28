package solutions

import (
	"errors"
	"fmt"
)


type ValidationError struct {
	Field string
	Msg   string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("validation failed on field %s: %s", e.Field, e.Msg)
}

var ErrConnection = errors.New("connection failed")

func CheckDatabaseError(err error) bool {
	return errors.Is(err, ErrConnection)
}

func SafeExecute(fn func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered panic: %v", r)
		}
	}()
	fn()
	return nil
}
