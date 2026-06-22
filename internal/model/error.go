package model

import "fmt"

type ErrPersonNotFound struct {
	ID int64
}

func (e ErrPersonNotFound) Error() string {
	return fmt.Sprintf("Person with ID %d not found", e.ID)
}
