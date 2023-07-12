package controller

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/2023-DSGW-Novel-Engineering/cation-backend/initializers"
	"github.com/2023-DSGW-Novel-Engineering/cation-backend/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Register(c *gin.Context) {
	var body struct {
		Name           string `json:"name"`
		UserID         string `json:"user_id"`
		Password       string `json:"password"`
		NativeLanguage string `json:"native_language"`
	}

	// Body값 bind
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	if body.Name == "" || body.UserID == "" || body.Password == "" || body.NativeLanguage == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Empty Each body values",
		})
	}

	// 레코드 생성
	user := models.User{Name: body.Name, UserId: body.UserID, Password: body.Password, NativeLanguage: body.NativeLanguage}
	res := initializers.DB.Create(&user)
	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user recode",
		})

		return
	}

	// 성공
	c.JSON(http.StatusOK, gin.H{})
}

func Login(c *gin.Context) {
	var body struct {
		UserID   string `json:"user_id"`
		Password string `json:"password"`
	}

	// body 값 bind
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	// Empty 확인
	if body.UserID == "" || body.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Empty Each body values",
		})

		return
	}

	// DB Select
	var user models.User
	res := initializers.DB.First(&user, "user_id = ?", body.UserID)
	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id or passsowrd",
		})

		return
	}

	// JWT 토큰 생성
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30 /*30일*/).Unix(),
	})

	// 최종적으로 생성
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})

		return
	}

	// Token을 되돌려 주다.
	c.SetCookie("Authorization", tokenString, 3600*24*30 /*30일*/, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})

}

func Logout(c *gin.Context) {
	c.SetCookie("Authorization", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "Success logout",
	})

}

func Vaildate(c *gin.Context) {
	user, _ := c.Get("user")

	fmt.Println(user)
	c.JSON(http.StatusOK, gin.H{})

}
