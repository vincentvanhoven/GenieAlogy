package models

type Person struct {
	Uuid           string  `json:"uuid"`
	Sex            string  `json:"sex"`
	Firstname      *string `json:"firstname"`
	Lastname       *string `json:"lastname"`
	Birthdate      *string `json:"birthdate"`
	Birthplace     *string `json:"birthplace"`
	FamilyUuid     *string `json:"family_uuid"`
	ProfilePicture *string `json:"profile_picture"`
	PositionX      int     `json:"position_x"`
	PositionY      int     `json:"position_y"`
}
