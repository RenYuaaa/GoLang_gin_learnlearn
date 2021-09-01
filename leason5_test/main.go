package main

import "github.com/gin-gonic/gin"

type Person struct {
	Id   int    `uri:"id"`
	Name string `uri:"name"`
}

func main() {
	r := gin.Default()

	r.GET("/posts", func(c *gin.Context) {
		c.String(200, "1")
	})

	//r.GET("/:id", func(c *gin.Context) {
	//	id := c.Param("id")
	//	c.JSON(200, gin.H{
	//		"id": id,
	//	})
	//})

	r.GET("/:id/:name", func(c *gin.Context) {
		var person Person
		if err := c.ShouldBindUri(&person); err != nil {
			c.Status(404)
			return
		}
		c.JSON(200, gin.H{
			"name": person.Name,
			"id":   person.Id,
		})
	})

	// 删除数据
	r.DELETE("/posts/:id", func(c *gin.Context) {

	})

	err := r.Run()
	if err != nil {
		return
	}
}
