package view

import (
	"encoding/json"

	"github.com/jesseck3013/cbdb-tool/internal/model"
	"github.com/jesseck3013/cbdb-tool/internal/repository"
)

type TextRenderer struct {
}

func (t TextRenderer) PersonByID(p *model.Person, fields []model.PersonField) ([]byte, error) {
	s := PrintPerson(p, fields)
	return []byte(s), nil
}

func (t TextRenderer) PersonByName(ps []repository.GetPersonByNameRow) ([]byte, error) {
	s := printPersonList(ps)
	return []byte(s), nil
}

type JSONRenderer struct {
}

func (t JSONRenderer) PersonByID(p *model.Person, fields []model.PersonField) ([]byte, error) {
	return json.Marshal(p)
}

func (t JSONRenderer) PersonByName(ps []repository.GetPersonByNameRow) ([]byte, error) {
	return json.Marshal(ps)
}
