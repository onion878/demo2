package app

import "github.com/gin-gonic/gin"

type Index struct {}

func (o Index) UserInit(router *gin.RouterGroup)  {
	router.GET("/save/:name", o.Save)
	router.GET("/delete", o.Delete)
	router.GET("/insert", o.Insert)
}

func (o Index) Save(c *gin.Context)  {
	c.String(200, c.Param("name"))
}

func (Index) Delete(c *gin.Context)  {
	c.String(200, "Delete")
}

func (Index) Insert(c *gin.Context)  {
	c.String(200, "Insert")
}