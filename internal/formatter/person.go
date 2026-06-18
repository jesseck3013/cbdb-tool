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
			Align(lipgloss.Center)

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
				joinTwoLang(name.CAltNameChn, safeValue(name.CAltName)),
				joinTwoLang(safeValue(name.CNameTypeDescChn), safeValue(name.CNameTypeDesc)),
			})
	}

	t := newTable([]string{"Alternative Name", "Type"}, rows)
	lipgloss.Println(t)
}
