package routers

import (
	"godemo/internal/goweb/gogin/proxy/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterViewStatHandler(group *gin.RouterGroup) {
	h := handlers.NewViewStatHandlers()
	group.Group("").GET("/view_stat", h.GetData)

}
