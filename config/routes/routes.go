package routes

import (
	"../admin"
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

	r = gin.Default()
	authorized := r.Group("/", gin.BasicAuth(acc))
	mux := http.NewServeMux()
	admin.Admin.MountTo("/admin", mux)
	authorized.Any("/admin/*w", gin.WrapH(mux))
	return r
}
