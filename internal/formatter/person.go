package formatter

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/jesseck3013/cbdb-tool/internal/repository"
)

const FALLBACK = "Unkown"

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
