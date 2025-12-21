package models

import "database/sql"

type Family struct {
	Uuid        string         `json:"uuid"`
	Person1Uuid sql.NullString `json:"person_1_uuid"`
	Person2Uuid sql.NullString `json:"person_2_uuid"`
}
