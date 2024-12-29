package repo

import (
	"file_server/internal/data/models"
	"file_server/internal/data/repo/user_avatar"
	"gorm.io/gorm"
)

type UserAvatarRepoInterface interface {
	InsertOrUpdate(userAvatar *models.UserAvatarModel) (avatar *models.UserAvatarModel, err error)
	Update(userAvatar *models.UserAvatarModel) (avatar *models.UserAvatarModel, err error)
	FindByID(id int64) (avatar *models.UserAvatarModel, err error)
}

func NewUserAvatarRepo(db *gorm.DB) UserAvatarRepoInterface {
	return &user_avatar.UserAvatarRepo{
		DB: db,
	}
}
