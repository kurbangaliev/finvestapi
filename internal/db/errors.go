package db

import (
	"errors"
	"fmt"
)

var ErrAutoMigrate = errors.New("error auto migrate")

type MigrateError struct {
	Model   string
	Message string
	Err     error // The underlying, wrapped error
}

func (e *MigrateError) Error() string {
	return fmt.Sprintf("model %s message: %s failed: %v", e.Model, e.Message, e.Err)
}

// Unwrap implements the unwrapping behavior for errors.Unwrap, errors.Is, and errors.As.
func (e *MigrateError) Unwrap() error {
	return e.Err
}
