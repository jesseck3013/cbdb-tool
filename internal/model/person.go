package model

import (
	"context"

	"github.com/jesseck3013/cbdb-tool/internal/repository"
)

type Person struct {
	BasicInfo repository.GetPersonByIDRow

	AltNames []repository.GetAltnamesByPersonIDRow

	Entries []repository.GetEntryByPersonIDRow

	Institutions []repository.GetInstitutionByPersonIDRow

	Postings []repository.GetPostingByPersonIDRow

	Status []repository.GetStatusByPersonIDRow

	Texts []repository.GetTextByPersonIDRow

	KinShip []repository.GetPersonKinShipByPersonIDRow

	Associations []repository.GetAltnamesByPersonIDRow
}

func GetPerson(ctx context.Context, q *repository.Queries, id int64) (*Person, error) {
	person, err := q.GetPersonByID(ctx, id)
	if err != nil {
		return nil, err
	}

	altnames, err := q.GetAltnamesByPersonID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &Person{
		BasicInfo: person,
		AltNames:  altnames,
	}, nil
}
