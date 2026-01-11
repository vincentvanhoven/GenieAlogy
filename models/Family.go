package models

type Family struct {
	Id        *int `json:"id"`
	Person1Id *int `json:"person_1_id"`
	Person2Id *int `json:"person_2_id"`
	PositionX int  `json:"position_x"`
	PositionY int  `json:"position_y"`
}
