package view

import (
	"encoding/json"
	"io"

	"github.com/jesseck3013/cbdb-tool/internal/model"
	"github.com/jesseck3013/cbdb-tool/internal/repository"
)

type TextRenderer struct {
}

func (t TextRenderer) PersonByID(w io.Writer, p *model.Person, fields []model.PersonField) error {
	PrintPerson(w, p, fields)
	return nil
}

func (t TextRenderer) PersonByName(w io.Writer, ps []repository.GetPersonByNameRow) error {
	printPersonList(w, ps)
	return nil
}

type JSONRenderer struct {
}

func (t JSONRenderer) PersonByID(w io.Writer, p *model.Person, fields []model.PersonField) error {
	en := json.NewEncoder(w)
	return en.Encode(p)
}

func (t JSONRenderer) PersonByName(w io.Writer, ps []repository.GetPersonByNameRow) error {
	en := json.NewEncoder(w)
	return en.Encode(ps)
}
