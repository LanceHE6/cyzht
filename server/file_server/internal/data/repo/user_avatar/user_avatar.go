package user_avatar

import (
	"errors"
	"file_server/internal/data/models"
	"gorm.io/gorm"
)

type UserAvatarRepo struct {
	DB *gorm.DB
}

func (u UserAvatarRepo) InsertOrUpdate(userAvatar *models.UserAvatarModel) (avatar *models.UserAvatarModel, err error) {
	avatar, err = u.FindByID(userAvatar.ID)
	if avatar != nil {
		// update
		return u.Update(userAvatar)
	}
	err = u.DB.Model(&models.UserAvatarModel{}).Create(userAvatar).Error
	return userAvatar, err
}

func (u UserAvatarRepo) Update(userAvatar *models.UserAvatarModel) (avatar *models.UserAvatarModel, err error) {
	err = u.DB.Model(&models.UserAvatarModel{}).Where("id = ?", userAvatar.ID).Updates(userAvatar).Error
	return userAvatar, err
}

func (u UserAvatarRepo) FindByID(id int64) (avatar *models.UserAvatarModel, err error) {
	err = u.DB.Model(&models.UserAvatarModel{}).Where("id = ?", id).First(&avatar).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return avatar, err
}
