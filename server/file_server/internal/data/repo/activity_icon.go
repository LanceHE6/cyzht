package repo

import (
	"errors"
	"file_server/internal/data/models"
	"gorm.io/gorm"
)

type ActivityIconRepoInterface interface {
	InsertOrUpdate(aa *models.ActivityIconModel) (avatar *models.ActivityIconModel, err error)
	Update(aa *models.ActivityIconModel) (avatar *models.ActivityIconModel, err error)
	FindByID(id int64) (avatar *models.ActivityIconModel, err error)
	FindByHash(hash string) (avatar *models.ActivityIconModel, err error)
}

func NewActivityIconRepo(db *gorm.DB) ActivityIconRepoInterface {
	return &activityIconRepo{
		DB: db,
	}
}

type activityIconRepo struct {
	DB *gorm.DB
}

func (r *activityIconRepo) InsertOrUpdate(icon *models.ActivityIconModel) (avatar *models.ActivityIconModel, err error) {
	avatar, err = r.FindByID(icon.ID)
	if avatar != nil {
		// update
		return r.Update(icon)
	}
	err = r.DB.Model(&models.ActivityIconModel{}).Create(icon).Error
	return icon, err
}

func (r *activityIconRepo) Update(aa *models.ActivityIconModel) (avatar *models.ActivityIconModel, err error) {
	err = r.DB.Model(&models.ActivityIconModel{}).Where("id = ?", aa.ID).Updates(aa).Error
	return aa, err
}

func (r *activityIconRepo) FindByID(id int64) (avatar *models.ActivityIconModel, err error) {
	err = r.DB.Model(&models.ActivityIconModel{}).Where("id = ?", id).First(&avatar).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return avatar, err
}

func (r *activityIconRepo) FindByHash(hash string) (*models.ActivityIconModel, error) {
	var icon models.ActivityIconModel
	err := r.DB.Where("hash = ?", hash).First(&icon).Error
	if err != nil {
		return nil, err
	}
	return &icon, nil
}
