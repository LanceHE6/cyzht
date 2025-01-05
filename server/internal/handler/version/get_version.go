package version

import (
	"github.com/gin-gonic/gin"
	"server/pkg/response"
)

// GetVersion
//
//	@Description: 获取版本号
//	@receiver v versionHandler
//	@param context *gin.Context
func (v versionHandler) GetVersion() gin.HandlerFunc {
	return func(context *gin.Context) {
		const version = "V0.0.0.241010_A"
		context.JSON(200, response.NewResponse(0, "Hello NetChat!", map[string]any{
			"version": version,
		}))
	}
}
