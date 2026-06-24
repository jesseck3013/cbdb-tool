package controller

type PersonField string

const (
	BasicInfo   PersonField = "info"
	AltName     PersonField = "altname"
	Entry       PersonField = "entry"
	Institution PersonField = "institution"
	Posting     PersonField = "posting"
	Status      PersonField = "status"
	Text        PersonField = "text"
	KinShip     PersonField = "kinship"
	Association PersonField = "association"
	Place       PersonField = "place"
)
