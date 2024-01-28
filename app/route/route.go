package route

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Run(c *gin.Engine, db *gorm.DB) {
	base := c.Group("/")

	RouteAdmin(base, db)
	RouteUser(base,db)
}
