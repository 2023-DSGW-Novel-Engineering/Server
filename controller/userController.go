package controller

import (
	"log"
	"net/http"

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
	TargetUserName string `json:"target_user_name"`
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

	userModel := new(models.User)
	res := initializers.DB.Table("users").Where("name = ?", body.TargetUserName).Find(userModel)
	if res.Error != nil {
		log.Println(res.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "No Select userModel",
		})

		return
	}

	friendModel := &models.Friend{Me: int(user.ID), Target: int(userModel.ID)}
	res = initializers.DB.Table("friends").Create(friendModel)
	if res.Error != nil {
		log.Println(res.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "No create firend data",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{})

}

func GetFriendList(c *gin.Context) {
	temp, _ := c.Get("user-from-middleware")
	user := temp.(models.User)

	friendModel := []models.Friend{}
	res := initializers.DB.Where("me = ?", user.ID).Find(&friendModel)
	if res.Error != nil {
		log.Println(res.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "No Find friendModel",
		})

		return
	}

	response := []string{}
	for _, v := range friendModel {
		userModel := new(models.User)
		res := initializers.DB.Table("users").Where("id = ?", v.Target).Find(userModel)
		if res.Error != nil {
			log.Println(res.Error)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "user table no select",
			})

			return
		}

		response = append(response, userModel.Name)
	}

	c.JSON(http.StatusOK, gin.H{
		"users": response,
	})
}
