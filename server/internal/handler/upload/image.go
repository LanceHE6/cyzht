package upload

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
	"server/pkg/response"
	"server/pkg/rpc/file_server/api/v1/file_server"
	"strconv"
)

var imageExt = map[string]bool{
	".png":  true,
	".jpg":  true,
	".jpeg": true,
	".gif":  true,
	".svg":  true,
}

func (u uploadHandler) UploadImage(ctx *gin.Context) {
	file, header, err := ctx.Request.FormFile("file")
	if err == nil {
		filename := header.Filename
		// 获取文件后缀
		extString := path.Ext(filename)

		if !imageExt[extString] {
			ctx.JSON(http.StatusInternalServerError, response.FailedResponse(10, "不支持的文件格式"))
			return
		}

		var data = make([]byte, header.Size)
		_, _ = file.Read(data)

		rep, err := u.FileRpcServer.UploadImage(context.Background(), &file_server.UploadImageRequest{
			FileName:    filename,
			FileContent: data,
		})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, response.FailedResponse(http.StatusInternalServerError, err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, response.SuccessResponse(rep))
	} else {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(-1, "获取图片文件失败", err))
		return
	}
}

func (u uploadHandler) GetImageUrl(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.FailedResponse(100, err.Error()))
		return
	}

	rep, err := u.FileRpcServer.GetFileInfo(context.Background(), &file_server.GetFileInfoRequest{
		FileId: id,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(-1, "获取图片url失败", err))
		return
	}
	ctx.JSON(http.StatusOK, response.SuccessResponse(map[string]any{
		"id":  id,
		"url": u.C.Server.FileServer.StaticURL + rep.FileUrl,
	}))

}
