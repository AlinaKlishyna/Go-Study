package api

import (
	"github.com/iotique-srl/health-monitor/core"

	"github.com/iotique-srl/health-monitor/internal/api/auth"
	exec "github.com/iotique-srl/health-monitor/internal/graphql/exec"
	resolver "github.com/iotique-srl/health-monitor/internal/graphql/resolver"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	authHandler := &auth.AuthHandler{}
	router.Use(authHandler.JWTMiddleware())

	protected := router.Group("/")

	authorization := router.Group("/auth")
	authorization.POST("/login", authHandler.LoginHandler)

	endpoints := protected.Group("/endpoints")
	endpointHandler := &EndpointHandler{}
	endpointHandler.Update(endpoints)
	endpointHandler.Create(endpoints)
	endpointHandler.List(endpoints)
	endpointHandler.Delete(endpoints)

	results := protected.Group("/checkLast")
	resultsHandler := &ResultsHandler{}
	resultsHandler.GetAscID(results)
	resultsHandler.List(results)

	users := protected.Group("/users")
	usersHandler := &UsersHandler{}
	usersHandler.Create(users)
	usersHandler.Update(users)
	usersHandler.GetUser(users)
	usersHandler.List(users)
	usersHandler.Delete(users)

	setupGraphQLRoute(protected)

	return router
}

func setupGraphQLRoute(group *gin.RouterGroup) {
	srv := handler.New(exec.NewExecutableSchema(exec.Config{
		Resolvers: &resolver.Resolver{},
	}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.Use(extension.Introspection{})

	group.POST("/query", func(c *gin.Context) {
		domain := c.MustGet("domain").(*core.Domain)

		ctx := core.WithDomainContext(c.Request.Context(), domain)

		srv.ServeHTTP(c.Writer, c.Request.WithContext(ctx))
	})
}
