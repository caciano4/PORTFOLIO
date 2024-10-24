package handlers

import (
	"github.com/caciano/portfolio/internal/handlers/users"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRoutes(router *gin.Engine, db *gorm.DB) {
	users.Init(router, db)
}
