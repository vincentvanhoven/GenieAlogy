package models

type Family struct {
	Id        *int `json:"id"`
	MaleId    *int `json:"male_id"`
	FemaleId  *int `json:"female_id"`
	PositionX int  `json:"position_x"`
	PositionY int  `json:"position_y"`
}
