package routes

import (
	"net/http"
	"strconv"

	"example.com/blog-app-backend-go/db"
	"example.com/blog-app-backend-go/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetComments(c *gin.Context) {

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

	var comments []models.Comment

	result := db.DB.Offset(startIndex).Limit(endIndex).Find(&comments)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":       comments,
		"data_count": len(comments),
	})
}

func GetCommentByID(c *gin.Context) {

	var comment models.Comment

	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	result := db.DB.First(&comment, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, comment)

}

func CreateComment(c *gin.Context) {

	userId, err := uuid.Parse(c.GetString("userId"))

	if err != nil {
		c.Status(http.StatusUnauthorized)
	}

	type CreateCommentInput struct {
		Content string    `form:"content" binding:"required"`
		PostID  uuid.UUID `form:"post_id" binding:"required"`
	}

	var input CreateCommentInput

	err = c.ShouldBind(&input)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	comment := models.Comment{Content: input.Content, UserID: userId, PostID: input.PostID}

	result := db.DB.Create(&comment)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, comment)

}

func UpdateComment(c *gin.Context) {

	type UpdateCommentInput struct {
		Content string `form:"content" binding:"required"`
	}

	var input UpdateCommentInput

	err := c.ShouldBind(&input)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	var id uuid.UUID

	id, err = uuid.Parse(c.Param("id"))

	if err != nil {
		c.Status(http.StatusUnauthorized)
		return
	}

	var comment models.Comment

	result := db.DB.First(&comment, id)

	if result.Error != nil {

		c.Status(http.StatusInternalServerError)
		return
	}

	if input.Content != "" {
		comment.Content = input.Content
	}

	result = db.DB.Save(&comment)

	if result.Error != nil {

		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, comment)

}

func DeleteComment(c *gin.Context) {

	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.Status(http.StatusUnauthorized)
		return
	}

	var comment models.Comment

	result := db.DB.First(&comment, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	result = db.DB.Delete(&comment)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, comment)

}

func GetAuthorByCommentID(c *gin.Context) {

	var comment models.Comment

	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	result := db.DB.First(&comment, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	var user models.User

	result = db.DB.First(&user, comment.UserID)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, user)
}
