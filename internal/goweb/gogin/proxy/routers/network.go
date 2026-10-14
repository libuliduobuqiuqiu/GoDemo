package routers

import (
	"godemo/internal/goweb/gogin/proxy/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterNetworkRouters(group *gin.RouterGroup) {
	networkGroup := group.Group("")
	networkGroup.GET("ping", handlers.Ping)
	networkGroup.GET("traceroute", handlers.TraceRoute)
	networkGroup.POST("upload", handlers.ParseFormData)
}
