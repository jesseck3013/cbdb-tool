package model

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jesseck3013/cbdb-tool/internal/repository"
)

type PersonFieldFunc func(context.Context, int64, *Person) error

type PersonField string

const (
	BasicInfo   PersonField = "info"
	AltName     PersonField = "altname"
	Entry       PersonField = "entry"
	Institution PersonField = "institution"
	Posting     PersonField = "posting"
	Status      PersonField = "status"
	Text        PersonField = "text"
	KinShip     PersonField = "kinship"
	Association PersonField = "association"
	Place       PersonField = "place"
)

var ALLPersonFileds = []PersonField{
	BasicInfo,
	AltName,
	Entry,
	Institution,
	Posting,
	Status,
	Text,
	KinShip,
	Association,
	Place,
}

type PersonByIDInput struct {
	ID     int64
	Fileds []PersonField
}

type Person struct {
	BasicInfo    repository.GetPersonBasicInfoByIDRow
	AltNames     []repository.GetAltnamesByPersonIDRow
	Entries      []repository.GetEntryByPersonIDRow
	Institutions []repository.GetInstitutionByPersonIDRow
	Postings     []repository.GetPostingByPersonIDRow
	Status       []repository.GetStatusByPersonIDRow
	Texts        []repository.GetTextByPersonIDRow
	KinShips     []repository.GetPersonKinShipByPersonIDRow
	Associations []repository.GetAssociationByPersonIDRow
	Places       []repository.GetPlaceByPersonIDRow
}

type PersonIDStore interface {
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
	GetTextByPersonID(ctx context.Context, cPersonid int64) ([]repository.GetTextByPersonIDRow, error)
}

type Store interface {
	PersonIDStore
}

type Service struct {
	store              Store
	personByIDRegistry map[PersonField]PersonFieldFunc
}

func NewSerice(s Store) *Service {
	service := &Service{store: s}

	service.personByIDRegistry = map[PersonField]PersonFieldFunc{
		BasicInfo: func(ctx context.Context, id int64, p *Person) error {
			v, err := s.GetPersonBasicInfoByID(ctx, id)
			if err != nil {
				return err
			}
			p.BasicInfo = v
			return nil

		},
		AltName: func(ctx context.Context, id int64, p *Person) error {
			v, err := s.GetAltnamesByPersonID(ctx, id)
			if err != nil {
				return err
			}
			p.AltNames = v
			return nil
		},
		Entry: func(ctx context.Context, id int64, p *Person) error {
			v, err := s.GetEntryByPersonID(ctx, id)
			if err != nil {
				return err
			}
			p.Entries = v
			return nil
		},
		Institution: func(ctx context.Context, id int64, p *Person) error {
			v, err := s.GetInstitutionByPersonID(ctx, id)
			if err != nil {
				return err
			}
			p.Institutions = v
			return nil
		},
		Posting: func(ctx context.Context, id int64, p *Person) error {
			v, err := s.GetPostingByPersonID(ctx, id)
			if err != nil {
				return err
			}
			p.Postings = v
			return nil
		},
		Status: func(ctx context.Context, id int64, p *Person) error {
			v, err := s.GetStatusByPersonID(ctx, id)
			if err != nil {
				return err
			}
			p.Status = v
			return nil
		},
		Text: func(ctx context.Context, id int64, p *Person) error {
			v, err := s.GetTextByPersonID(ctx, id)
			if err != nil {
				return err
			}
			p.Texts = v
			return nil
		},
		KinShip: func(ctx context.Context, id int64, p *Person) error {
			v, err := s.GetPersonKinShipByPersonID(ctx, id)
			if err != nil {
				return err
			}
			p.KinShips = v
			return nil
		},
		Association: func(ctx context.Context, id int64, p *Person) error {
			v, err := s.GetAssociationByPersonID(ctx, id)
			if err != nil {
				return err
			}
			p.Associations = v
			return nil
		},
		Place: func(ctx context.Context, id int64, p *Person) error {
			v, err := s.GetPlaceByPersonID(ctx, id)
			if err != nil {
				return err
			}
			p.Places = v
			return nil
		},
	}

	return service
}

func (s *Service) FetchPersonByID(ctx context.Context, input PersonByIDInput) (*Person, error) {
	p := &Person{}

	// required field
	info, err := s.store.GetPersonBasicInfoByID(ctx, input.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrPersonNotFound{ID: input.ID}
		} else {
			return nil, err
		}
	}
	p.BasicInfo = info

	for _, field := range input.Fileds {
		fn, ok := s.personByIDRegistry[field]
		if ok {
			err := fn(ctx, input.ID, p)
			if err != nil {
				if err == sql.ErrNoRows {
					return nil, ErrPersonNotFound{ID: input.ID}
				} else {
					return nil, err
				}
			}
		} else {
			err := fmt.Sprintf("Developer Error: %v func is not registerd in personByIDRegistry", field)
			panic(err)
		}
	}

	return p, nil
}
