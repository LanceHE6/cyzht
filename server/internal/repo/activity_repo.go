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
}

type activityRepo struct {
	MyDB *gorm.DB
}

func (a *activityRepo) SelectByID(id int64) (*model.ActivityModel, error) {
	var activity model.ActivityModel
	err := a.modelDB().Where("id = ?", id).First(&activity).Error
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

func NewActivityRepo(mysqlConn *gorm.DB) ActivityRepoInterface {
	return &activityRepo{
		MyDB: mysqlConn,
	}
}
