package model_test

import (
	"context"
	"errors"
	"math"
	"testing"

	"github.com/jesseck3013/cbdb-tool/internal/model"
	"github.com/jesseck3013/cbdb-tool/internal/repository"
)

const (
	MOCK_PERSON_ID = 100
)

var person100 = model.Person{
	BasicInfo: repository.GetPersonBasicInfoByIDRow{
		PersonID: MOCK_PERSON_ID,
	},

	AltNames:     []repository.GetAltnamesByPersonIDRow{},
	Entries:      []repository.GetEntryByPersonIDRow{},
	Institutions: []repository.GetInstitutionByPersonIDRow{},
	Postings:     []repository.GetPostingByPersonIDRow{},
	Status:       []repository.GetStatusByPersonIDRow{},
	Texts:        []repository.GetTextByPersonIDRow{},
	KinShips:     []repository.GetPersonKinShipByPersonIDRow{},
	Associations: []repository.GetAssociationByPersonIDRow{},
	Places:       []repository.GetPlaceByPersonIDRow{},
}

var DB = map[int64]model.Person{MOCK_PERSON_ID: person100}

type MockStore struct {
	db map[int64]model.Person
}

func (m *MockStore) GetPersonBasicInfoByID(ctx context.Context, id int64) (repository.GetPersonBasicInfoByIDRow, error) {
	person, ok := m.db[id]
	if !ok {
		return repository.GetPersonBasicInfoByIDRow{}, model.ErrPersonNotFound{ID: id}
	}

	return person.BasicInfo, nil
}

func (m *MockStore) GetPersonByName(ctx context.Context, name string) ([]repository.GetPersonByNameRow, error) {
	return []repository.GetPersonByNameRow{}, nil
}

func (m *MockStore) GetAltnamesByPersonID(ctx context.Context, cPersonid int64) ([]repository.GetAltnamesByPersonIDRow, error) {
	return []repository.GetAltnamesByPersonIDRow{}, nil
}

func (m *MockStore) GetAssociationByPersonID(ctx context.Context, cPersonid int64) ([]repository.GetAssociationByPersonIDRow, error) {
	return []repository.GetAssociationByPersonIDRow{}, nil
}

func (m *MockStore) GetEntryByPersonID(ctx context.Context, cPersonid int64) ([]repository.GetEntryByPersonIDRow, error) {
	return []repository.GetEntryByPersonIDRow{}, nil
}

func (m *MockStore) GetInstitutionByPersonID(ctx context.Context, cPersonid int64) ([]repository.GetInstitutionByPersonIDRow, error) {
	return []repository.GetInstitutionByPersonIDRow{}, nil
}

func (m *MockStore) GetPersonKinShipByPersonID(ctx context.Context, cPersonid int64) ([]repository.GetPersonKinShipByPersonIDRow, error) {
	return []repository.GetPersonKinShipByPersonIDRow{}, nil
}
func (m *MockStore) GetPlaceByPersonID(ctx context.Context, cPersonid int64) ([]repository.GetPlaceByPersonIDRow, error) {
	return []repository.GetPlaceByPersonIDRow{}, nil
}
func (m *MockStore) GetPostingByPersonID(ctx context.Context, personID int64) ([]repository.GetPostingByPersonIDRow, error) {
	return []repository.GetPostingByPersonIDRow{}, nil
}
func (m *MockStore) GetStatusByPersonID(ctx context.Context, cPersonid int64) ([]repository.GetStatusByPersonIDRow, error) {
	return []repository.GetStatusByPersonIDRow{}, nil
}
func (m *MockStore) GetTextByPersonID(ctx context.Context, cPersonid int64) ([]repository.GetTextByPersonIDRow, error) {
	return []repository.GetTextByPersonIDRow{}, nil
}

func assertBasicInfo(t *testing.T, want, got repository.GetPersonBasicInfoByIDRow) {
	t.Helper()
	if want != got {
		t.Errorf("Expected %v, got %v", want, got)
	}
}

