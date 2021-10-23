package golang_jwt_project

import (
	routes "github.com/ducthong2802/golang-jwt-project/routes"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)
	
	router.GET("/api-1", func(context *gin.Context) {
		 context.JSONP(http.StatusOK, gin.H{
			 "success" : "Access granted for api-1",
		 })
	})
	router.GET("/api-2", func(context *gin.Context) {
		context.JSONP(http.StatusOK, gin.H{
			"success" : "Access granted for api-2",
		})
	})

	router.Run(":" + port)
}
