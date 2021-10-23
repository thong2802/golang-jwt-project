package controllers

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	helper "github.com/ducthong2802/golang-jwt-project/helpers"
	"github.com/ducthong2802/golang-jwt-project/models"
	"github.com/ducthong2802/golang-jwt-project/database"
	"golang.org/x/crypto/bcrypt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

)

var userCollectoin *mongo.Collection = database.OpenCollection(database.Client, "uset")
var validate = validator.New()

//HashPassword encrypts user password
func HashPassword()  {
   
}

func VerifyPassword()  {
	
}

func SignUp() gin.HandlerFunc  {
	return func(context *gin.Context) {

	}
}

func Login() gin.HandlerFunc  {
	return func(context *gin.Context) {

	}
}

func GetUsers()gin.HandlerFunc  {
	return func(context *gin.Context) {
		
	}
}

func GetUser() gin.HandlerFunc  {
	return func(context *gin.Context) {
		
	}
}