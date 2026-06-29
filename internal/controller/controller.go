package controller

import (
	"context"

	"github.com/jesseck3013/cbdb-tool/internal/model"
)

type Controller interface {
	PersonController
}

type PersonController interface {
	PersonByID(ctx context.Context, input model.PersonByIDInput) ([]byte, error)
	PersonByName(ctx context.Context, name string) ([]byte, error)
}
