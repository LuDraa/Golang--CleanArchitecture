package models

import "golang.org/x/crypto/bcrypt"

type Address struct {
	State   string `json:"state" bson:"state"`
	City    string `json:"city" bson:"city"`
	Country string `json:"country" bson:"country"`
	Pincode int    `json:"pincode" bson:"pincode"`
}

type User struct {
	ID       int     `json:"id" bson:"_id"`
	Name     string  `json:"name" bson:"name"`
	Age      int     `json:"age" bson:"age"`
	Phone    string  `json:"phone" bson:"phone"`
	Email    string  `json:"email" bson:"email"`
	Address  Address `json:"address" bson:"address"`
	Password string  `json:"password" bson:"password"`
	IsAdmin  bool    `json:"isAdmin" bson:"isAdmin"`
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}
