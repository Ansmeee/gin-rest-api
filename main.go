package main

import (
	router "gin-rest-api/web"
	"github.com/gin-gonic/gin"
)

func main()  {
	// initialize app
	app := gin.Default()

	// register routers
	router.Register(app)

	// run app
	app.Run("localhost:9090")
}
