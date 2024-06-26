package controllers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ladiesman2127/mongo-go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userController struct {
	mongoClient *mongo.Client
}

func NewUserController(mongoClient *mongo.Client) *userController {
	return &userController{mongoClient}
}

func (uc *userController) GetUser(ctx *gin.Context) {
	id, err := primitive.ObjectIDFromHex(ctx.Params.ByName("id"))
	if err != nil {
		ctx.Writer.WriteHeader(http.StatusNotFound)
		return
	}
	user := models.User{}
	users := uc.mongoClient.Database("mc_db").Collection("users")
	users.FindOne(context.Background(), bson.M{"_id": id}).Decode(&user)

	jUser, err := json.Marshal(user)
	if err != nil {
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		return
	}
	ctx.Writer.Header().Set("Content-Type", "application/json")
	ctx.Writer.WriteHeader(http.StatusOK)
	ctx.Writer.Write(jUser)
}

func (uc *userController) CreateUser(ctx *gin.Context) {
	user := models.User{}
	users := uc.mongoClient.Database("mc_db").Collection("users")
	json.NewDecoder(ctx.Request.Body).Decode(&user)
	user.Id = primitive.NewObjectID()
	if _, err := users.InsertOne(context.Background(), user); err != nil {
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		return
	}
	jUser, err := json.Marshal(user)
	if err != nil {
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		return
	}
	ctx.Writer.Header().Set("Content-Type", "application/json")
	ctx.Writer.WriteHeader(http.StatusOK)
	ctx.Writer.Write(jUser)
}

func (uc *userController) DeleteUser(ctx *gin.Context) {
	id, err := primitive.ObjectIDFromHex(ctx.Params.ByName("id"))
	if err != nil {
		ctx.Writer.WriteHeader(http.StatusNotFound)
		return
	}

	users := uc.mongoClient.Database("mc_db").Collection("users")
	if _, err := users.DeleteOne(context.Background(), bson.M{"_id": id}); err != nil {
		ctx.Writer.WriteHeader(http.StatusNotFound)
		return
	}
	ctx.Writer.WriteHeader(http.StatusOK)
}
