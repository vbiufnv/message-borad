package api

import (
	"github.com/gin-gonic/gin"
)

func SetRouter() {
	r := gin.Default()

	r.POST("/login", Login)
	r.POST("/register", Register)

	r.Use(AuthMiddleware())

	r.POST("/user/quit", QuitLogin)
	r.PUT("/user/update", Update)

	r.POST("/post", CreatePost)
	r.PUT("/post", UpdatePost)
	r.DELETE("/post", DeletePost)
	r.GET("/post", ListPosts)

	r.GET("/post/:id", IdToComments)

	//点赞
	r.POST("/post/:post_id/like", LikePost)

	r.POST("/comment", CreateComment)
	r.PUT("/comment", UpdateComment)
	r.DELETE("/comment", DeleteComment)

	r.Run(":8080")
}
