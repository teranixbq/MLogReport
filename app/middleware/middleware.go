package middleware

import (
	"mlogreport/utils/auth"
	"mlogreport/utils/helper"

	"github.com/gin-gonic/gin"
)

func IsRole(role_type string) gin.HandlerFunc{
	return func(c *gin.Context) {
		id,role,errExtract := auth.ExtractToken(c)
		if errExtract != nil {
			c.AbortWithStatusJSON(400,helper.ErrorResponse(errExtract.Error()))
			return
		}

		if role != role_type{
			c.AbortWithStatusJSON(401,helper.ErrorResponse("error : unauthorize"))
		}

		c.Set("id",id)
		c.Next()
	}

}