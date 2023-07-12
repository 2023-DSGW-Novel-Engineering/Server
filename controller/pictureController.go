package controller

import "github.com/gin-gonic/gin"

const USER_PICTURES_PATH string = "pictures/users/"

func GetPicture(c *gin.Context) {
	name := c.Param("name")

	c.File(USER_PICTURES_PATH + name + ".png")
}
