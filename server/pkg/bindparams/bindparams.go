package bindparams

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/pkg/response"
)

// BindParams 泛型函数 绑定参数,出现绑定错误返回nil
func BindParams[T any](c *gin.Context) *T {
	var data T
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, response.FailedResponse(100, err.Error()))
		return nil
	}
	return &data
}
