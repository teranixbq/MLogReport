package route

import (
	"github.com/gin-gonic/gin"
	supabase "github.com/supabase-community/storage-go"
	"gorm.io/gorm"
)

func Run(c *gin.Engine, db *gorm.DB, sb *supabase.Client) {
	base := c.Group("/")

	RouteAdmin(base, db)
	RouteUser(base, db)
	RouteWeekly(base, db)
	RouteReport(base, db, sb)
}
