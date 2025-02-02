package bindparams

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/pkg/response"
)

// BindPostParams 泛型函数 绑定post参数,出现绑定错误返回nil
func BindPostParams[T any](c *gin.Context) *T {
	var data T
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, response.FailedResponse(100, err.Error()))
		return nil
	}
	return &data
}

// BindQueryParams 泛型函数 绑定query参数,出现绑定错误返回空T
func BindQueryParams[T any](c *gin.Context) *T {
	var params T
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, response.FailedResponse(100, err.Error()))
		return &params
	}
	return &params
}
