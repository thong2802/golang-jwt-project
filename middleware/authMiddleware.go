package middleware

import "github.com/gin-gonic/gin"
// ở cái AuthMiddleware mình sẽ viết một middleware isAuth có chức năng bảo vệ những api cần bảo mật,
func Authenticate()gin.HandlerFunc  {
	return func(context *gin.Context) {

	}
}
