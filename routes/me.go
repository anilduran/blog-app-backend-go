package routes

import (
	"net/http"
	"strconv"

	"example.com/blog-app-backend-go/db"
	"example.com/blog-app-backend-go/models"
	"example.com/blog-app-backend-go/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetMyCredentials(c *gin.Context) {

	userId, _ := uuid.Parse(c.GetString("userId"))

	var user models.User

	result := db.DB.First(&user, userId)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, user)

}

func UpdateMyCredentials(c *gin.Context) {

	userId, _ := uuid.Parse(c.GetString("userId"))

	type UpdateMyCredentialsInput struct {
		Username string `form:"username"`
		Email    string `form:"email"`
		Password string `form:"password"`
	}

	var input UpdateMyCredentialsInput

	err := c.ShouldBind(&input)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	var user models.User

	result := db.DB.First(&user, userId)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	if input.Username != "" {
		user.Username = input.Username
	}

	if input.Email != "" {
		user.Email = input.Email
	}

	if input.Password != "" {
		hashedPassword := utils.HashPassword(input.Password)
		user.Password = hashedPassword
	}

	result = db.DB.Save(&user)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, user)

}

func GetMyPosts(c *gin.Context) {

	userId, _ := uuid.Parse(c.GetString("userId"))

	var posts []models.Post

	result := db.DB.Where("author_id = ?", userId).Find(&posts)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": posts,
	})

}

func GetMyComments(c *gin.Context) {

	page, err := strconv.ParseInt(c.Query("page"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Enter a valid page number",
		})
		return
	}

	limit, err := strconv.ParseInt(c.Query("limit"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Enter a valid limit number",
		})
		return
	}

	startIndex := int((page - 1) * limit)
	endIndex := int(page * limit)

	userId, _ := uuid.Parse(c.GetString("userId"))

	var comments []models.Comment

	result := db.DB.Where("user_id = ?", userId).Offset(startIndex).Limit(endIndex).Find(&comments)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": comments,
	})

}

func GetMyReadingLists(c *gin.Context) {

	page, err := strconv.ParseInt(c.Query("page"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Enter a valid page number",
		})
		return
	}

	limit, err := strconv.ParseInt(c.Query("limit"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Enter a valid limit number",
		})
		return
	}

	startIndex := int((page - 1) * limit)
	endIndex := int(page * limit)

	userId, _ := uuid.Parse(c.GetString("userId"))

	// var user models.User

	var readingLists []models.ReadingList

	result := db.DB.Where("user_id = ?", userId).Offset(startIndex).Limit(endIndex).Find(&readingLists)

	// result := db.DB.Preload("ReadingLists").First(&user, userId)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": readingLists,
	})

}

func GetMyBookmarks(c *gin.Context) {

	userId, _ := uuid.Parse(c.GetString("userId"))

	var user models.User

	result := db.DB.Preload("Bookmarks").First(&user, userId)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user.Bookmarks,
	})

}

func GetPresignedUrl(c *gin.Context) {

	userId, _ := uuid.Parse(c.GetString("userId"))

	url, key, err := utils.PresignerInstance.PutObject(userId.String())

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"url": url,
		"key": key,
	})

}
