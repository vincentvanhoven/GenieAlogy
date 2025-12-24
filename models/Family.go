package models

type Family struct {
	Uuid        string  `json:"uuid"`
	Person1Uuid *string `json:"person_1_uuid"`
	Person2Uuid *string `json:"person_2_uuid"`
}