func assertNoError(t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Fatalf("Not expecting error but got: %v", got)
	}
}

func assertNonNilField[T any](t *testing.T, fieldName model.PersonField, got []T) {
	t.Helper()
	if got == nil {
		t.Errorf("Field %s: expected non nil but got %v", fieldName, got)
	}
}

func TestFetchPersonByID(t *testing.T) {
	db := &MockStore{
		db: DB,
	}
	s := model.NewSerice(db)

	t.Run("Found person", func(t *testing.T) {
		input := model.PersonByIDInput{
			ID:     MOCK_PERSON_ID,
			Fileds: []model.PersonField{model.BasicInfo},
		}

		p, err := s.FetchPersonByID(context.Background(), input)
		assertNoError(t, err)
		assertBasicInfo(t, person100.BasicInfo, p.BasicInfo)
	})

	t.Run("Not found error", func(t *testing.T) {
		input := model.PersonByIDInput{
			ID:     math.MaxInt,
			Fileds: []model.PersonField{model.BasicInfo},
		}

		_, err := s.FetchPersonByID(context.Background(), input)
		want := model.ErrPersonNotFound{}
		if !errors.As(err, &want) {
			t.Fatalf("Expected Person not found but got %v", err)
		}
	})

}

func TestFetchPersonByIDWithFields(t *testing.T) {
	db := &MockStore{
		db: DB,
	}
	s := model.NewSerice(db)

	testCases := []struct {
		description        string
		fields             []model.PersonField
		expectBasicInfo    bool
		expectAltName      bool
		expectEntries      bool
		expectInstitutions bool
		expectPostings     bool
		expectStatus       bool
		expectTexts        bool
		expectKinShips     bool
		expectAssociations bool
		expectPlaces       bool
	}{
		{
			description:     "Expect Basic Info",
			fields:          []model.PersonField{model.BasicInfo},
			expectBasicInfo: true,
		},
		{
			description:        "Expect Altnames, Association",
			fields:             []model.PersonField{model.BasicInfo, model.AltName, model.Association},
			expectBasicInfo:    true,
			expectAltName:      true,
			expectAssociations: true,
		},
		{
			description: "Expect All",
			fields: []model.PersonField{
				model.BasicInfo,
				model.AltName,
				model.Entry,
				model.Institution,
				model.Posting,
				model.Status,
				model.Text,
				model.KinShip,
				model.Association,
				model.Place,
			},
			expectBasicInfo:    true,
			expectAltName:      true,
			expectEntries:      true,
			expectInstitutions: true,
			expectPostings:     true,
			expectStatus:       true,
			expectTexts:        true,
			expectKinShips:     true,
			expectAssociations: true,
			expectPlaces:       true,
		},
	}

	input := model.PersonByIDInput{ID: MOCK_PERSON_ID}
	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			input.Fileds = testCase.fields
			got, err := s.FetchPersonByID(context.Background(), input)
			assertNoError(t, err)

			if testCase.expectBasicInfo {
				assertBasicInfo(t, person100.BasicInfo, got.BasicInfo)
			}
			if testCase.expectAltName {
				assertNonNilField(t, model.AltName, got.AltNames)
			}
			if testCase.expectAssociations {
				assertNonNilField(t, model.Association, got.Associations)
			}
			if testCase.expectEntries {
				assertNonNilField(t, model.Entry, got.Entries)
			}
			if testCase.expectInstitutions {
				assertNonNilField(t, model.Institution, got.Institutions)
			}
			if testCase.expectKinShips {
				assertNonNilField(t, model.KinShip, got.KinShips)
			}
			if testCase.expectPlaces {
				assertNonNilField(t, model.Place, got.Places)
			}
			if testCase.expectPostings {
				assertNonNilField(t, model.Posting, got.Postings)
			}
			if testCase.expectStatus {
				assertNonNilField(t, model.Status, got.Status)
			}
			if testCase.expectTexts {
				assertNonNilField(t, model.Text, got.Texts)
			}
		})
	}
}
