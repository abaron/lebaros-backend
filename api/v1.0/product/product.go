package product

import (
	"github.com/abaron/lebaros-backend/lib/middlewares"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	product := r.Group("/product")
	{
		product.POST("/import-exists", middlewares.Authorized, importExists)
		product.GET("/stocks", middlewares.Authorized, stocks)
	}
}
