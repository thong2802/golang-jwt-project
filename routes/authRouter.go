package routes

import (
	"github.com/gin-gonic/gin"
	 controller "github.com/ducthong2802/golang-jwt-project/controllers"
)

func AuthRoutes(incomingRoutes *gin.Engine)  {
	incomingRoutes.POST("user/signup", controller.SignUp())
	incomingRoutes.POST("user/login", controller.Login())
}
