package view

import (
	"strconv"

	"charm.land/lipgloss/v2"
	"github.com/jesseck3013/cbdb-tool/internal/repository"
)

func printPersonList(ps []repository.GetPersonByNameRow) string {
	header := headerStyle.Render("# Found Persons")
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

	t := printTable([]string{"ID", "Name", "Lifespan", "Dynasty", "Dynasty Years"}, rows)
	return lipgloss.JoinVertical(lipgloss.Left, header, t)
}
