package sample

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Test01Request struct {
	Name string `json:"name"`
}

// @Summary Test01
// @Description Test01
// @Tags middleman
// @Produce  json
// @Param body body Test01Request true "content"
// @Success 200 {string} 12345
// @Router /test01 [post]
func (*SampleService) Test01(c *gin.Context) {
	c.String(http.StatusOK, "1111", "")
	return
}
