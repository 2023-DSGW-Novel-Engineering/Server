package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/2023-DSGW-Novel-Engineering/cation-backend/initializers"
	"github.com/2023-DSGW-Novel-Engineering/cation-backend/models"
	"github.com/gin-gonic/gin"
)

func GetUserInfo(c *gin.Context) {
	temp, _ := c.Get("user-from-middleware")
	user := temp.(models.User)

	c.JSON(http.StatusOK, gin.H{
		"name":            user.Name,
		"user_id":         user.UserID,
		"native_language": user.NativeLanguage,
	})

}

type AddFriendBody struct {
	TargetUserID string `json:"target_user_id"`
}

func AddFriend(c *gin.Context) {
	body := new(AddFriendBody)
	if err := c.BindJSON(body); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	temp, _ := c.Get("user-from-middleware")
	user := temp.(models.User)

	i, _ := strconv.Atoi(body.TargetUserID)

	friendModel := &models.Friend{Me: int(user.ID), Target: i}
	res := initializers.DB.Create(friendModel)
	if res.Error != nil {
		log.Println(res.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "No create firend data",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{})

}
