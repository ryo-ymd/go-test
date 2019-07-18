package server

import (
	"github.com/gin-gonic/gin"
	"github.com/ryo-ymd/go-test/controller/tag"

	"github.com/ryo-ymd/go-test/controller/article"
)

// Serverはルーティングなどの設定を行う
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

	t := r.Group("/tags")
	{
		ctrl := tag.Controller{}
		t.GET("", ctrl.Index)
		t.GET("/:id", ctrl.Show)
		t.GET("/:id/ancestors", ctrl.Ancestors)
	}

	return r
}
