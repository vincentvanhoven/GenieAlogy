package models

type SaveFile struct {
	People   []Person `json:"people"`
	Families []Family `json:"families"`
}
