package server

import (
	"github.com/gin-gonic/gin"

	"github.com/ryo-ymd/go-test/controller"
)

func Init() {
	r := router()
	r.Run()
}

func router() *gin.Engine {
	r := gin.Default()

	a := r.Group("/articles")
	{
		ctrl := article.Controller{}
		a.GET("", ctrl.Index)
	}

	return r
}
