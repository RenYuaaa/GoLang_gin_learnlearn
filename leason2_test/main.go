package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	// 加载静态文件
	// 加载指定路径下所有的文件
	r.Static("/images", "./images")

	//
	r.StaticFS("/static", http.Dir("./static"))

	// 加载单独的静态文件
	r.StaticFile("index", "index.html")

	err := r.Run()
	if err != nil {
		return
	}
}
