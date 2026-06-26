package view

import (
	"fmt"
	"io"
	"strconv"

	"charm.land/lipgloss/v2"
	"charm.land/lipgloss/v2/table"
	"github.com/jesseck3013/cbdb-tool/internal/model"
	"github.com/jesseck3013/cbdb-tool/internal/repository"
)

var (
	primaryColor = lipgloss.Color("#c7a587")

	labelStyle = lipgloss.NewStyle().
			Height(1).
			Width(20).
			Bold(true)

	// Style for valid data values
	valueStyle = lipgloss.NewStyle().
			Height(1).
			Foreground(lipgloss.Color("255")).
			Align(lipgloss.Left)

	// Style specifically for missing database points
	emptyValueStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("240")).
			Italic(true)

	// Header style for "Person Profile:"
	headerStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(primaryColor).
			Margin(1, 0, 1, 0)

	tableHeaderStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("255")).
				Align(lipgloss.Center)
)

func renderLine(key, value string) string {
	return lipgloss.JoinHorizontal(lipgloss.Left,
		labelStyle.Render(key),
		valueStyle.Render(value))
}

// Combine Chinese and English Data Entry
func joinTwoLang(lang1, lang2 string) string {
	return fmt.Sprintf("%s / %s", lang1, lang2)
}

// Combine two years to form a period
// Example: 1000 2000 -> 1000 - 2000
func joinTwoYear(lang1, lang2 string) string {
	return fmt.Sprintf("%s - %s", lang1, lang2)
}

func printBasicInfo(w io.Writer, info repository.GetPersonBasicInfoByIDRow) {
	fmt.Fprintln(w, headerStyle.Render("# Profile"))
	rows := []string{
		renderLine(" ID:", strconv.FormatInt(info.PersonID, 10)),
		renderLine(" Name:", joinTwoLang(safeValue(info.NameCh), safeValue(info.NameEn))),
		renderLine(" Dynasty:", joinTwoLang(safeValue(info.DynastyEn), safeValue(info.DynastyCh))),
		renderLine(" Choronym:", joinTwoLang(safeValue(info.ChoronymCh), safeValue(info.ChoronymEn))),
		renderLine(" Birth Year:", safeValue(info.BirthYear)),
		renderLine(" Death Year:", safeValue(info.DeathYear)),
	}

	profileBlock := lipgloss.JoinVertical(lipgloss.Left, rows...)

	fmt.Fprintln(w, profileBlock)
}

func PrintPerson(w io.Writer, person *model.Person, fields []model.PersonField) {
	for _, field := range fields {
		fn, ok := personByIDRegistry[field]
		if ok {
			fn(w, person)
		}
	}
}

func newTable(headers []string, rows [][]string) *table.Table {
	t := table.New().
		Border(lipgloss.ASCIIBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("255"))).
		Width(80).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch {
			case row == table.HeaderRow:
				return tableHeaderStyle
			default:
				return valueStyle
			}
		}).
		Headers(headers...).
		Rows(rows...)

	return t
}

func printTable(w io.Writer, headers []string, rows [][]string) {
	if len(rows) == 0 {
		fmt.Fprintln(w, valueStyle.Render("Data not found"))
		fmt.Fprintln(w, "")
		return
	}

	t := newTable(headers, rows)
	lipgloss.Fprintln(w, t)
}

func printAltNames(w io.Writer, names []repository.GetAltnamesByPersonIDRow) {
	fmt.Fprintln(w, headerStyle.Render("# Alternative Names"))
	rows := [][]string{}

	for _, name := range names {
		rows = append(rows,
			[]string{
				joinTwoLang(safeValue(name.NameTypeCh), safeValue(name.NameType)),
				joinTwoLang(name.AltnameCh, safeValue(name.AltnameEn)),
			})
	}

	printTable(w, []string{"Type", "Alternative Name"}, rows)
}

func printKinships(w io.Writer, kinships []repository.GetPersonKinShipByPersonIDRow) {
	fmt.Fprintln(w, headerStyle.Render("# Kinship"))
	rows := [][]string{}

	for _, kinship := range kinships {
		rows = append(rows,
			[]string{
				strconv.FormatInt(kinship.KinID, 10),
				joinTwoLang(safeValue(kinship.NameCh), safeValue(kinship.Name)),
				joinTwoLang(kinship.KinRelCh, kinship.KinRelEn),
			})
	}

	printTable(w, []string{"ID", "Name", "Type"}, rows)
}

func printAssociations(w io.Writer, associations []repository.GetAssociationByPersonIDRow) {
	fmt.Fprintln(w, headerStyle.Render("# Associations"))
	rows := [][]string{}

	for _, association := range associations {
		rows = append(rows,
			[]string{
				strconv.FormatInt(association.AssocID, 10),
				joinTwoLang(safeValue(association.NameCh), safeValue(association.NameEn)),
				joinTwoLang(safeValue(association.AssocTypeCh), safeValue(association.AssocTypeEn)),
			})
	}

	printTable(w, []string{"ID", "Name", "Type"}, rows)
}

