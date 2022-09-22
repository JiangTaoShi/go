package sample

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type Test01Request struct {
	Name string `json:"name"binding:"required"`
}

// @Summary Test01
// @Description Test01
// @Tags sample
// @Produce  json
// @Param body body sample.Test01Request true "content"
// @Success 200 {string} 12345
// @Router /test01 [post]
func (*SampleService) Test01(c *gin.Context) {
	var request Test01Request
	validate := validator.New()
	validate.Struct(request)

	//err := c.ShouldBindJSON(&request)
	//if err != nil {
	//	c.String(http.StatusBadRequest, "", err.Error())
	//	return
	//}
	c.String(http.StatusOK, "", "1111")
	return
}
