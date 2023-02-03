package main

import (
	"github.com/gin-gonic/gin"
	"github.com/soohyun-lee/side_project_overseas/initializers"
	"github.com/soohyun-lee/side_project_overseas/routes"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {
	r := gin.Default()

	r.POST("/login", routes.UserLogin)
	r.POST("/board", routes.BoardCreate)
	r.GET("/board/list", routes.BoardList)
	r.Run()
}
