package routes

import (
	"Golang-User-Auth/controllers"

	"github.com/gin-gonic/gin"
)

func Handlers() *gin.Engine {

	r := gin.Default()
	//r.Use(cors.Default())
	r.LoadHTMLGlob("templates/static/*")
	r.Static("/assets", "./templates/assets")
	r.Static("/admin/assets", "./templates/assets")
	r.StaticFile("/favicon.ico", "./templates/assets/favicon.ico")
	r.GET("/", controllers.HomePage)
	r.POST("/", controllers.LoginForm)

	r.GET("/dashboard", controllers.UserDashBoard) //User dashboard

	r.GET("/register", controllers.RegisterHomePage)
	r.POST("/register", controllers.RegisterForm)

	r.GET("/admin", controllers.AdminDashBoard) //Admin dashboard
	r.DELETE("/admin/:id", controllers.DeleteUser)

	r.PUT("/admin/:id", controllers.UpdateUser)
	r.GET("/admin/list", controllers.ListUser)
	r.GET("/logout", controllers.Logout) //Common Logout
	//r.GET("/t/:id", controllers.UpdateUser)
	r.NoRoute(func(c *gin.Context) {
		c.HTML(404, "404.html", gin.H{"message": "Page not found"})
	})

	return r
}
