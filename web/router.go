package web

import "github.com/gin-gonic/gin"
import home "gin-rest-api/controllers"
import user "gin-rest-api/controllers"
import blog "gin-rest-api/controllers"

// Register routers
func Register(engine *gin.Engine)  {
	engine.GET("/", home.Index)
	engine.GET("/rest/latestBlog", blog.LatestBlog)
	engine.GET("/rest/info", home.Info)
	engine.POST("/signup", user.Signup)
	engine.POST("/signin", user.Signin)
}