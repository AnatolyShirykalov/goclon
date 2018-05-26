package routes

import (
	"../../app/controllers"
	"../admin"
	auth "../siteauth"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"net/http"
)

var r *gin.Engine

func Router() *gin.Engine {
	if r != nil {
		return r
	}
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000", "https://clonclient.shirykalov.ru", "http://localhost:5292", "http://192.168.1.103:3000"}
	config.AllowCredentials = true
	config.AllowMethods = []string{"GET", "PUT"}

	Auth := auth.Auth

	r = gin.Default()
	r.Use(cors.New(config))
	r.Any("/auth/*w", gin.WrapH(Auth.NewServeMux()))
	mux := http.NewServeMux()
	admin.Admin.MountTo("/admin", mux)
	r.Any("/admin/*w", gin.WrapH(mux))
	r.GET("/", controllers.HomeIndex)
	comments := r.Group("/api/comments")
	comments.PUT("/", controllers.CommentsUpdate)
	comments.GET("/", controllers.CommentsIndex)
	return r
}
