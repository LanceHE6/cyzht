package repo

import (
	"errors"
	"file_server/internal/data/models"
	"gorm.io/gorm"
)

type ActivityAvatarRepoInterface interface {
	InsertOrUpdate(userAvatar *models.ActivityAvatarModel) (avatar *models.ActivityAvatarModel, err error)
	Update(userAvatar *models.ActivityAvatarModel) (avatar *models.ActivityAvatarModel, err error)
	FindByID(id int64) (avatar *models.ActivityAvatarModel, err error)
}

func NewActivityAvatarRepo(db *gorm.DB) ActivityAvatarRepoInterface {
	return &activityAvatarRepo{
		DB: db,
	}
}

type activityAvatarRepo struct {
	DB *gorm.DB
}

func (u activityAvatarRepo) InsertOrUpdate(userAvatar *models.ActivityAvatarModel) (avatar *models.ActivityAvatarModel, err error) {
	avatar, err = u.FindByID(userAvatar.ID)
	if avatar != nil {
		// update
		return u.Update(userAvatar)
	}
	err = u.DB.Model(&models.ActivityAvatarModel{}).Create(userAvatar).Error
	return userAvatar, err
}

func (u activityAvatarRepo) Update(userAvatar *models.ActivityAvatarModel) (avatar *models.ActivityAvatarModel, err error) {
	err = u.DB.Model(&models.ActivityAvatarModel{}).Where("id = ?", userAvatar.ID).Updates(userAvatar).Error
	return userAvatar, err
}

func (u activityAvatarRepo) FindByID(id int64) (avatar *models.ActivityAvatarModel, err error) {
	err = u.DB.Model(&models.ActivityAvatarModel{}).Where("id = ?", id).First(&avatar).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return avatar, err
}
