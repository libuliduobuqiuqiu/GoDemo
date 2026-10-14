package routers

import (
	"fmt"
	"godemo/internal/goweb/gogin/proxy/public"
	"godemo/internal/goweb/gogin/proxy/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(proxyTargets []string, enableSSH bool) (*gin.Engine, error) {
	router := gin.New()
	router.Use(gin.Recovery())

	if gin.Mode() == gin.DebugMode {
		router.Use(gin.Logger())
	}
	router.GET("/healthz", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, public.ProxyResponse{
			Err:     http.StatusNotFound,
			Message: "The request url was not found on the server",
		})
	})

	// Register router
	rootGroup := router.Group("/netac/base")
	RegisterNetworkRouters(rootGroup)
	RegisterViewStatHandler(rootGroup)
	if err := RegisterProxyRouter(rootGroup, proxyTargets); err != nil {
		return nil, err
	}
	if enableSSH {
		if _, err := service.HostKeyCallback(); err != nil {
			return nil, err
		}
		RegisterSshRouters(rootGroup)
	}
	return router, nil
}

func StartProxyService(address string, port int, proxyTargets []string, enableSSH bool) error {
	router, err := NewRouter(proxyTargets, enableSSH)
	if err != nil {
		return err
	}
	return router.Run(fmt.Sprintf("%s:%d", address, port))
}
