package exhibitor

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/internal/model"
	"server/pkg/bindparams"
	"server/pkg/jwt"
	"server/pkg/response"
)

func (e *exhibitorHandler) AddExhibitor(ctx *gin.Context) {
	type addExhibitorRequest struct {
		Name       string `json:"name" binding:"required"`
		Introduce  string `json:"introduce"`
		ActivityID int64  `json:"aid,string" binding:"required"`
	}
	data := bindparams.BindPostParams[addExhibitorRequest](ctx)
	if data == nil {
		return
	}
	claims, _ := jwt.GetClaimsByContext(ctx)
	err := e.ExhibitorRepo.Insert(&model.ExhibitorModel{
		Name:       data.Name,
		Introduce:  data.Introduce,
		ActivityID: data.ActivityID,
		CreatorID:  claims.ID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.FailedResponse(-1, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.SuccessResponse(nil))
}
