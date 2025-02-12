package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"product-go/internal/model"
	"product-go/internal/service"
	"strconv"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(service *service.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) GetAll(c *gin.Context) {
	products := h.service.GetAll()
	c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) Get(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if product, exists := h.service.Get(id); exists {
		c.JSON(http.StatusOK, product)
	} else {
		c.Status(http.StatusNotFound)
	}
}

func (h *ProductHandler) Create(c *gin.Context) {
	var product model.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created := h.service.Create(product)
	c.JSON(http.StatusCreated, created)
}

func (h *ProductHandler) Update(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var product model.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if updated, exists := h.service.Update(id, product); exists {
		c.JSON(http.StatusOK, updated)
	} else {
		c.Status(http.StatusNotFound)
	}
}

func (h *ProductHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if h.service.Delete(id) {
		c.Status(http.StatusOK)
	} else {
		c.Status(http.StatusNotFound)
	}
}
