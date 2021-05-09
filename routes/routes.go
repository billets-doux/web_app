package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web_app/logger"
	"web_app/settings"
)

func SetUp() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "ok")
	})
	r.GET("/version", func(context *gin.Context) {

		context.String(http.StatusOK, settings.Conf.AppConfig.Version)
	})
	return r
}
