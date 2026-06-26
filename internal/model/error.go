package model

import "fmt"

type ErrPersonNotFound struct {
	ID int64
}

func (e ErrPersonNotFound) Error() string {
	return fmt.Sprintf("Person with ID %d is not found", e.ID)
}

type ErrPersonNameFound struct {
	Name string
}

func (e ErrPersonNameFound) Error() string {
	return fmt.Sprintf("Person name %s not found", e.Name)
}

type ErrFieldsNotRegistered struct {
	Field string
}

func (e ErrFieldsNotRegistered) Error() string {
	return fmt.Sprintf("Developer Error: %v func is not registerd in personByIDRegistry", e.Field)
}
