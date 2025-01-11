package version

import (
	"github.com/gin-gonic/gin"
	"server/pkg/response"
)

// GetVersion
//
//	@Description: 获取版本号
//	@receiver v versionHandler
//	@param ctx *gin.Context
func (v versionHandler) GetVersion(ctx *gin.Context) {
	const version = "V0.0.0.241010_A"
	ctx.JSON(200, response.NewResponse(0, "Hello cyzht!", map[string]any{
		"version": version,
	}))
}
