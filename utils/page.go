package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Pagination(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func (db *gorm.DB) *gorm.DB {
		
		page, _ := strconv.Atoi(c.Query("page"))
		if page == 0 {
			page = 1
		}

		perPage, _ := strconv.Atoi(c.Query("per_page"))
		if (perPage == 0) {
			perPage = 10
		}

		offSet := (page - 1) * perPage

		return db.Offset(offSet).Limit(perPage)
	}
}