package main

import (
	"context"
	"log"

	"github.com/ladiesman2127/mongo-go/controllers"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	r := gin.Default()
	userController := controllers.NewUserController(getSession())

	r.GET("/user/:id", userController.GetUser)
	r.POST("user", userController.CreateUser)
	r.DELETE("/user/:id", userController.DeleteUser)

	log.Fatal(r.Run())
}

func getSession() *mongo.Client {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:2717"))
	if err != nil {
		log.Fatal(err)
	}
	return client
}
