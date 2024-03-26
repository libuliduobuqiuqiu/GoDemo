package gowebsockets

import (
	"fmt"
	"net/http"

	"github.com/beego/beego/v2/server/web"
	"github.com/gorilla/websocket"
)

func init() {
	web.CtrlGet("/chat/join", (*ChatController).Join)
}

type ChatController struct {
	web.Controller
}

func NewWsError(prefix string, err error) map[string]interface{} {
	errResp := map[string]interface{}{
		"code": http.StatusInternalServerError,
	}

	if err != nil {
		errResp["msg"] = fmt.Sprintf("%s: %s", prefix, err.Error())
	} else {
		errResp["msg"] = prefix
	}
	return errResp
}

func (c *ChatController) Join() {
	uname := c.GetString("name")

	// 处理websocket连接
	upgrader := websocket.Upgrader{}
	ws, err := upgrader.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil)
	if err != nil {
		errResp := NewWsError("Create Connection Failed", err)
		c.Ctx.Output.Status = http.StatusInternalServerError
		c.Ctx.JSONResp(errResp)
		return
	}

	defer ws.Close()

	// 验证用户名是否重复
	if err := isUserExist(uname); err != nil {
		ws.WriteMessage(websocket.TextMessage, []byte(err.Error()))
		c.Ctx.Output.Status = http.StatusBadRequest
		return
	}

	// 发送登录成功消息
	ws.WriteMessage(websocket.TextMessage, []byte("Welcome join the ChatRoom."))

	// 订阅用户通道注册用户
	join(Subscriber{Name: uname, Conn: ws})
	defer leave(uname)

	// 监听接受用户输入
	for {
		_, input, err := ws.ReadMessage()
		if err != nil {
			errResp := NewWsError("Read Message Failed", err)
			c.Ctx.Output.Status = http.StatusInternalServerError
			c.Ctx.JSONResp(errResp)
			return
		}
		publish <- Event{EventType: EventMessage, Name: uname, Message: string(input)}
	}
}
