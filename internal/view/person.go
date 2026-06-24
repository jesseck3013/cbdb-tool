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
		renderLine(" ID:", strconv.FormatInt(info.CPersonid, 10)),
		renderLine(" Name:", joinTwoLang(safeValue(info.CNameChn), safeValue(info.CName))),
		renderLine(" Dynasty:", joinTwoLang(safeValue(info.CDynastyChn), safeValue(info.CDynasty))),
		renderLine(" Choronym:", joinTwoLang(safeValue(info.CChoronymChn), safeValue(info.CChoronymDesc))),
		renderLine(" Birth Year:", safeValue(info.CBirthyear)),
		renderLine(" Death Year:", safeValue(info.CDeathyear)),
	}

	profileBlock := lipgloss.JoinVertical(lipgloss.Left, rows...)

	fmt.Println(profileBlock)
}

func PrintPerson(w io.Writer, person *model.Person) {
	printBasicInfo(person.BasicInfo)
	printAltNames(person.AltNames)
	printStatus(person.Status)
	printKinships(person.KinShips)
	printAssociations(removeDupliate(person.Associations))
	printPlaces(person.Places)
	printEntries(person.Entries)
	printPostings(person.Postings)
	printInstitutions(person.Institutions)

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
				joinTwoLang(safeValue(name.CNameTypeDescChn), safeValue(name.CNameTypeDesc)),
				joinTwoLang(name.CAltNameChn, safeValue(name.CAltName)),
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
				strconv.FormatInt(kinship.CKinID, 10),
				joinTwoLang(safeValue(kinship.CNameChn), safeValue(kinship.CName)),
				joinTwoLang(kinship.CKinrelChn, kinship.CKinrel),
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
				strconv.FormatInt(association.CAssocID, 10),
				joinTwoLang(safeValue(association.CNameChn), safeValue(association.CNameChn)),
				joinTwoLang(safeValue(association.CAssocTypeDescChn), safeValue(association.CAssocTypeShortDesc)),
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
				joinTwoLang(v.CStatusDescChn,
					v.CStatusDesc),
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
				joinTwoLang(safeValue(v.CNameChn), safeValue(v.CName)),
				joinTwoLang(safeValue(v.CAddrDescChn), safeValue(v.CAddrDesc)),
				joinTwoYear(safeValue(v.CFirstyear), safeValue(v.CLastyear)),
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
				joinTwoLang(v.CEntryTypeDescChn, v.CEntryTypeDesc),
				strconv.FormatInt(int64(v.CYear), 10),
				safeValue(v.CAge),
				safeValue(v.CExamRank),
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
				safeValue(v.COfficeChn),
				joinTwoLang(safeValue(v.CApptDescChn), safeValue(v.CApptDesc)),
				joinTwoYear(safeValue(v.CFirstyear), safeValue(v.CLastyear)),
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
				joinTwoLang(safeValue(v.CInstNameHz), safeValue(v.CInstNamePy)),
				joinTwoLang(safeValue(v.CBiRoleChn), safeValue(v.CBiRoleDesc)),
				joinTwoYear(safeValue(v.CBiBeginYear), safeValue(v.CBiEndYear)),
			})
	}

	printTable([]string{"Institution", "Role", "Period"}, rows)
}

// TODO: build a generic version
func removeDupliate(s []repository.GetAssociationByPersonIDRow) []repository.GetAssociationByPersonIDRow {
	set := make(map[int64]*repository.GetAssociationByPersonIDRow)
	for _, v := range s {
		set[v.CAssocID] = &v
	}

	res := make([]repository.GetAssociationByPersonIDRow, 0)

	for _, v := range set {
		res = append(res, *v)
	}

	return res
}
