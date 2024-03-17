package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {

	v1 := router.Group("/api/v1")
	{
		//show Opening
		v1.GET("opening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "GET opening",
			})
		})

		v1.POST("openings", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "POST opening",
			})
		})

		v1.DELETE("openings", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "DELETE opening",
			})
		})

		v1.PUT("openings", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "PUT opening",
			})
		})

		v1.GET("openings", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "GET openings",
			})
		})

	}

}
