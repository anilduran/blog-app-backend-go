package routes

import (
	"example.com/blog-app-backend-go/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	auth := r.Group("/api/auth")
	auth.POST("/sign-in", SignIn)
	auth.POST("/sign-up", SignUp)

	me := r.Group("/api/me")
	me.Use(middlewares.Auth)
	me.GET("/", GetMyCredentials)
	me.PUT("/", UpdateMyCredentials)
	me.GET("/posts", GetMyPosts)
	me.GET("/comments", GetMyComments)
	me.GET("/reading-lists", GetMyReadingLists)
	me.GET("/bookmarks", GetMyBookmarks)
	me.GET("/presigned-url", GetPresignedUrl)

	category := r.Group("/api/categories")
	category.Use(middlewares.Auth)
	category.GET("/", GetCategories)
	category.GET("/:id", GetCategoryByID)
	category.POST("/", CreateCategory)
	category.PUT("/:id", UpdateCategory)
	category.DELETE("/:id", DeleteCategory)
	category.GET("/:id/posts", GetPostsByCategoryID)

	post := r.Group("/api/posts")
	post.Use(middlewares.Auth)
	post.GET("/", GetPosts)
	post.GET("/:id", GetPostByID)
	post.POST("/", CreatePost)
	post.PUT("/:id", UpdatePost)
	post.DELETE("/:id", DeletePost)
	post.GET("/:id/comments", GetCommentsByPostID)
	post.GET("/:id/author", GetAuthorByPostID)
	post.POST("/:id/bookmark", BookmarkPost)
	post.DELETE("/:id/bookmark", UnbookmarkPost)

	comment := r.Group("/api/comments")
	comment.Use(middlewares.Auth)
	comment.GET("/", GetComments)
	comment.GET("/:id", GetCommentByID)
	comment.POST("/", CreateComment)
	comment.PUT("/:id", UpdateComment)
	comment.DELETE("/:id", DeleteComment)
	comment.GET("/:id/author", GetAuthorByCommentID)

	user := r.Group("/api/users")
	user.Use(middlewares.Auth)
	user.GET("/", GetUsers)
	user.GET("/:id", GetUserByID)
	user.POST("/", CreateUser)
	user.PUT("/:id", UpdateUser)
	user.DELETE("/:id", DeleteUser)

	role := r.Group("/api/roles")
	role.Use(middlewares.Auth)
	role.GET("/", GetRoles)
	role.GET("/:id", GetRoleByID)
	role.POST("/", CreateRole)
	role.PUT("/:id", UpdateRole)
	role.DELETE("/:id", DeleteRole)

	readingList := r.Group("/api/reading-lists")
	readingList.Use(middlewares.Auth)
	readingList.GET("/", GetReadingLists)
	readingList.GET("/:id", GetReadingListByID)
	readingList.POST("/", CreateReadingList)
	readingList.PUT("/:id", UpdateReadingList)
	readingList.DELETE("/:id", DeleteReadingList)
	readingList.GET("/:id/posts", GetReadingListPosts)
	readingList.POST("/:id/posts/:postId", AddPostToReadingList)
	readingList.DELETE("/:id/posts/:postId", RemovePostFromReadingList)

}
