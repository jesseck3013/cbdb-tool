package controller

import (
	"context"
	"io"

	"github.com/jesseck3013/cbdb-tool/internal/model"
	"github.com/jesseck3013/cbdb-tool/internal/repository"
)

type Store interface {
	GetPersonBasicInfoByID(context.Context, int64) (repository.GetPersonBasicInfoByIDRow, error)
	GetAltnamesByPersonID(ctx context.Context, cPersonid int64) ([]repository.GetAltnamesByPersonIDRow, error)
	GetAssociationByPersonID(ctx context.Context, cPersonid int64) ([]repository.GetAssociationByPersonIDRow, error)
	GetEntryByPersonID(ctx context.Context, cPersonid int64) ([]repository.GetEntryByPersonIDRow, error)
	GetInstitutionByPersonID(ctx context.Context, cPersonid int64) ([]repository.GetInstitutionByPersonIDRow, error)
	GetPersonByName(ctx context.Context, cNameChn *string) ([]repository.GetPersonByNameRow, error)
	GetPersonKinShipByPersonID(ctx context.Context, cPersonid int64) ([]repository.GetPersonKinShipByPersonIDRow, error)
	GetPlaceByPersonID(ctx context.Context, cPersonid int64) ([]repository.GetPlaceByPersonIDRow, error)
	GetPostingByPersonID(ctx context.Context, personID int64) ([]repository.GetPostingByPersonIDRow, error)
	GetStatusByPersonID(ctx context.Context, cPersonid int64) ([]repository.GetStatusByPersonIDRow, error)
}

type Renderer interface {
	PersonByID(io.Writer, *model.Person, []model.PersonField) error
	PersonByName(io.Writer, []repository.GetPersonByNameRow) error
}
