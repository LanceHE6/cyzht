package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
	"server/pkg/jwt"
	"server/pkg/response"
)

var imageExt = map[string]bool{
	".png":  true,
	".jpg":  true,
	".jpeg": true,
	".gif":  true,
	".svg":  true,
}

func (s userHandler) UpdateAvatar(ctx *gin.Context) {
	file, header, err := ctx.Request.FormFile("file")
	if err == nil {
		filename := header.Filename
		userInfo, _ := jwt.GetClaimsByContext(ctx)
		// 获取文件后缀
		extString := path.Ext(filename)

		if !imageExt[extString] {
			ctx.JSON(http.StatusInternalServerError, response.FailedResponse(http.StatusInternalServerError, "不支持的文件格式"))
			return
		}

		var data = make([]byte, header.Size)
		_, _ = file.Read(data)

		err := s.UserRepo.UploadAvatar(userInfo.ID, filename, data)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, response.FailedResponse(http.StatusInternalServerError, err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, response.SuccessResponse("更新头像成功"))
	}
}
