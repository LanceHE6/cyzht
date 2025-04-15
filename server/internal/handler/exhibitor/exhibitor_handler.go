package exhibitor

import (
	"github.com/gin-gonic/gin"
	"server/internal/repo/exhibitor"
	"server/internal/repo/exhibitoruser"
)

type HandlerInterface interface {
	AddExhibitor(ctx *gin.Context)
	DeleteExhibitor(ctx *gin.Context)
	ListExhibitor(ctx *gin.Context)
	JoinExhibitor(ctx *gin.Context)
	WithdrawExhibitor(ctx *gin.Context)
}

type exhibitorHandler struct {
	ExhibitorRepo     exhibitor.RepoInterface
	ExhibitorUserRepo exhibitoruser.RepoInterface
}

func NewExhibitorHandler(exhibitorRepo exhibitor.RepoInterface,
	exhibitorUserRepo exhibitoruser.RepoInterface) HandlerInterface {
	return &exhibitorHandler{
		ExhibitorRepo:     exhibitorRepo,
		ExhibitorUserRepo: exhibitorUserRepo,
	}
}
