package helpers

import (
	"context"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/ducthong2802/golang-jwt-project/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"
	"time"
)

var userCollection*mongo.Collection = database.OpenCollection(database.Client, "user")
var SECRET_KEY string = os.Getenv("SECRET_KEY")

// Định nghĩa những thông tin của user mà bạn muốn lưu vào token ở đây
type SignedDetails struct {
	Email      string
	First_name string
	Last_name  string

	        string
	User_type  string
	jwt.StandardClaims
}
//generateToken – tạo token
func GenerateAllTokens(email string, firstName string, lastName string, userType string, uid string) (signedToken string, signedRefreshToken string, err error){
	// Thực hiện ký và tạo token
	claims := &SignedDetails {
	  Email : email,
	  First_name: firstName,
	  Last_name: lastName,
	  Uid: uid,
	  User_type: userType,
	  StandardClaims : jwt.StandardClaims{
		  ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
	  },
  }

  refreshClaims := &SignedDetails{
	  StandardClaims: jwt.StandardClaims{
		  ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
	  },
  }

 token, err := jwt.NewWithClaims(jwt.SigningMethodES256, claims).SignedString([]byte(SECRET_KEY))
 refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodES256, refreshClaims).SignedString([]byte(SECRET_KEY))

	if err != nil {
		log.Panic(err)
		return
	}


 return token, refreshToken, err
}


// UpdateAllTokens
func UpdateAllTokens(signedToken string, signedRefreshToken string, userId string)  {
	 var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	 var updateObj primitive.D

}
// verifyToken – xác minh token có hợp lệ hay không.

