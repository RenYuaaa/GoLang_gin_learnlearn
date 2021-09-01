package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// GET请求--获取所有文章信息
	r.GET("/posts", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// POST请求--创建一篇新文章
	r.POST("/posts", func(c *gin.Context) {
		c.String(http.StatusOK, "POST")
	})

	// PUT请求--修改一篇文章
	r.PUT("/posts/:id", func(c *gin.Context) {
		c.String(http.StatusOK, fmt.Sprintf("PUT id : %s", c.Param("id")))
	})

	// DELETE请求--删除一篇文章
	r.DELETE("/posts", func(c *gin.Context) {
		c.String(http.StatusOK, "DELETE")
	})

	// 任何方法都可以进入该请求
	r.Any("/users", func(c *gin.Context) {
		c.String(200, "any")
	})

	err := r.Run()
	if err != nil {
		return
	}
}
