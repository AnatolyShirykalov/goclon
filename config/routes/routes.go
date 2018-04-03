package routes

import (
	"../../app/controllers"
	"../admin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

var r *gin.Engine

func Router() *gin.Engine {
	if r != nil {
		return r
	}
	acc := gin.Accounts{
		"anatoly": "1234",
	}
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowCredentials = true
	config.AllowMethods = []string{"GET", "PUT"}

	r = gin.Default()
	r.Use(cors.New(config))
	authorized := r.Group("/", gin.BasicAuth(acc))
	mux := http.NewServeMux()
	admin.Admin.MountTo("/admin", mux)
	authorized.Any("/admin/*w", gin.WrapH(mux))
	authorized.GET("/", controllers.HomeIndex)
	comments := r.Group("/api/comments")
	comments.PUT("/", controllers.CommentsUpdate)
	comments.GET("/", controllers.CommentsIndex)
	return r
}
