package main

import (
	"incrediblefour/config"
	cntD "incrediblefour/features/content/data"
	cntH "incrediblefour/features/content/handler"
	cntS "incrediblefour/features/content/services"
	usrD "incrediblefour/features/user/data"
	usrH "incrediblefour/features/user/handler"
	usrS "incrediblefour/features/user/services"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	db := config.InitDB(*cfg)
	config.Migrate(db)

	userData := usrD.New(db)
	userSrv := usrS.New(userData)
	userHdl := usrH.New(userSrv)

	contentData := cntD.New(db)
	contentSrv := cntS.New(contentData)
	contentHdl := cntH.New(contentSrv)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, error=${error}\n",
	}))

	e.POST("/register", userHdl.Register())
	e.POST("/login", userHdl.Login())

	e.GET("/contents/:id", contentHdl.MyContent())
	e.GET("/contents", contentHdl.ContentList())

	auth := e.Group("")
	auth.Use(middleware.JWT([]byte(config.JWTKey)))
	
	auth.GET("/users", userHdl.Profile())
	auth.PUT("/users", userHdl.Update())
	auth.DELETE("/users", userHdl.Deactivate())
	
	auth.POST("/contents", contentHdl.Add())
	auth.PUT("/contents/:id", contentHdl.Update())
	auth.DELETE("/contents/:id", contentHdl.Delete())

	if err := e.Start(":8000"); err != nil {
		log.Println(err.Error())
	}
}