package controller

import (
	"net/http"
	"strconv"

	"github.com/2023-DSGW-Novel-Engineering/cation-backend/initializers"
	"github.com/2023-DSGW-Novel-Engineering/cation-backend/models"
	"github.com/gin-gonic/gin"
)

func GetUserInfo(c *gin.Context) {
	// Path/value 파싱
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user ID",
		})
		return
	}

	// DB select
	user := new(models.User)
	if res := initializers.DB.First(user, id); res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":              user.ID,
		"name":            user.Name,
		"native_language": user.NativeLanguage,
		"image_path":      user.ImagePath,
	})
}
