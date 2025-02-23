package server

import (
	"github.com/gin-gonic/gin"
	"log"
	"product-go/internal/handler"
	"product-go/internal/service"
)

type Server struct {
	router *gin.Engine
}

func myCustomMiddleware(c *gin.Context) {
	// do something before request is sent
	log.Println(c.Request.URL.Path)

	c.Next()

	// do something with response
	status := c.Writer.Status()
	log.Println(status)
}

func New() *Server {
	router := gin.Default()

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	productService := service.NewProductService()
	productHandler := handler.NewProductHandler(productService)

	api := router.Group("/api")
	{
		products := api.Group("/v1/products")
		{
			products.GET("", myCustomMiddleware, productHandler.GetAll)
			products.GET("/:id", productHandler.Get)
			products.POST("", productHandler.Create)
			products.PUT("/:id", productHandler.Update)
			products.DELETE("/:id", productHandler.Delete)
		}
	}

	return &Server{router: router}
}

func (s *Server) Start(addr string) error {
	return s.router.Run(addr)
}
