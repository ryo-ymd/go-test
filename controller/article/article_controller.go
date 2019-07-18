package article

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/ryo-ymd/go-test/service/article"
)

type Controller struct {}

// ArticleControllerはArticleのアクションを定義する
func (pc Controller) Index(c *gin.Context) {
	var s article.Service

	a, err := s.GetAll()

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, a)
	}
}
