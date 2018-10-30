package controllers

import (
	"log"

	"github.com/airdb/beews/models"
	"time"

	_ "fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
	"github.com/gorilla/websocket"
)

type WebSocketController struct {
	beego.Controller
}

var upgrader = websocket.Upgrader{}

func (c *WebSocketController) Get() {

	ws, err := upgrader.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil)
	if err != nil {
		log.Fatal(err)
	}
	//defer ws.Close()

	clients[ws] = true

	//不断的广播发送到页面上
	// for {
	// 	//目前存在问题 定时效果不好 需要在业务代码替换时改为beego toolbox中的定时器
	// 	time.Sleep(time.Second * 3)
	// 	msg := models.Message{Message: "这是向页面发送的数据 " + time.Now().Format("2006-01-02 15:04:05")}
	// 	broadcast <- msg
	// }
	timeTask()

	//如果从 socket 中读取数据有误，我们假设客户端已经因为某种原因断开。我们记录错误并从全局的 “clients” 映射表里删除该客户端，这样一来，我们不会继续尝试与其通信。
	//另外，HTTP 路由处理函数已经被作为 goroutines 运行。这使得 HTTP 服务器无需等待另一个连接完成，就能处理多个传入连接。
	// for {
	// 	time.Sleep(time.Second * 3)

	// 	var msg models.Message // Read in a new message as JSON and map it to a Message object
	// 	err := ws.ReadJSON(&msg)
	// 	if err != nil {
	// 		log.Printf("页面可能断开啦 ws.ReadJSON error: %v", err)
	// 		delete(clients, ws)
	// 		break
	// 	} else {
	// 		fmt.Println("接受到从页面上反馈回来的信息 ", msg.Message)
	// 	}

	// 	broadcast <- msg
	// }

}

func timeTask() {

	timeStr := "0/3 * * * * *" //每隔3秒执行

	t1 := toolbox.NewTask("timeTask", timeStr, func() error {

		//todo do what you want
		msg := models.Message{Message: "这是向页面发送的数据 " + time.Now().Format("2006-01-02 15:04:05")}
		broadcast <- msg

		return nil
	})

	toolbox.AddTask("tk1", t1)
	toolbox.StartTask()
}
