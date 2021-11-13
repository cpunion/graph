package graph

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/cpunion/graph/app/graph/generated"
	"github.com/gin-gonic/gin"
)

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler(basePath string) gin.HandlerFunc {
	h := playground.Handler("GraphQL", basePath+"/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func InitRoutes(gin *gin.RouterGroup) {

	gin.GET("/", playgroundHandler(gin.BasePath()))
	gin.POST("/query", graphqlHandler())
}
