package handlers

import (
	"godemo/internal/goweb/gogin/proxy/public"
	"godemo/internal/goweb/gogin/proxy/service"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

// 1. 接受http请求，升级为ws请求
// 2. 根据传入的http请求参数，新建ssh配置，初始化client，根据client新建session
// 3. 发送建立伪终端, 建立通道将终端输出到ws中，将ws读取的命令输入到终端

func Terminal(ctx *gin.Context) {
	// 升级为Websocket请求
	upgrader := websocket.Upgrader{}

	ws, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		public.HandleErrJson(ctx, err)
		return
	}

	defer ws.Close()

	// 读取传入参数
	params := public.ProxyParams{}
	if err = ws.ReadJSON(&params); err != nil {
		public.HandleErrMessage(ws, err)
		return
	}

	terminalService := service.NewTerminalService()
	if err = terminalService.BuildTerminal(params, ws); err != nil {
		logrus.Error(err)
	}

}
