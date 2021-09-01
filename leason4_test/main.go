package main

import "github.com/gin-gonic/gin"

/**
路由分组
*/
func main() {
	r := gin.Default()

	// 路由分组
	//r.GET("/posts", GetHandler)
	//r.POST("/posts", PostHandler)
	//// 删除id为1的文章
	//r.DELETE("/posts/1", DeleteHandler)

	p := r.Group("/posts")
	{
		// 路由分组，下面单独的/也可以省去不写
		p.GET("/", GetHandler)
		p.POST("/", PostHandler)
		// 删除id为1的文章
		p.DELETE("/1", DeleteHandler)
	}

	v1 := r.Group("/api/v1")
	{
		v1.POST("/login", loginEndpoint)
		v1.POST("/submit", submitEndpoint)
		v1.POST("/read", readEndpoint)
	}

	v2 := r.Group("/api/v2")
	{
		v2.POST("/login", loginEndpoint)
		v2.POST("/submit", submitEndpoint)
		v2.POST("/read", readEndpoint)
	}

	// 嵌套式路径
	api := r.Group("/api")
	{
		v3 := api.Group("/api/v3")
		{
			v3.POST("/login", loginEndpoint)
			v3.POST("/submit", submitEndpoint)
			v3.POST("/read", readEndpoint)
		}
	}

	err := r.Run()
	if err != nil {
		return
	}
}

func readEndpoint(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "readEndpoint",
	})
}

func submitEndpoint(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "submitEndpoint",
	})
}

func loginEndpoint(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "loginEndpoint",
	})
}

func DeleteHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "DELETE",
	})
}

func PostHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "POST",
	})
}

func GetHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "GET",
	})
}
