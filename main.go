package main

import (
	"theraq/router"
	"github.com/gin-gonic/gin"
	
	
)

func main() {

	
	
	
	
	
	ginRouter := gin.Default()
	router.RegisterRoutes(ginRouter)
	
}
