package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo}
}

func (s *Service) GetProducts(ctx *gin.Context) {
	// extracting query string from request
	query := ctx.Request.URL.RawQuery

	// calling DB to get filtered result
	products, err := s.repo.GetProducts(ctx, query)

	// Handling errors with response code 500 (Internal server error)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error getting products",
		})
		return
	}

	// sending response with status code 200
	ctx.JSON(http.StatusOK, GetProductsRes{
		Products: products,
		Count:    len(products),
	})
}
