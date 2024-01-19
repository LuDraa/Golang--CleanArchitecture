package models

type Products struct {
	ID          int     `json:"id" bson:"_id"`
	Name        string  `json:"name" bson:"name"`
	Type        string  `json:"type" bson:"type"`
	Cost        float64 `json:"cost" bson:"cost"`
	IsAvailable bool    `json:"isAvailable" bson:"isAvailable"`
}
