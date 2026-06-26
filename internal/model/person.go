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
	All         PersonField = "all"
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

var ALLPersonFileds = map[PersonField]struct{}{
	BasicInfo:   {},
	AltName:     {},
	Entry:       {},
	Institution: {},
	Posting:     {},
	Status:      {},
	Text:        {},
	KinShip:     {},
	Association: {},
	Place:       {},
}

func GetAllPersonFieldsSlice() []PersonField {
	fields := make([]PersonField, 0)

	for field, _ := range ALLPersonFileds {
		fields = append(fields, field)
	}

	return fields
}

func SelectFields(fields []string) []PersonField {
	set := make(map[PersonField]struct{}, 0)
	set[BasicInfo] = struct{}{}
	for _, field := range fields {
		set[PersonField(field)] = struct{}{}
	}

	res := make([]PersonField, 0)
	for field, _ := range set {
		res = append(res, field)
	}

	return res
}

type PersonByIDInput struct {
	ID     int64
	Fileds []PersonField
}

type Person struct {
	BasicInfo    repository.GetPersonBasicInfoByIDRow       `json:"basicInfo"`
	AltNames     []repository.GetAltnamesByPersonIDRow      `json:"altnames,omitempty"`
	Entries      []repository.GetEntryByPersonIDRow         `json:"entries,omitempty"`
	Institutions []repository.GetInstitutionByPersonIDRow   `json:"institutions,omitempty"`
	Postings     []repository.GetPostingByPersonIDRow       `json:"postings,omitempty"`
	Status       []repository.GetStatusByPersonIDRow        `json:"status,omitempty"`
	Texts        []repository.GetTextByPersonIDRow          `json:"texts,omitempty"`
	KinShips     []repository.GetPersonKinShipByPersonIDRow `json:"kinship,omitempty"`
	Associations []repository.GetAssociationByPersonIDRow   `json:"associations,omitempty"`
	Places       []repository.GetPlaceByPersonIDRow         `json:"places,omitempty"`
}

type PersonIDStore interface {
	GetPersonBasicInfoByID(context.Context, int64) (repository.GetPersonBasicInfoByIDRow, error)
	GetAltnamesByPersonID(ctx context.Context, cPersonid int64) ([]repository.GetAltnamesByPersonIDRow, error)
	GetAssociationByPersonID(ctx context.Context, cPersonid int64) ([]repository.GetAssociationByPersonIDRow, error)
	GetEntryByPersonID(ctx context.Context, cPersonid int64) ([]repository.GetEntryByPersonIDRow, error)
	GetInstitutionByPersonID(ctx context.Context, cPersonid int64) ([]repository.GetInstitutionByPersonIDRow, error)
	GetPersonKinShipByPersonID(ctx context.Context, cPersonid int64) ([]repository.GetPersonKinShipByPersonIDRow, error)
	GetPlaceByPersonID(ctx context.Context, cPersonid int64) ([]repository.GetPlaceByPersonIDRow, error)
	GetPostingByPersonID(ctx context.Context, personID int64) ([]repository.GetPostingByPersonIDRow, error)
	GetStatusByPersonID(ctx context.Context, cPersonid int64) ([]repository.GetStatusByPersonIDRow, error)
	GetTextByPersonID(ctx context.Context, cPersonid int64) ([]repository.GetTextByPersonIDRow, error)
}

type PersonNameStore interface {
	GetPersonByName(ctx context.Context, name string) ([]repository.GetPersonByNameRow, error)
}

type Store interface {
	PersonIDStore
	PersonNameStore
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

type PersonSearchResult struct {
	ID         int64  `json:"id"`
	Name       string `json:name,omitempty`
	NameChn    string `json:"nameChn,omitempty"`
	BirthYear  int16  `json:"birthYear,omitempty"`
	DeathYear  int16  `json:"deathYear,omitempty"`
	Dynasty    string `json:"dynasty,omitempty"`
	DynastyChn string `json:"dynastyChn,omitempty"`
	Start      int16  `json:"start"`
	End        int16  `json:"end"`
}

func (s *Service) FetchPersonByName(ctx context.Context, name string) ([]repository.GetPersonByNameRow, error) {
	ps, err := s.store.GetPersonByName(ctx, name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrPersonNameFound{name}
		} else {
			return nil, err
		}
	}

	return ps, nil
}
