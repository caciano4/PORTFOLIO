package users

import (
	"net/http"

	"github.com/caciano/portfolio/internal/services/users"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Init(router *gin.Engine, db *gorm.DB) {
	router.GET("/users", func(c *gin.Context) {
		HandleGetAllUsers(c, db)
	})
	router.GET("/users/:id", func(c *gin.Context) {
		HandleGetUserById(c, db)
	})
}

func HandleGetUserById(c *gin.Context, db *gorm.DB) {
	user, err := users.GetUserById(db, c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user, "message": "User fetched successfully"})
}

// HandleGetAllUsers is the handler for the GET /users endpoint
func HandleGetAllUsers(c *gin.Context, db *gorm.DB) {
	users, err := users.GetAllUsers(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": users, "message": "Users fetched successfully"})
}
