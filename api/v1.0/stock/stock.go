package stock

import (
	"github.com/abaron/lebaros-backend/lib/middlewares"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	product := r.Group("/stock")
	{
		product.POST("/in", middlewares.Authorized, in)
		product.POST("/out", middlewares.Authorized, out)
	}
}
