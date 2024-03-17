package router

import (
	"github.com/daviht7/goopportunitties/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {

	v1 := router.Group("/api/v1")
	{
		//show Opening
		v1.GET("opening", handler.ShowOpeningHandler)

		v1.POST("openings", handler.CreateOpeningHandler)

		v1.DELETE("openings", handler.DeleteOpeningHandler)

		v1.PUT("openings", handler.UpdateOpeningHandler)

		v1.GET("openings", handler.ListOpeningHandler)

	}

}
