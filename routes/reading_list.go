package routes

import (
	"net/http"
	"strconv"

	"example.com/blog-app-backend-go/db"
	"example.com/blog-app-backend-go/models"
	"github.com/gin-gonic/gin"
)

func GetReadingLists(c *gin.Context) {

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

	var readingLists []models.ReadingList

	result := db.DB.Offset(startIndex).Limit(endIndex).Find(&readingLists)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": readingLists,
	})

}

func GetReadingListByID(c *gin.Context) {

	id := c.Param("id")

	var readingList models.ReadingList

	result := db.DB.First(&readingList, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, readingList)

}

func CreateReadingList(c *gin.Context) {

	userId := c.GetUint("userId")

	type CreateReadingListInput struct {
		Name        string `form:"name" binding:"required"`
		Description string `form:"description" binding:"required"`
	}

	var input CreateReadingListInput

	if err := c.ShouldBind(&input); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	readingList := models.ReadingList{
		Name:        input.Name,
		Description: input.Description,
		UserID:      userId,
	}

	result := db.DB.Create(&readingList)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, readingList)

}

func UpdateReadingList(c *gin.Context) {

	type UpdateReadingListInput struct {
		Name        string `form:"name"`
		Description string `form:"description"`
	}

	var input UpdateReadingListInput

	err := c.ShouldBind(&input)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	var readingList models.ReadingList

	id := c.Param("id")

	result := db.DB.First(&readingList, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	if input.Name != "" {
		readingList.Name = input.Name
	}

	if input.Description != "" {
		readingList.Description = input.Description
	}

	result = db.DB.Save(&readingList)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, readingList)

}

func DeleteReadingList(c *gin.Context) {

	id := c.Param("id")

	var readingList models.ReadingList

	result := db.DB.First(&readingList, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	result = db.DB.Delete(&readingList)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, readingList)

}

func GetReadingListPosts(c *gin.Context) {

	id := c.Param("id")

	var readingList models.ReadingList

	result := db.DB.Preload("Posts").First(&readingList, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": readingList.Posts,
	})

}

func AddPostToReadingList(c *gin.Context) {

	id := c.Param("id")

	var readingList models.ReadingList

	result := db.DB.First(&readingList, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	postId := c.Param("postId")

	var post models.Post

	result = db.DB.First(&post, postId)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	err := db.DB.Model(&readingList).Association("Posts").Append(&post)

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, post)

}

func RemovePostFromReadingList(c *gin.Context) {

	id := c.Param("id")

	var readingList models.ReadingList

	result := db.DB.First(&readingList, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	postId := c.Param("postId")

	var post models.Post

	result = db.DB.First(&post, postId)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	err := db.DB.Model(&readingList).Association("Posts").Delete(&post)

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, post)

}
