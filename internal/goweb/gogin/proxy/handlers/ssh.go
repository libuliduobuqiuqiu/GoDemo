package handlers

import (
	"godemo/internal/goweb/gogin/proxy/public"
	"godemo/internal/goweb/gogin/proxy/service"

	"github.com/gin-gonic/gin"
)

func ExecSshCommand(ctx *gin.Context) {
	var (
		params public.ProxySshParams
		cmdRes map[string]string
	)

	if err := ctx.BindJSON(&params); err != nil {
		public.HandleErrJson(ctx, err)
		return
	}

	sshService := service.NewSshService()
	cmdRes, err := sshService.ExecCommand(params)
	if err != nil {
		public.HandleErrJson(ctx, err)
		return
	}
	public.HandleSuccessJson(ctx, cmdRes)
}
