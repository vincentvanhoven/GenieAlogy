package models

type Person struct {
	Id             *int    `json:"id"`
	Sex            string  `json:"sex"`
	Firstname      *string `json:"firstname"`
	Lastname       *string `json:"lastname"`
	Birthdate      *string `json:"birthdate"`
	Birthplace     *string `json:"birthplace"`
	FamilyId       *int    `json:"family_id"`
	ProfilePicture *string `json:"profile_picture"`
	PositionX      int     `json:"position_x"`
	PositionY      int     `json:"position_y"`
}
