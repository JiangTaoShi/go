package sample

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.gin/pkg/zip"
	"net/http"
	"os"
	"path/filepath"
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
	//var request Test01Request
	//err := c.ShouldBindJSON(request)
	//if err != nil {
	//	c.String(http.StatusBadRequest, "", err.Error())
	//	return
	//}
	compressUrlList := []string{"./temp/02.jpg", "./temp/01.jpg"}
	var files = []*os.File{}
	for i := 0; i < len(compressUrlList); i++ {
		curl, _ := filepath.Abs(compressUrlList[i])
		cfile, err := os.Open(curl)
		if err != nil {
			fmt.Println(errors.New("图片压缩异常"))
			return
		}
		files = append(files, cfile)
		defer cfile.Close()
	}
	zipPath, _ := filepath.Abs(fmt.Sprintf("./temp/%s.zip", "test"))
	err := zip.CompressArr(files, zipPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.String(http.StatusOK, "", "1111")
	return
}
