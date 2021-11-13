package main

import (
	"github.com/cpunion/graph/app/graph"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	ginRoute := gin.Default()
	ginRoute.Use(cors.Default())

	graphql := ginRoute.Group("graphql")
	{
		graph.InitRoutes(graphql)
	}

	ginRoute.Run(":8080")
}
