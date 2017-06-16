package main

import "github.com/gin-gonic/gin"

func initRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/md5", md5Handler)
	return router
}

func main() {
	initRouter().Run()
}