func printStatus(w io.Writer, status []repository.GetStatusByPersonIDRow) {
	fmt.Fprintln(w, headerStyle.Render("# Status"))
	rows := [][]string{}

	for _, v := range status {
		rows = append(rows,
			[]string{
				joinTwoLang(v.StatusCh, v.StatusEn),
			})
	}

	printTable(w, []string{"Status"}, rows)
}

func printPlaces(w io.Writer, places []repository.GetPlaceByPersonIDRow) {
	fmt.Fprintln(w, headerStyle.Render("# Biographical Place Information"))
	rows := [][]string{}

	for _, v := range places {
		rows = append(rows,
			[]string{
				joinTwoLang(safeValue(v.NameCh), safeValue(v.NameEn)),
				joinTwoLang(safeValue(v.AddrCh), safeValue(v.AddrEn)),
				joinTwoYear(safeValue(v.FirstYear), safeValue(v.LastYear)),
			})
	}

	printTable(w, []string{"Place", "Type", "Period"}, rows)
}

func printEntries(w io.Writer, entries []repository.GetEntryByPersonIDRow) {
	fmt.Fprintln(w, headerStyle.Render("# Modes of Entry"))
	rows := [][]string{}

	for _, v := range entries {
		rows = append(rows,
			[]string{
				joinTwoLang(v.EntryCh, v.EntryTypeEn),
				strconv.FormatInt(int64(v.Year), 10),
				safeValue(v.Age),
				safeValue(v.ExamRank),
			})
	}

	printTable(w, []string{"Mode of Entry", "Year", "Age", "Exam Rank"}, rows)
}

func printPostings(w io.Writer, postings []repository.GetPostingByPersonIDRow) {
	fmt.Fprintln(w, headerStyle.Render("# Offices and Postings"))
	rows := [][]string{}

	for _, v := range postings {
		rows = append(rows,
			[]string{
				safeValue(v.OfficeCh),
				joinTwoLang(safeValue(v.ApptCn), safeValue(v.ApptEn)),
				joinTwoYear(safeValue(v.FirstYear), safeValue(v.LastYear)),
			})
	}

	printTable(w, []string{"Office", "Type", "Period"}, rows)
}

func printInstitutions(w io.Writer, insts []repository.GetInstitutionByPersonIDRow) {
	fmt.Fprintln(w, headerStyle.Render("# Social Institutions"))
	rows := [][]string{}

	for _, v := range insts {
		rows = append(rows,
			[]string{
				joinTwoLang(safeValue(v.InstNameHz), safeValue(v.InstNamePy)),
				joinTwoLang(safeValue(v.BiRoleCh), safeValue(v.BiRoleEn)),
				joinTwoYear(safeValue(v.BiBeginYear), safeValue(v.BiEndYear)),
			})
	}

	printTable(w, []string{"Institution", "Role", "Period"}, rows)
}

func printTexts(w io.Writer, insts []repository.GetTextByPersonIDRow) {
	fmt.Fprintln(w, headerStyle.Render("# Texts"))
	rows := [][]string{}

	for _, v := range insts {
		rows = append(rows,
			[]string{
				safeValue(v.TitleCh),
				safeValue(v.TextYear),
			})
	}

	printTable(w, []string{"Title", "Year"}, rows)
}

// TODO: build a generic version
func removeDupliate(s []repository.GetAssociationByPersonIDRow) []repository.GetAssociationByPersonIDRow {
	set := make(map[int64]*repository.GetAssociationByPersonIDRow)
	for _, v := range s {
		set[v.AssocID] = &v
	}

	res := make([]repository.GetAssociationByPersonIDRow, 0)

	for _, v := range set {
		res = append(res, *v)
	}

	return res
}

var personByIDRegistry = map[model.PersonField]func(io.Writer, *model.Person){
	model.BasicInfo: func(w io.Writer, p *model.Person) {
		printBasicInfo(w, p.BasicInfo)
	},
	model.AltName: func(w io.Writer, p *model.Person) {
		printAltNames(w, p.AltNames)
	},
	model.Entry: func(w io.Writer, p *model.Person) {
		printEntries(w, p.Entries)
	},
	model.Institution: func(w io.Writer, p *model.Person) {
		printInstitutions(w, p.Institutions)
	},
	model.Posting: func(w io.Writer, p *model.Person) {
		printPostings(w, p.Postings)
	},
	model.Status: func(w io.Writer, p *model.Person) {
		printStatus(w, p.Status)
	},
	model.Text: func(w io.Writer, p *model.Person) {
		printTexts(w, p.Texts)
	},
	model.KinShip: func(w io.Writer, p *model.Person) {
		printKinships(w, p.KinShips)
	},
	model.Association: func(w io.Writer, p *model.Person) {
		printAssociations(w, removeDupliate(p.Associations))
	},
	model.Place: func(w io.Writer, p *model.Person) {
		printPlaces(w, p.Places)
	},
}
