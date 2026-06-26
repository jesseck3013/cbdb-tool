package view

import (
	"fmt"
	"io"
	"strconv"

	"github.com/jesseck3013/cbdb-tool/internal/repository"
)

func printPersonList(w io.Writer, ps []repository.GetPersonByNameRow) {
	fmt.Fprintln(w, headerStyle.Render("# Found Persons"))
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

	printTable(w, []string{"ID", "Name", "Lifespan", "Dynasty", "Dynasty Years"}, rows)
}
