package repo

import (
	"errors"
	"file_server/internal/data/models"
	"gorm.io/gorm"
)

type ImageRepoInterface interface {
	InsertOrUpdate(i *models.ImageModel) (image *models.ImageModel, err error)
	Update(i *models.ImageModel) (image *models.ImageModel, err error)
	FindByID(id int64) (image *models.ImageModel, err error)
	FindByHash(hash string) (image *models.ImageModel, err error)
}

func NewImageRepo(db *gorm.DB) ImageRepoInterface {
	return &imageRepo{
		DB: db,
	}
}

type imageRepo struct {
	DB *gorm.DB
}

func (r *imageRepo) InsertOrUpdate(icon *models.ImageModel) (image *models.ImageModel, err error) {
	image, err = r.FindByID(icon.ID)
	if image != nil {
		// update
		return r.Update(icon)
	}
	err = r.DB.Model(&models.ImageModel{}).Create(icon).Error
	return icon, err
}

func (r *imageRepo) Update(i *models.ImageModel) (image *models.ImageModel, err error) {
	err = r.DB.Model(&models.ImageModel{}).Where("id = ?", i.ID).Updates(i).Error
	return i, err
}

func (r *imageRepo) FindByID(id int64) (image *models.ImageModel, err error) {
	err = r.DB.Model(&models.ImageModel{}).Where("id = ?", id).First(&image).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return image, err
}

func (r *imageRepo) FindByHash(hash string) (*models.ImageModel, error) {
	var icon models.ImageModel
	err := r.DB.Where("hash = ?", hash).First(&icon).Error
	if err != nil {
		return nil, err
	}
	return &icon, nil
}
