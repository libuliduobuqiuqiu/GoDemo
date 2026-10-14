package routers

import (
	"godemo/internal/goweb/gogin/proxy/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterProxyRouter(group *gin.RouterGroup, allowedTargets []string) error {
	handler, err := handlers.NewProxyHandler(allowedTargets)
	if err != nil {
		return err
	}
	proxyGroup := group.Group("proxy")
	proxyGroup.Any("", handler)
	return nil
}
