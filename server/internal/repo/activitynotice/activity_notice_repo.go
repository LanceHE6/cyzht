package activitynotice

import (
	"github.com/jinzhu/gorm"
	"server/internal/db"
	"server/internal/model"
)

type RepoInterface interface {
	Insert(activityNotice *model.ActivityNoticeModel) (*model.ActivityNoticeModel, error)
	SelectByID(id int64) (*model.ActivityNoticeModel, error)
	SelectByAID(activityID int64) (*[]model.ActivityNoticeModel, error)
	DeleteByID(id int64) error
	Update(activityNotice *model.ActivityNoticeModel) error
}

type activityNoticeRepo struct {
	MyDB *gorm.DB
}

func (a *activityNoticeRepo) Update(activityNotice *model.ActivityNoticeModel) error {
	return a.modelDB().Where("id = ?", activityNotice.ID).Update(&activityNotice).Error
}

func (a *activityNoticeRepo) SelectByAID(activityID int64) (*[]model.ActivityNoticeModel, error) {
	var activityNotice []model.ActivityNoticeModel
	err := a.modelDB().Preload("Activity").Preload("Creator").Where("activity_id = ?", activityID).Find(&activityNotice).Error
	return &activityNotice, err
}

func (a *activityNoticeRepo) modelDB() *gorm.DB {
	return a.MyDB.Model(&model.ActivityNoticeModel{})
}
func (a *activityNoticeRepo) Insert(activityNotice *model.ActivityNoticeModel) (*model.ActivityNoticeModel, error) {
	if err := a.modelDB().Create(&activityNotice).Error; err != nil {
		return nil, err
	}
	return activityNotice, nil
}

func (a *activityNoticeRepo) SelectByID(id int64) (*model.ActivityNoticeModel, error) {
	var activityNotice model.ActivityNoticeModel
	err := a.modelDB().Where("id = ?", id).First(&activityNotice).Error
	if err != nil {
		return nil, err
	}
	return &activityNotice, nil
}

func (a *activityNoticeRepo) DeleteByID(id int64) error {
	return a.modelDB().Where("id = ?", id).Delete(&model.ActivityNoticeModel{}).Error
}

func NewActivityNoticeRepo(dbConn *db.DBConn) RepoInterface {
	return &activityNoticeRepo{
		MyDB: dbConn.MySQLConn,
	}
}
