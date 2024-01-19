package models

type Order struct {
	ID         int     `json:"id" bson:"_id,omitempty"`
	User       string  `json:"userName" bson:"userName"`
	Product    string  `json:"productName" bson:"productName"`
	Quantity   int     `json:"quantity" bson:"quantity"`
	TotalPrice float64 `json:"totalPrice" bson:"totalPrice"`
}
