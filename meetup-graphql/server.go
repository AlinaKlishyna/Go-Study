package main

import (
	"meetup-graphql/graphql/exec"
	"meetup-graphql/graphql/middleware"
	"meetup-graphql/graphql/resolver"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/vektah/gqlparser/v2/ast"
)

func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema и Config находятся в файле exec/generated.go
	// Resolver находится в файле resolver.go
	srv := handler.New(exec.NewExecutableSchema(exec.Config{Resolvers: &resolver.Resolver{}}))

	// Настройка сервера
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	return func(ctx *gin.Context) {
		srv.ServeHTTP(ctx.Writer, ctx.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	srv := playground.Handler("GraphQL playground", "/query")

	return func(ctx *gin.Context) {
		srv.ServeHTTP(ctx.Writer, ctx.Request)
	}
}

func main() {
	router := gin.Default()
	// Подключаем middleware ДО обработки запросов
	// Теперь каждый запрос будет упаковывать gin.Context в обычный context.Context,
	// чтобы потом мы могли доставать его внутри GraphQL резолверов.
	router.Use(middleware.GinContextToContextMiddleware())

	router.POST("/query", graphqlHandler())
	router.GET("/", playgroundHandler()) // Playground – обычная HTML страничка для тестирования

	router.Run() // Запускаем HTTP сервер (по умолчанию на :8080)
}
