package data

import (
	"context"
	"ecommerce/gmr/interfaces"
	"ecommerce/gmr/models"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type AuthDataLayerImp struct {
	userCollection *mongo.Collection
	ctx            context.Context
}

func NewAuthDataLayerImp(collection *mongo.Collection, ctx context.Context) interfaces.AuthDataLayer {
	return &AuthDataLayerImp{
		userCollection: collection,
		ctx:            ctx,
	}
}

func (adl *AuthDataLayerImp) VerifyCredentials(username string, password string) (string, error) {
	var user models.User
	fmt.Println(username)
	err := adl.userCollection.FindOne(adl.ctx, bson.M{"name": username}).Decode(&user)

	if err != nil {
		fmt.Println("shardul3")
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}

	return user.Name, nil
}
