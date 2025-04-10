package repo

import (
	"context"
	"errors"
	"file_server/internal/data/models"
	"gorm.io/gorm"
	"time"
)

type FileRepoInterface interface {
	Upsert(file *models.FileModel) (*models.FileModel, error)
	FindByID(id int64) (*models.FileModel, error)
	FindByHash(hash string) (*models.FileModel, error)
	FindByIDs(ids []int64) ([]*models.FileModel, error)
	Delete(id int64) error
}

type fileRepo struct {
	db *gorm.DB
}

func NewFileRepo(db *gorm.DB) FileRepoInterface {
	return &fileRepo{db: db}
}

func (r *fileRepo) Upsert(file *models.FileModel) (*models.FileModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// SQLite 专用的 Upsert 实现
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 先尝试查询是否存在
		var existing models.FileModel
		if err := tx.Where("hash = ?", file.Hash).First(&existing).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// 不存在则创建
				return tx.Create(file).Error
			}
			return err
		}

		// 已存在则更新
		file.ID = existing.ID
		return tx.Model(&existing).Updates(file).Error
	})

	if err != nil {
		return nil, err
	}
	return file, nil
}

func (r *fileRepo) FindByID(id int64) (*models.FileModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var file models.FileModel
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&file).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &file, nil
}

func (r *fileRepo) FindByHash(hash string) (*models.FileModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var file models.FileModel
	err := r.db.WithContext(ctx).Where("hash = ?", hash).First(&file).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &file, nil
}

func (r *fileRepo) FindByIDs(ids []int64) ([]*models.FileModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var files []*models.FileModel
	err := r.db.WithContext(ctx).Where("id IN ?", ids).Find(&files).Error
	if err != nil {
		return nil, err
	}

	// 确保返回顺序与请求顺序一致
	result := make([]*models.FileModel, len(ids))
	fileMap := make(map[int64]*models.FileModel, len(files))
	for _, f := range files {
		fileMap[f.ID] = f
	}

	for i, id := range ids {
		if f, ok := fileMap[id]; ok {
			result[i] = f
		}
	}

	return result, nil
}

func (r *fileRepo) Delete(id int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return r.db.WithContext(ctx).Delete(&models.FileModel{}, id).Error
}
