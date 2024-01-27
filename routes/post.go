package routes

import (
	"net/http"
	"strconv"

	"example.com/blog-app-backend-go/db"
	"example.com/blog-app-backend-go/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetPosts(c *gin.Context) {

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

	var posts []models.Post

	result := db.DB.Offset(startIndex).Limit(endIndex).Find(&posts)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": posts,
	})

}

func GetPostByID(c *gin.Context) {

	id := c.Param("id")

	var post models.Post

	result := db.DB.First(&post, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, post)

}

func CreatePost(c *gin.Context) {

	userId, _ := uuid.Parse(c.GetString("userId"))

	type CreatePostInput struct {
		Title       string `form:"title" binding:"required"`
		Description string `form:"description" binding:"required"`
		Content     string `form:"content" binding:"required"`
	}

	var input CreatePostInput

	err := c.ShouldBind(&input)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	post := models.Post{
		Title:       input.Title,
		Description: input.Description,
		Content:     input.Content,
		AuthorID:    userId,
	}

	result := db.DB.Create(&post)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, post)

}

func UpdatePost(c *gin.Context) {

	type UpdatePostInput struct {
		Title       string `form:"title"`
		Description string `form:"description"`
		Content     string `form:"content"`
	}

	var input UpdatePostInput

	err := c.ShouldBind(&input)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	var post models.Post

	id := c.Param("id")

	result := db.DB.First(&post, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	userId, _ := uuid.Parse(c.GetString("userId"))

	if post.AuthorID != userId {
		c.Status(http.StatusForbidden)
		return
	}

	if input.Title != "" {
		post.Title = input.Title
	}

	if input.Description != "" {
		post.Description = input.Description
	}

	if input.Content != "" {
		post.Content = input.Content
	}

	result = db.DB.Save(&post)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, post)

}

func DeletePost(c *gin.Context) {

	id := c.Param("id")

	var post models.Post

	result := db.DB.First(&post, id)

	userId, _ := uuid.Parse(c.GetString("userId"))

	if post.AuthorID != userId {
		c.Status(http.StatusForbidden)
		return
	}

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	result = db.DB.Delete(&post)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, post)

}

func GetCommentsByPostID(c *gin.Context) {

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

	id := c.Param("id")

	var comments []models.Comment

	result := db.DB.Where("post_id = ?", id).Offset(startIndex).Limit(endIndex).Find(&comments)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": comments,
	})

}

func GetAuthorByPostID(c *gin.Context) {

	id := c.Param("id")

	var post models.Post

	result := db.DB.First(&post, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	var author models.User

	result = db.DB.First(&author, post.AuthorID)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, author)
}

func BookmarkPost(c *gin.Context) {

	userId := c.GetUint("userId")

	var user models.User

	result := db.DB.First(&user, userId)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	id := c.Param("id")

	var post models.Post

	result = db.DB.First(&post, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	err := db.DB.Model(&user).Association("Bookmarks").Append(&post)

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, post)

}

func UnbookmarkPost(c *gin.Context) {

	userId := c.GetUint("userId")

	var user models.User

	result := db.DB.First(&user, userId)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	id := c.Param("id")

	var post models.Post

	result = db.DB.First(&post, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	err := db.DB.Model(&user).Association("Bookmarks").Delete(&post)

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, post)

}
