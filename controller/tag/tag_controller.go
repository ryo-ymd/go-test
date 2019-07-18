package tag

import (
	"fmt"
	"github.com/ryo-ymd/go-test/service/tag"
	"strconv"

	"github.com/gin-gonic/gin"

)

type Controller struct {}

// TagControllerはTagのアクションを定義する
func (pc Controller) Index(c *gin.Context) {
	var s tag.Service

	a, err := s.GetAll()

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, a)
	}
}

func (pc Controller) Show(c *gin.Context) {
	var s tag.Service

	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	a, err := s.FindById(uint(id))

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
		return
	}

	c.JSON(200, a)
}

func (pc Controller) Ancestors(c *gin.Context) {
	var s tag.Service

	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	a, err := s.GetTree(uint(id))

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
		return
	}
	c.JSON(200, a)
}
