package main

type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Person struct {
	Id         int      `json:"id"`
	Sex        string   `json:"sex"`
	Firstname  string   `json:"firstname"`
	Lastname   string   `json:"lastname"`
	Birthdate  string   `json:"birthdate"`
	Birthplace string   `json:"birthplace"`
	FamilyId   int      `json:"family_id"`
	Position   Position `json:"position"`
}

type Family struct {
	Id        int `json:"id"`
	Person1Id int `json:"person_1_id"`
	Person2Id int `json:"person_2_id"`
}

type SaveFile struct {
	People   []Person `json:"people"`
	Families []Family `json:"families"`
}
