package route

import (
	"curd/controller"

	"github.com/gin-gonic/gin"
)

func Start(port string) {
	app := gin.Default()
	app.POST("/user/add", controller.AddUser)
	app.POST("/user/edit", controller.UpdateUser)
	app.GET("/user/find", controller.FindUser)
	app.GET("/user/all", controller.FindAllUser)
	app.POST("/user/delete", controller.DeteleUser)
	app.Run(port)

}
