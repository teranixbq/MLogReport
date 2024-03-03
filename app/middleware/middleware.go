package middleware

import (
	"mlogreport/utils/auth"
	"mlogreport/utils/helper"
	"strconv"

	"github.com/gin-gonic/gin"
)

func IsRole(roleType string) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, role, errExtract := auth.ExtractToken(c)
		if errExtract != nil {
			c.AbortWithStatusJSON(400, helper.ErrorResponse(errExtract.Error()))
			return
		}

		if role != roleType {
			c.AbortWithStatusJSON(401, helper.ErrorResponse("error : unauthorize"))
		}

		c.Set("id", id)
		c.Next()
	}

}

func Pagination() gin.HandlerFunc {
	return func(c *gin.Context) {
		pageStr := c.Query("page")
		limitStr := c.Query("limit")

		var page, limit int
		var err error

		if pageStr != "" {
			if page, err = strconv.Atoi(pageStr); err != nil {
				c.AbortWithStatusJSON(400, helper.ErrorResponse("error : failed to convert page to integer"))
				return
			}
		}

		if limitStr != "" {
			if limit, err = strconv.Atoi(limitStr); err != nil {
				c.AbortWithStatusJSON(400, helper.ErrorResponse("error : failed to convert limit to integer"))
				return
			}
		}

		c.Set("pagination", gin.H{
			"page":  page,
			"limit": limit,
		})

		c.Next()
	}
}
