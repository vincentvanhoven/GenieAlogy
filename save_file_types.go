package main

type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Person struct {
	Uuid       string   `json:"uuid"`
	Sex        string   `json:"sex"`
	Firstname  string   `json:"firstname"`
	Lastname   string   `json:"lastname"`
	Birthdate  string   `json:"birthdate"`
	Birthplace string   `json:"birthplace"`
	FamilyUuid string   `json:"family_uuid"`
	Position   Position `json:"position"`
}

type Family struct {
	Uuid        string `json:"uuid"`
	Person1Uuid string `json:"person_1_uuid"`
	Person2Uuid string `json:"person_2_uuid"`
}

type SaveFile struct {
	People   []Person `json:"people"`
	Families []Family `json:"families"`
}
