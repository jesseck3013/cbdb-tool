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

func printBasicInfo(info repository.GetPersonBasicInfoByIDRow) {
	fmt.Println(headerStyle.Render("# Profile"))
	rows := []string{
		renderLine(" ID:", strconv.FormatInt(info.PersonID, 10)),
		renderLine(" Name:", joinTwoLang(safeValue(info.NameCh), safeValue(info.NameEn))),
		renderLine(" Dynasty:", joinTwoLang(safeValue(info.DynastyEn), safeValue(info.DynastyCh))),
		renderLine(" Choronym:", joinTwoLang(safeValue(info.ChoronymCh), safeValue(info.ChoronymEn))),
		renderLine(" Birth Year:", safeValue(info.BirthYear)),
		renderLine(" Death Year:", safeValue(info.DeathYear)),
	}

	profileBlock := lipgloss.JoinVertical(lipgloss.Left, rows...)

	fmt.Println(profileBlock)
}

func PrintPerson(w io.Writer, person *model.Person, fields []model.PersonField) {
	for _, field := range fields {
		fn, ok := personByIDRegistry[field]
		if ok {
			fn(person)
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

func printTable(headers []string, rows [][]string) {
	if len(rows) == 0 {
		// t := newTable(headers, [][]string{
		// 	{"No data found in the database"},
		// })
		// lipgloss.Println(t)
		fmt.Println(valueStyle.Render("Data not found"))
		fmt.Println("")
		return
	}

	t := newTable(headers, rows)
	lipgloss.Println(t)
}

func printAltNames(names []repository.GetAltnamesByPersonIDRow) {
	fmt.Println(headerStyle.Render("# Alternative Names"))
	rows := [][]string{}

	for _, name := range names {
		rows = append(rows,
			[]string{
				joinTwoLang(safeValue(name.NameTypeCh), safeValue(name.NameType)),
				joinTwoLang(name.AltnameCh, safeValue(name.AltnameEn)),
			})
	}

	printTable([]string{"Type", "Alternative Name"}, rows)
}

func printKinships(kinships []repository.GetPersonKinShipByPersonIDRow) {
	fmt.Println(headerStyle.Render("# Kinship"))
	rows := [][]string{}

	for _, kinship := range kinships {
		rows = append(rows,
			[]string{
				strconv.FormatInt(kinship.KinID, 10),
				joinTwoLang(safeValue(kinship.NameCh), safeValue(kinship.Name)),
				joinTwoLang(kinship.KinRelCh, kinship.KinRelEn),
			})
	}

	printTable([]string{"ID", "Name", "Type"}, rows)
}

func printAssociations(associations []repository.GetAssociationByPersonIDRow) {
	fmt.Println(headerStyle.Render("# Associations"))
	rows := [][]string{}

	for _, association := range associations {
		rows = append(rows,
			[]string{
				strconv.FormatInt(association.AssocID, 10),
				joinTwoLang(safeValue(association.NameCh), safeValue(association.NameEn)),
				joinTwoLang(safeValue(association.AssocTypeCh), safeValue(association.AssocTypeEn)),
			})
	}

	printTable([]string{"ID", "Name", "Type"}, rows)
}

func printStatus(status []repository.GetStatusByPersonIDRow) {
	fmt.Println(headerStyle.Render("# Status"))
	rows := [][]string{}

	for _, v := range status {
		rows = append(rows,
			[]string{
				joinTwoLang(v.StatusCh, v.StatusEn),
			})
	}

	printTable([]string{"Status"}, rows)
}

func printPlaces(places []repository.GetPlaceByPersonIDRow) {
	fmt.Println(headerStyle.Render("# Biographical Place Information"))
	rows := [][]string{}

	for _, v := range places {
		rows = append(rows,
			[]string{
				joinTwoLang(safeValue(v.NameCh), safeValue(v.NameEn)),
				joinTwoLang(safeValue(v.AddrCh), safeValue(v.AddrEn)),
				joinTwoYear(safeValue(v.FirstYear), safeValue(v.LastYear)),
			})
	}

	printTable([]string{"Place", "Type", "Period"}, rows)
}

func printEntries(entries []repository.GetEntryByPersonIDRow) {
	fmt.Println(headerStyle.Render("# Modes of Entry"))
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

	printTable([]string{"Mode of Entry", "Year", "Age", "Exam Rank"}, rows)
}

func printPostings(postings []repository.GetPostingByPersonIDRow) {
	fmt.Println(headerStyle.Render("# Offices and Postings"))
	rows := [][]string{}

	for _, v := range postings {
		rows = append(rows,
			[]string{
				safeValue(v.OfficeCh),
				joinTwoLang(safeValue(v.ApptCn), safeValue(v.ApptEn)),
				joinTwoYear(safeValue(v.FirstYear), safeValue(v.LastYear)),
			})
	}

	printTable([]string{"Office", "Type", "Period"}, rows)
}

func printInstitutions(insts []repository.GetInstitutionByPersonIDRow) {
	fmt.Println(headerStyle.Render("# Social Institutions"))
	rows := [][]string{}

	for _, v := range insts {
		rows = append(rows,
			[]string{
				joinTwoLang(safeValue(v.InstNameHz), safeValue(v.InstNamePy)),
				joinTwoLang(safeValue(v.BiRoleCh), safeValue(v.BiRoleEn)),
				joinTwoYear(safeValue(v.BiBeginYear), safeValue(v.BiEndYear)),
			})
	}

	printTable([]string{"Institution", "Role", "Period"}, rows)
}

func printTexts(insts []repository.GetTextByPersonIDRow) {
	fmt.Println(headerStyle.Render("# Texts"))
	rows := [][]string{}

	for _, v := range insts {
		rows = append(rows,
			[]string{
				safeValue(v.TitleCh),
				safeValue(v.TextYear),
			})
	}

	printTable([]string{"Title", "Year"}, rows)
}

func removeDupliate(s []repository.GetAssociationByPersonIDRow) []repository.GetAssociationByPersonIDRow {
	set := make(map[int64]bool)
	res := make([]repository.GetAssociationByPersonIDRow, 0)
	for _, v := range s {
		if !set[v.AssocID] {
			set[v.AssocID] = true
			res = append(res, v)
		}
	}

	return res
}

var personByIDRegistry = map[model.PersonField]func(*model.Person){
	model.BasicInfo: func(p *model.Person) {
		printBasicInfo(p.BasicInfo)
	},
	model.AltName: func(p *model.Person) {
		printAltNames(p.AltNames)
	},
	model.Entry: func(p *model.Person) {
		printEntries(p.Entries)
	},
	model.Institution: func(p *model.Person) {
		printInstitutions(p.Institutions)
	},
	model.Posting: func(p *model.Person) {
		printPostings(p.Postings)
	},
	model.Status: func(p *model.Person) {
		printStatus(p.Status)
	},
	model.Text: func(p *model.Person) {
		printTexts(p.Texts)
	},
	model.KinShip: func(p *model.Person) {
		printKinships(p.KinShips)
	},
	model.Association: func(p *model.Person) {
		printAssociations(removeDupliate(p.Associations))
	},
	model.Place: func(p *model.Person) {
		printPlaces(p.Places)
	},
}
