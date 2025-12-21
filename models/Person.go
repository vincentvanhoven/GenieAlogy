package models

import "database/sql"

type Person struct {
	Uuid           string         `json:"uuid"`
	Sex            string         `json:"sex"`
	Firstname      sql.NullString `json:"firstname"`
	Lastname       sql.NullString `json:"lastname"`
	Birthdate      sql.NullString `json:"birthdate"`
	Birthplace     sql.NullString `json:"birthplace"`
	FamilyUuid     sql.NullString `json:"family_uuid"`
	ProfilePicture sql.NullString `json:"profile_picture"`
	Position       Position       `json:"position"`
}
