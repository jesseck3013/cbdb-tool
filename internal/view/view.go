package view

import (
	"encoding/json"
	"io"

	"github.com/jesseck3013/cbdb-tool/internal/model"
)

type TextRenderer struct {
}

func (t TextRenderer) PersonByID(w io.Writer, p *model.Person) error {
	PrintPerson(w, p)
	return nil
}

type JSONRenderer struct {
}

func (t JSONRenderer) PersonByID(w io.Writer, p *model.Person) error {
	en := json.NewEncoder(w)
	return en.Encode(p)
}
