package formatter

import (
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"

	"charm.land/lipgloss/v2"
	"charm.land/lipgloss/v2/table"
	"github.com/jesseck3013/cbdb-tool/internal/model"
	"github.com/jesseck3013/cbdb-tool/internal/repository"
)

type PersonDisplay struct {
	CPersonid     int64  `json:"c_personid"`
	CName         string `json:"c_name"`
	CNameChn      string `json:"c_name_chn"`
	CMingzi       string `json:"c_mingzi"`
	CFemale       string `json:"c_female"`
	CBirthyear    string `json:"c_birthyear"`
	CDeathyear    string `json:"c_deathyear"`
	CDynasty      string `json:"c_dynasty"`
	CDynastyChn   string `json:"c_dynasty_chn"`
	CChoronymDesc string `json:"c_choronym_desc"`
	CChoronymChn  string `json:"c_choronym_chn"`
}

func newPersonDisplay(p repository.GetPersonByIDRow) PersonDisplay {
	return PersonDisplay{
		CPersonid:     p.CPersonid,
		CName:         valueOrFallBack(p.CName, FALLBACK),
		CNameChn:      valueOrFallBack(p.CNameChn, FALLBACK),
		CMingzi:       valueOrFallBack(p.CMingzi, FALLBACK),
		CFemale:       valueOrFallBack(p.CFemale, FALLBACK),
		CBirthyear:    valueOrFallBack(p.CBirthyear, FALLBACK),
		CDeathyear:    valueOrFallBack(p.CDeathyear, FALLBACK),
		CDynasty:      valueOrFallBack(p.CDynasty, FALLBACK),
		CDynastyChn:   valueOrFallBack(p.CDynastyChn, FALLBACK),
		CChoronymDesc: valueOrFallBack(p.CChoronymDesc, FALLBACK),
		CChoronymChn:  valueOrFallBack(p.CChoronymChn, FALLBACK),
	}
}

func Person(p repository.GetPersonByIDRow) {
	person := newPersonDisplay(p)
	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprintln(tw, "Person Profile:")
	fmt.Fprintf(tw, " ID:\t%d\n", person.CPersonid)
	fmt.Fprintf(tw, " Name:\t%s\n", person.CName)
	fmt.Fprintf(tw, " Name (Chinese):\t%s\n", person.CNameChn)
	fmt.Fprintf(tw, " Dynasty:\t%s\n", person.CDynasty)
	fmt.Fprintf(tw, " Dynasty (Chinese):\t%s\n", person.CDynastyChn)
	fmt.Fprintf(tw, " Birth Year:\t%s\n", person.CBirthyear)
	fmt.Fprintf(tw, " Death Year:\t%s\n", person.CDeathyear)
	fmt.Fprintf(tw, " Choronym:\t%s\n", person.CChoronymDesc)
	fmt.Fprintf(tw, " Choronym (Chinese):\t%s\n", person.CChoronymChn)

	tw.Flush()
}

var (
	primaryColor = lipgloss.Color("#c7a587")

	labelStyle = lipgloss.NewStyle().
			Width(20).
			Bold(true).
			Foreground(primaryColor)

	// Style for valid data values
	valueStyle = lipgloss.NewStyle().
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
			Underline(true).
			MarginBottom(1)

	tableHeaderStyle = lipgloss.NewStyle().
				Foreground(primaryColor).
				Bold(true).
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

func printBasicInfo(info repository.GetPersonByIDRow) {
	fmt.Println(headerStyle.Render("Person Profile:"))
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

func PrintPerson(person *model.Person) {
	printBasicInfo(person.BasicInfo)
	printAltNames(person.AltNames)
	printStatus(person.Status)
	printKinships(person.KinShips)
	//printAssociations(person.Associations)
	printPlaces(person.Places)
	printEntries(person.Entries)
	printPostings(person.Postings)

}

func newTable(headers []string, rows [][]string) *table.Table {
	return table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(primaryColor)).
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
}

func printAltNames(names []repository.GetAltnamesByPersonIDRow) {
	rows := [][]string{}

	for _, name := range names {
		rows = append(rows,
			[]string{
				joinTwoLang(safeValue(name.CNameTypeDescChn), safeValue(name.CNameTypeDesc)),
				joinTwoLang(name.CAltNameChn, safeValue(name.CAltName)),
			})
	}

	t := newTable([]string{"Type", "Alternative Name"}, rows)
	lipgloss.Println(t)
}

func printKinships(kinships []repository.GetPersonKinShipByPersonIDRow) {
	rows := [][]string{}

	for _, kinship := range kinships {
		rows = append(rows,
			[]string{
				joinTwoLang(kinship.CKinrelChn, kinship.CKinrel),
				joinTwoLang(safeValue(kinship.CNameChn), safeValue(kinship.CName)),
			})
	}

	t := newTable([]string{"Kinship Type", "Name"}, rows)
	lipgloss.Println(t)
}

func printAssociations(associations []repository.GetAssociationByPersonIDRow) {
	rows := [][]string{}

	for _, association := range associations {
		rows = append(rows,
			[]string{
				joinTwoLang(safeValue(association.CAssocTypeDescChn), safeValue(association.CAssocTypeDesc)),
				joinTwoLang(safeValue(association.CNameChn), safeValue(association.CNameChn)),
			})
	}

	t := newTable([]string{"Association Type", "Name"}, rows)
	lipgloss.Println(t)
}

func printStatus(status []repository.GetStatusByPersonIDRow) {
	rows := [][]string{}

	for _, v := range status {
		rows = append(rows,
			[]string{
				joinTwoLang(v.CStatusDescChn,
					v.CStatusDesc),
			})
	}

	t := newTable([]string{"Status"}, rows)
	lipgloss.Println(t)
}

func printPlaces(places []repository.GetPlaceByPersonIDRow) {
	rows := [][]string{}

	for _, v := range places {
		rows = append(rows,
			[]string{
				joinTwoLang(safeValue(v.CNameChn), safeValue(v.CName)),
				joinTwoLang(safeValue(v.CAddrDescChn), safeValue(v.CAddrDesc)),
				joinTwoYear(safeValue(v.CFirstyear), safeValue(v.CLastyear)),
			})
	}

	t := newTable([]string{"Place", "Type", "Period"}, rows)
	lipgloss.Println(t)
}

func printEntries(entries []repository.GetEntryByPersonIDRow) {
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

	t := newTable([]string{"Mode of Entry", "Year", "Age", "Exam Rank"}, rows)
	lipgloss.Println(t)
}

func printPostings(postings []repository.GetPostingByPersonIDRow) {
	rows := [][]string{}

	for _, v := range postings {
		rows = append(rows,
			[]string{
				safeValue(v.COfficeChn),
				joinTwoLang(safeValue(v.CApptDescChn), safeValue(v.CApptDesc)),
				joinTwoYear(safeValue(v.CFirstyear), safeValue(v.CLastyear)),
			})
	}

	t := newTable([]string{"Office", "Type", "Period"}, rows)
	lipgloss.Println(t)
}
