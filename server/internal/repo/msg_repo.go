package repo

import (
	"github.com/jinzhu/gorm"
	"server/internal/model"
)

type MsgRepoInterface interface {
	// Insert 插入
	Insert(msg *model.MsgModel) error
	// SelectByID 依id查询
	SelectByID(id int64) (*model.MsgModel, error)
	// DeleteByID 删除
	DeleteByID(id int64) error
	// update 更新
	update(msg *model.MsgModel) error
}

type msgRepo struct {
	MyDB *gorm.DB
}

func (e *msgRepo) DeleteByID(id int64) error {
	return e.modelDB().Delete(&model.MsgModel{}, "id = ?", id).Error
}

func (e *msgRepo) SelectByID(id int64) (*model.MsgModel, error) {
	var msg model.MsgModel
	err := e.modelDB().Where("id = ?", id).First(&msg).Error
	if err != nil {
		return nil, err
	}
	return &msg, nil
}

func (e *msgRepo) modelDB() *gorm.DB {
	return e.MyDB.Model(&model.MsgModel{})
}

func (e *msgRepo) Insert(msg *model.MsgModel) error {
	return e.modelDB().Create(&msg).Error
}

func (e *msgRepo) update(msg *model.MsgModel) error {
	return e.modelDB().Save(&msg).Error
}

func NewMsgRepo(mysqlConn *gorm.DB) MsgRepoInterface {
	return &msgRepo{
		MyDB: mysqlConn,
	}
}
