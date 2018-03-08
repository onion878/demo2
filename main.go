package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/static"
	m "./app"
	"os"
	"io"
)


func main() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./static/dist", true)))
	user := r.Group("/user")
	{
		m.User{}.UserInit(user)
	}
	base := r.Group("/base")
	{
		m.Base{}.UserInit(base)
	}


	r.Run() // listen and serve on 0.0.0.0:8080
}