package main

import (
	"incrediblefour/config"
	cmtD "incrediblefour/features/comment/data"
	cmtH "incrediblefour/features/comment/handler"
	cmtS "incrediblefour/features/comment/services"
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

	commentData := cmtD.New(db)
	commentSrv := cmtS.New(commentData)
	commentHdl := cmtH.New(commentSrv)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, error=${error}\n",
	}))

	e.POST("/register", userHdl.Register())
	e.POST("/login", userHdl.Login())

	e.GET("/contents/:id", contentHdl.ContentDetail())
	e.GET("/contents", contentHdl.ContentList())

	user := e.Group("/users")

	content := e.Group("")
	content.Use(middleware.JWT([]byte(config.JWTKey)))

	comment := e.Group("")
	comment.Use(middleware.JWT([]byte(config.JWTKey)))

	user.GET("/:username", contentHdl.GetProfile())
	user.GET("", userHdl.Profile(), middleware.JWT([]byte(config.JWTKey)))
	user.PUT("", userHdl.Update(), middleware.JWT([]byte(config.JWTKey)))
	user.DELETE("", userHdl.Deactivate(), middleware.JWT([]byte(config.JWTKey)))

	content.POST("/contents", contentHdl.Add())
	content.PUT("/contents/:id", contentHdl.Update())
	content.DELETE("/contents/:id", contentHdl.Delete())

	comment.POST("/comments/:id", commentHdl.Add())
	comment.DELETE("/comments/:id", commentHdl.Delete())

	if err := e.Start(":8000"); err != nil {
		log.Println(err.Error())
	}
}
