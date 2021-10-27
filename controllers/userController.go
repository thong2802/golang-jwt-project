package controllers

import (
	"context"
	"fmt"
	"github.com/ducthong2802/golang-jwt-project/database"
	helper "github.com/ducthong2802/golang-jwt-project/helpers"
	"github.com/ducthong2802/golang-jwt-project/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
)

var userCollectoin *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()

//HashPassword encrypts user password
func HashPassword(password string) string  {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password),14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func VerifyPassword(userPassword string, providedPassword string)(bool, string)  {
	// use library bcrypt so sanh password da duoc băm va passord input co trung nhau k
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	msg := ""

	if err != nil {
		msg   =fmt.Sprintf("email of password is incorrect")
		check = false
	}
	return check, msg
}
//AuthController.js này sẽ bao gồm 2 controller login – thực hiện chức năng đăng nhập, tạo token và controller refreshToken – làm mới lại token khi hết hạn.
func SignUp() gin.HandlerFunc  {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(),100*time.Second)
		var user models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		validationErr := validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":validationErr.Error(),
			})
			return
		}
		//check email
		count, err := userCollectoin.CountDocuments(ctx, bson.M{"email" : user.Email})
		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":"error occured while checking for the email",
			})
		}
		password := HashPassword(*user.Password)
		user.Password = &password

		//check phone
		count, err = userCollectoin.CountDocuments(ctx, bson.M{"phone" : user.Phone})
		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":"error occured while checking for the email",
			})
		}

		//if count ton tai thi toast err
		if count > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":"this email or phone number already exists",
			})
		}

		// format save in database
		user.Creat_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		user.User_id = user.ID.Hex()
		token, refreshToken, _ :=  helper.GenerateAllTokens(*user.Email, *user.First_name, *user.Last_name, *user.User_type, *&user.User_id)
		user.Token = &token
		user.Refresh_token = &refreshToken

		// save vao database
		resultInsertionNumber, insertErr := userCollectoin.InsertOne(ctx, user)
		if  insertErr != nil{
			msg := fmt.Sprintf("User item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":msg,
			})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, resultInsertionNumber)
	}
}

// user login
func Login() gin.HandlerFunc  {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User
		var foundUser models.User

		if err := c.BindJSON(&user); err != nil{
			c.JSON(http.StatusBadRequest, gin.H{
				"err" : err.Error(),
			})
			return
		}

		// check email
		err := userCollectoin.FindOne(ctx, bson.M{"email" : user.Email}).Decode(&foundUser)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error" : "email or password is incorrect",
			})
			return
		}

			// xac minh password nhap dung khong
		 passwordIsValid, msg := VerifyPassword(*user.Password, *foundUser.Password)
		 defer cancel()
		 if passwordIsValid != true{
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": msg,
				})
			 return
		 }

		if foundUser.Email == nil  {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":"user not found",
			})
			return
		}
		c.JSON(http.StatusOK, foundUser)


		 // token
		 token, refrestoken,_ := helper.GenerateAllTokens(*foundUser.Email, *foundUser.First_name, *foundUser.Last_name, *foundUser.User_type, foundUser.User_id)
		 helper.UpdateAllTokens(token, refrestoken, foundUser.User_id)
		 err = userCollectoin.FindOne(ctx, bson.M{"user_id":foundUser.User_id}).Decode(&foundUser)

		if err !=  nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error" : err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, foundUser)
	}
}

func GetUsers()gin.HandlerFunc  {
	return func(context *gin.Context) {
		
	}
}

func GetUser() gin.HandlerFunc  {
	return func(c *gin.Context) {
		userId := c.Param("user_id")
		if err := helper.MatchUserTypeToUid(c, userId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":err.Error(),
			})
			return
		}
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

		var user models.User

	    err := userCollectoin.FindOne(ctx, bson.M{"user_id" : userId}).Decode(user)
		defer cancel()
		if err != nil{
			c.JSON(http.StatusBadRequest, gin.H{
				"error":err.Error(),
			})
		}
		c.JSON(http.StatusOK, user)
	}
}