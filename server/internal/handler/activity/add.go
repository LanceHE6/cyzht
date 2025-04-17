package activity

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
	"server/internal/model"
	"server/pkg/bindparams"
	"server/pkg/jwt"
	"server/pkg/response"
	"server/pkg/timeconv"
)

var imageExt = map[string]bool{
	".png":  true,
	".jpg":  true,
	".jpeg": true,
	".gif":  true,
	".svg":  true,
}

// AddActivity 新增活动
func (a *activityHandler) AddActivity(ctx *gin.Context) {
	type addActivityRequest struct {
		Name      string `json:"name" binding:"required" form:"name"`
		Introduce string `json:"introduce" form:"introduce"`
		StartAt   string `json:"start_at" binding:"required" form:"start_at"`
		EndAt     string `json:"end_at" binding:"required" form:"end_at"`
		Location  string `json:"location" binding:"required" form:"location"`
	}
	data := bindparams.BindPostParams[addActivityRequest](ctx)
	if data == nil {
		return
	}
	startAt, err := timeconv.ParesStrToTime(data.StartAt)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.FailedResponse(100, "start_at is in the wrong time format"))
		return
	}
	endAt, err := timeconv.ParesStrToTime(data.EndAt)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.FailedResponse(100, "end_at is in the wrong time format"))
		return
	}
	// 插入新展会记录
	claims, _ := jwt.GetClaimsByContext(ctx)
	activity := model.ActivityModel{
		Name:      data.Name,
		Introduce: data.Introduce,
		StartAt:   *startAt,
		EndAt:     *endAt,
		Location:  data.Location,
		CreatorID: claims.ID,
	}
	aid, err := a.ActivityRepo.Insert(&activity)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(-1, "failed in insert activity", err))
		return
	}
	// 创建者自动加入活动
	if err := a.ActivityUserRepo.Insert(claims.ID, aid, 3); err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(-1, "failed in insert activity user", err))
		return
	}

	file, header, err := ctx.Request.FormFile("avatar")
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

		err := a.ActivityRepo.UploadIcon(aid, filename, data)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, response.FailedResponse(http.StatusInternalServerError, err.Error()))
			return
		}
		//ctx.JSON(http.StatusOK, response.SuccessResponse("更新头像成功"))
	}
	ctx.JSON(http.StatusOK, response.SuccessResponse(nil))
	return
}
