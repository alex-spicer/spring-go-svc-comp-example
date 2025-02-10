package server

import (
	"github.com/gin-gonic/gin"
	"product-go/internal/handler"
	"product-go/internal/service"
)

type Server struct {
	router *gin.Engine
}

func New() *Server {
	router := gin.Default()

	productService := service.NewProductService()
	productHandler := handler.NewProductHandler(productService)

	api := router.Group("/api")
	{
		products := api.Group("/products")
		{
			products.GET("", productHandler.GetAll)
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
