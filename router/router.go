package router

import (
	"os"

	"github.com/gin-gonic/gin"
	"test.com/req/controllers"
)

func Routers() {
	port := os.Getenv("PORT")
	// """""""controllers"""""""
	login := controllers.Login
	signin := controllers.Signin
	cors := controllers.Cors
	logout := controllers.Logout
	// """""""""routers"""""""""
	router := gin.Default()
	// """"""""""static files""""""""""
	router.StaticFS("/static",gin.Dir("./static/",true))

	router.Use(cors)
	router.POST("/login",login)
	router.POST("/signin",signin)
	router.GET("/logout",logout)

	router.Run(":" + port)
}
