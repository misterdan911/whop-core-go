package routes

import (
	"whop-core-go/controllers"
	"whop-core-go/db"
	"whop-core-go/middleware"
	"whop-core-go/models"

	"github.com/gin-gonic/gin"
)

func Setup(route *gin.Engine) {

	var auth controllers.Auth

	/* ---------------------------  Public routes  --------------------------- */
	public := route.Group("/api/v1")
	public.POST("/login", auth.Login)

	/* ---------------------------  Private routes  --------------------------- */
	private := route.Group("/api/v1")
	private.Use(middleware.JwtVerifier)
	private.GET("/user/:id", func(ctx *gin.Context) {

		userId := ctx.Param("id")

		var user models.Users
		db.Whop.First(&user, userId)

		ctx.JSON(200, gin.H{
			"message": user.Username,
		})
	})

}
