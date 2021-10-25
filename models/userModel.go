package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID				primitive.ObjectID  `bson:"_id"`
	First_name		*string				`json:"first_name" validate:"required, min=2, max=100"`
	Last_name		*string				`json:"last_name" validate:"required, min=2, max=100"`
	Password		*string				`json:"password" validate:"required, min=6, max=20"`
	Email			*string				`json:"email" validate:"email, required"`
	Phone			*string				`json:"phone" validate:"required"`
	Token			*string				`json:"token"`
	User_type		*string				`json:"user_type" validate:"required, eq=ADMIN|eq=USER"`
	Refresh_token	*string				`json:"refresh_token"`
	Creat_at		time.Time			`json:"creat_at"`
	Updated_at		time.Time			`json:"updated"`
	User_id			string				`json:"user_id"`
}
