package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
	"server/pkg/response"
	"server/pkg/rpc/file_server/api/v1/file_server"
)

var imageExt = map[string]bool{
	".png":  true,
	".jpg":  true,
	".jpeg": true,
	".gif":  true,
	".svg":  true,
}

func (s userService) UpdateAvatar() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		file, header, err := ctx.Request.FormFile("file")
		if err == nil {
			filename := header.Filename
			userInfo, _ := GetUserInfoByContext(ctx)
			// 判断文件类型是否为图片
			// 获取文件后缀
			extString := path.Ext(filename)
			if !imageExt[extString] {
				ctx.JSON(http.StatusBadRequest, response.FailedResponse(10, "文件类型不支持"))
				return
			}
			// 转换为base64
			var data = make([]byte, header.Size)
			_, _ = file.Read(data)
			rep, err := s.FileRpcServer.UploadAvatar(context.Background(), &file_server.UploadAvatarRequest{
				Id:          userInfo.ID,
				FileContent: data,
				FileName:    filename,
				FileType:    extString,
			})
			if err != nil || rep.FileUrl == "" {
				ctx.JSON(http.StatusInternalServerError, response.FailedResponse(-1, "上传文件失败 "+err.Error()))
				return
			}
			//base64Str := base64.StdEncoding.EncodeToString(data)
			//
			//userAvatar := mongodb.UserAvatarModel{
			//	CreatedAt: time.Now(),
			//	UpdatedAt: time.Now(),
			//	UID:       userInfo.ID,
			//	FileName:  filename,
			//	FileSize:  header.Size,
			//	Base64:    base64Str,
			//}
			err = s.UserRepo.UpdateAvatar(userInfo.ID, rep.FileUrl)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(-2, "更新头像失败", err))
				return
			}
			ctx.JSON(http.StatusOK, response.SuccessResponse("更新头像成功"))
		}
	}
}
