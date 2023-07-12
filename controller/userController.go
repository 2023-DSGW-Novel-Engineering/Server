package controller

import (
	"net/http"

	"github.com/2023-DSGW-Novel-Engineering/cation-backend/initializers"
	"github.com/2023-DSGW-Novel-Engineering/cation-backend/models"
	"github.com/gin-gonic/gin"
)

func GetUserInfo(c *gin.Context) {
	var body struct {
		Name string `json:"name"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	if body.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Empty Each body values",
		})

		return
	}

	// DB select

	// db.Where("name = ?", "John").Find(&users)
	// select * from users where name = body.Name
	user := new(models.User)
	if res := initializers.DB.Where("name = ?", body.Name).Find(user); res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User not found",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"name":            user.Name,
		"user_id":         user.UserID,
		"native_language": user.NativeLanguage,
	})
}

func AddFriend(c *gin.Context) {

}
