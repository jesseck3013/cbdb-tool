package controller

import (
	"context"

	"github.com/jesseck3013/cbdb-tool/internal/model"
)

type Controller interface {
	PersonController
}

type PersonController interface {
	personByID(ctx context.Context, input model.PersonByIDInput) ([]byte, error)

	personByName(ctx context.Context, name string) ([]byte, error)
}
