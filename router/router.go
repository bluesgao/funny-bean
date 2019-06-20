package router

import (
	. "funny-bean/handler"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", IndexUsers) //http://localhost:8806

	//路由群组
	users := router.Group("api/v1/user")
	{
		users.GET("", GetAll)         //http://localhost:8806/api/v1/user
		users.POST("/add", AddUser)  //http://localhost:8806/api/v1/user/add
		users.GET("/get/:id", GetOne) //http://localhost:8806/api/v1/user/get/5
		users.POST("/update", UpdateUser)
		users.POST("/del", DelUser)
	}

	return router
}
