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

	Associations []repository.GetAssociationByPersonIDRow
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

	entries, err := q.GetEntryByPersonID(ctx, id)
	if err != nil {
		return nil, err
	}

	institutions, err := q.GetInstitutionByPersonID(ctx, id)
	if err != nil {
		return nil, err
	}

	postings, err := q.GetPostingByPersonID(ctx, id)
	if err != nil {
		return nil, err
	}

	status, err := q.GetStatusByPersonID(ctx, id)
	if err != nil {
		return nil, err
	}

	texts, err := q.GetTextByPersonID(ctx, id)
	if err != nil {
		return nil, err
	}

	kinShip, err := q.GetPersonKinShipByPersonID(ctx, id)
	if err != nil {
		return nil, err
	}

	associations, err := q.GetAssociationByPersonID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &Person{
		BasicInfo:    person,
		AltNames:     altnames,
		Entries:      entries,
		Institutions: institutions,
		Postings:     postings,
		Status:       status,
		Texts:        texts,
		KinShip:      kinShip,
		Associations: associations,
	}, nil
}
