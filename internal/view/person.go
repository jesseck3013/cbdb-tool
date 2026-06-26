package view

import (
	"fmt"
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

func printBasicInfo(info repository.GetPersonBasicInfoByIDRow) string {
	header := headerStyle.Render("# Profile")
	rows := []string{
		renderLine(" ID:", strconv.FormatInt(info.PersonID, 10)),
		renderLine(" Name:", joinTwoLang(safeValue(info.NameCh), safeValue(info.NameEn))),
		renderLine(" Dynasty:", joinTwoLang(safeValue(info.DynastyEn), safeValue(info.DynastyCh))),
		renderLine(" Choronym:", joinTwoLang(safeValue(info.ChoronymCh), safeValue(info.ChoronymEn))),
		renderLine(" Birth Year:", safeValue(info.BirthYear)),
		renderLine(" Death Year:", safeValue(info.DeathYear)),
	}

	profileBlock := lipgloss.JoinVertical(lipgloss.Left, rows...)

	return lipgloss.JoinVertical(lipgloss.Left, header, profileBlock)
}

func PrintPerson(person *model.Person, fields []model.PersonField) string {
	var res = ""
	for _, field := range fields {
		fn, ok := personByIDRegistry[field]
		if ok {
			res = lipgloss.JoinVertical(lipgloss.Left, res, fn(person))
		}
	}
	return res
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

func printTable(headers []string, rows [][]string) string {
	if len(rows) == 0 {
		return valueStyle.Render("No data is found\n")
	}

	t := newTable(headers, rows)
	return t.String() + "\n"
}

func printAltNames(names []repository.GetAltnamesByPersonIDRow) string {
	header := headerStyle.Render("# Alternative Names")
	rows := [][]string{}

	for _, name := range names {
		rows = append(rows,
			[]string{
				joinTwoLang(safeValue(name.NameTypeCh), safeValue(name.NameType)),
				joinTwoLang(name.AltnameCh, safeValue(name.AltnameEn)),
			})
	}

	t := printTable([]string{"Type", "Alternative Name"}, rows)
	return lipgloss.JoinVertical(lipgloss.Left, header, t)
}

func printKinships(kinships []repository.GetPersonKinShipByPersonIDRow) string {
	header := headerStyle.Render("# Kinship")
	rows := [][]string{}

	for _, kinship := range kinships {
		rows = append(rows,
			[]string{
				strconv.FormatInt(kinship.KinID, 10),
				joinTwoLang(safeValue(kinship.NameCh), safeValue(kinship.Name)),
				joinTwoLang(kinship.KinRelCh, kinship.KinRelEn),
			})
	}

	t := printTable([]string{"ID", "Name", "Type"}, rows)
	return lipgloss.JoinVertical(lipgloss.Left, header, t)
}

func printAssociations(associations []repository.GetAssociationByPersonIDRow) string {
	header := headerStyle.Render("# Associations")
	rows := [][]string{}

	for _, association := range associations {
		rows = append(rows,
			[]string{
				strconv.FormatInt(association.AssocID, 10),
				joinTwoLang(safeValue(association.NameCh), safeValue(association.NameEn)),
				joinTwoLang(safeValue(association.AssocTypeCh), safeValue(association.AssocTypeEn)),
			})
	}

	t := printTable([]string{"ID", "Name", "Type"}, rows)
	return lipgloss.JoinVertical(lipgloss.Left, header, t)
}

func printStatus(status []repository.GetStatusByPersonIDRow) string {
	header := headerStyle.Render("# Status")
	rows := [][]string{}

	for _, v := range status {
		rows = append(rows,
			[]string{
				joinTwoLang(v.StatusCh, v.StatusEn),
			})
	}

	t := printTable([]string{"Status"}, rows)
	return lipgloss.JoinVertical(lipgloss.Left, header, t)
}

func printPlaces(places []repository.GetPlaceByPersonIDRow) string {
	header := headerStyle.Render("# Biographical Place Information")
	rows := [][]string{}

	for _, v := range places {
		rows = append(rows,
			[]string{
				joinTwoLang(safeValue(v.NameCh), safeValue(v.NameEn)),
				joinTwoLang(safeValue(v.AddrCh), safeValue(v.AddrEn)),
				joinTwoYear(safeValue(v.FirstYear), safeValue(v.LastYear)),
			})
	}

	t := printTable([]string{"Place", "Type", "Period"}, rows)
	return lipgloss.JoinVertical(lipgloss.Left, header, t)
}

func printEntries(entries []repository.GetEntryByPersonIDRow) string {
	header := headerStyle.Render("# Modes of Entry")
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

	t := printTable([]string{"Mode of Entry", "Year", "Age", "Exam Rank"}, rows)
	return lipgloss.JoinVertical(lipgloss.Left, header, t)
}

func printPostings(postings []repository.GetPostingByPersonIDRow) string {
	header := headerStyle.Render("# Offices and Postings")
	rows := [][]string{}

	for _, v := range postings {
		rows = append(rows,
			[]string{
				safeValue(v.OfficeCh),
				joinTwoLang(safeValue(v.ApptCn), safeValue(v.ApptEn)),
				joinTwoYear(safeValue(v.FirstYear), safeValue(v.LastYear)),
			})
	}

	t := printTable([]string{"Office", "Type", "Period"}, rows)
	return lipgloss.JoinVertical(lipgloss.Left, header, t)
}

func printInstitutions(insts []repository.GetInstitutionByPersonIDRow) string {
	header := headerStyle.Render("# Social Institutions")
	rows := [][]string{}

	for _, v := range insts {
		rows = append(rows,
			[]string{
				joinTwoLang(safeValue(v.InstNameHz), safeValue(v.InstNamePy)),
				joinTwoLang(safeValue(v.BiRoleCh), safeValue(v.BiRoleEn)),
				joinTwoYear(safeValue(v.BiBeginYear), safeValue(v.BiEndYear)),
			})
	}

	t := printTable([]string{"Institution", "Role", "Period"}, rows)
	return lipgloss.JoinVertical(lipgloss.Left, header, t)
}

func printTexts(insts []repository.GetTextByPersonIDRow) string {
	header := headerStyle.Render("# Texts")
	rows := [][]string{}

	for _, v := range insts {
		rows = append(rows,
			[]string{
				safeValue(v.TitleCh),
				safeValue(v.TextYear),
			})
	}

	t := printTable([]string{"Title", "Year"}, rows)
	return lipgloss.JoinVertical(lipgloss.Left, header, t)
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

var personByIDRegistry = map[model.PersonField]func(*model.Person) string{
	model.BasicInfo: func(p *model.Person) string {
		return printBasicInfo(p.BasicInfo)
	},
	model.AltName: func(p *model.Person) string {
		return printAltNames(p.AltNames)
	},
	model.Entry: func(p *model.Person) string {
		return printEntries(p.Entries)
	},
	model.Institution: func(p *model.Person) string {
		return printInstitutions(p.Institutions)
	},
	model.Posting: func(p *model.Person) string {
		return printPostings(p.Postings)
	},
	model.Status: func(p *model.Person) string {
		return printStatus(p.Status)
	},
	model.Text: func(p *model.Person) string {
		return printTexts(p.Texts)
	},
	model.KinShip: func(p *model.Person) string {
		return printKinships(p.KinShips)
	},
	model.Association: func(p *model.Person) string {
		return printAssociations(removeDupliate(p.Associations))
	},
	model.Place: func(p *model.Person) string {
		return printPlaces(p.Places)
	},
}
