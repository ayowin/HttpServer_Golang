package main

import (
	"http-server-golang/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	var engine *gin.Engine
	engine = gin.Default()

	/* IndexController */
	var indexRouterGroup *gin.RouterGroup
	indexRouterGroup = engine.Group("/index")
	indexRouterGroup.Any("/", controller.IndexController.Index)

	/* UserController */
	var userRouterGroup *gin.RouterGroup
	userRouterGroup = engine.Group("/user")
	userRouterGroup.Any("/users", controller.UserController.Users)
	userRouterGroup.POST("/login", controller.UserController.Login)

	engine.Run(":80")
}
