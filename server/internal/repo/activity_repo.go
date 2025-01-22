package repo

import (
	"github.com/jinzhu/gorm"
	"server/internal/model"
)

type ActivityRepoInterface interface {
	// Insert 插入
	Insert(activity *model.ActivityModel) error
	// SelectByID 依id查询
	SelectByID(id int64) (*model.ActivityModel, error)
	// DeleteByID 删除
	DeleteByID(id int64) error
	// update 更新
	update(activity *model.ActivityModel) error
}

type activityRepo struct {
	MyDB *gorm.DB
}

func (a *activityRepo) DeleteByID(id int64) error {
	target, err := a.SelectByID(id)
	if err != nil {
		return err
	}
	if target == nil {
		return nil
	}
	target.IsDeleted = true
	return a.update(target)
}

func (a *activityRepo) SelectByID(id int64) (*model.ActivityModel, error) {
	var activity model.ActivityModel
	err := a.modelDB().Where("id = ? and is_deleted = false", id).First(&activity).Error
	if err != nil {
		return nil, err
	}
	return &activity, nil
}

func (a *activityRepo) modelDB() *gorm.DB {
	return a.MyDB.Model(&model.ActivityModel{})
}

func (a *activityRepo) Insert(activity *model.ActivityModel) error {
	return a.modelDB().Create(&activity).Error
}

func (a *activityRepo) update(activity *model.ActivityModel) error {
	return a.modelDB().Save(&activity).Error
}

func NewActivityRepo(mysqlConn *gorm.DB) ActivityRepoInterface {
	return &activityRepo{
		MyDB: mysqlConn,
	}
}
