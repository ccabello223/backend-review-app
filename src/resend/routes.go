package resend

import (
	"time"

	"github.com/NetKBs/backend-reviewapp/src/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	resend := router.Group("/code", middlewares.AuthMiddleware())
	{

		resend.POST("/generate", middlewares.RateLimitMiddleware(1, 30*time.Second), generateVerificationCodeController)
		resend.POST("/verify", verifyVerificationCodeController)

	}
}
