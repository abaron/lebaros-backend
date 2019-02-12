package apiv1

import (
	"github.com/abaron/lebaros-backend/api/v1.0/auth"
	"github.com/abaron/lebaros-backend/api/v1.0/posts"
	"github.com/abaron/lebaros-backend/api/v1.0/product"
	"github.com/abaron/lebaros-backend/api/v1.0/stock"
	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/v1.0")
	{
		v1.GET("/ping", ping)
		auth.ApplyRoutes(v1)
		product.ApplyRoutes(v1)
		stock.ApplyRoutes(v1)
		posts.ApplyRoutes(v1)
	}
}
