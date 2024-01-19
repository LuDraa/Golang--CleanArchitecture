package main

import (
	"context"
	"ecommerce/gmr/controllers"
	"ecommerce/gmr/data"
	"ecommerce/gmr/interfaces"
	"ecommerce/gmr/middleware"
	"ecommerce/gmr/services"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server      *gin.Engine
	ad          interfaces.AuthDataLayer
	as          interfaces.AuthServiceLayer
	ac          controllers.AuthController
	ud          interfaces.UserDataLayer
	us          interfaces.UserService
	uc          controllers.UserController
	ctx         context.Context
	userc       *mongo.Collection
	mongoclient *mongo.Client
	err         error
)

func init() {
	ctx = context.TODO()

	mongoconn := options.Client().ApplyURI("mongodb://localhost:27017")
	mongoclient, err = mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal("error while connecting with mongo", err)
	}
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("error while trying to ping mongo", err)
	}

	fmt.Println("mongo connection established")

	userc = mongoclient.Database("userdb").Collection("users")
	ud = data.NewUserDataLayerImpl(userc, ctx)
	us = services.NewUserService(ud)
	uc = controllers.New(us)

	ad = data.NewAuthDataLayerImp(userc, ctx)
	as = services.NewAuthService(ad)
	ac = controllers.NewAuthController(as)

	server = gin.Default()
}

func main() {

	defer mongoclient.Disconnect(ctx)
	server.Use(middleware.AuthMiddleware(as))

	basepath := server.Group("/")
	uc.RegisterUserRoutes(basepath)

	ac := controllers.NewAuthController(as)
	ac.RegisterAuthRoutes(basepath)

	log.Fatal(server.Run(":9090"))

}
