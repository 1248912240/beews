package routers

import (
	"github.com/airdb/beews/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/ws", &controllers.WebSocketController{})
}
