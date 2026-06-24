package view

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/jesseck3013/cbdb-tool/internal/model"
	"github.com/jesseck3013/cbdb-tool/internal/repository"
)

type TextRenderer struct {
}

func (t TextRenderer) PersonByID(w io.Writer, p *model.Person, fields []model.PersonField) error {
	PrintPerson(w, p, fields)
	return nil
}

func printPersonList(ps []repository.GetPersonByNameRow) {
	fmt.Println(headerStyle.Render("# Found Persons"))
	rows := [][]string{}

	for _, v := range ps {
		rows = append(rows,
			[]string{
				strconv.FormatInt(v.PersonID, 10),
				joinTwoLang(safeValue(v.NameCh), safeValue(v.NameEn)),
				joinTwoYear(safeValue(v.BirthYear), safeValue(v.DeathYear)),
				joinTwoLang(safeValue(v.DynastyCh), safeValue(v.DynastyEn)),
				joinTwoYear(strconv.FormatInt(int64(v.DynastyStart), 10), strconv.FormatInt(int64(v.DynastyEnd), 10)),
			})
	}

	printTable([]string{"ID", "Name", "Lifespan", "Dynasty", "Dynasty Years"}, rows)
}

func PrintPersonByName(w io.Writer, ps []repository.GetPersonByNameRow) {

}

func (t TextRenderer) PersonByName(w io.Writer, ps []repository.GetPersonByNameRow) error {
	printPersonList(ps)
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
