package msg

import (
	"github.com/jinzhu/gorm"
	"server/internal/db"
	"server/internal/model"
)

type RepoInterface interface {
	// Insert 插入
	Insert(msg *model.MsgModel) (*model.MsgModel, error)
	// SelectByID 依id查询
	SelectByID(id int64) (*model.MsgModel, error)
	// DeleteByID 删除
	DeleteByID(id int64) error
	// UpdateStatus 更新消息状态
	UpdateStatus(id int64, status uint8) error
	// update 更新
	update(msg *model.MsgModel) error

	// GetMsg 获取历史消息
	// 使用option模式
	// 使用示例: GetMsg(WithPage(&page,&limit))
	GetMsg(activityID, exhibitorID int64, params ...PagingParams) (*[]model.MsgModel, int, error)
}

type msgRepo struct {
	MyDB *gorm.DB
}

// GetMsg 获取历史消息
// 使用option模式
// 使用示例: GetMsg(WithPage(&page,&limit))
func (e *msgRepo) GetMsg(activityID, exhibitorID int64, params ...PagingParams) (*[]model.MsgModel, int, error) {
	db := e.modelDB().Where("activity_id = ? AND exhibitor_id = ?", activityID, exhibitorID)
	for _, param := range params {
		db = param(db)
	}

	var msgs []model.MsgModel
	// 获取总数 需清除原有的分页,防止影响总数
	countDB := db.Offset(-1).Limit(-1)
	var total int
	if err := countDB.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := db.Find(&msgs).Error; err != nil {
		return nil, 0, err
	}

	return &msgs, total, nil
}

type PagingParams func(*gorm.DB) *gorm.DB

// WithPage 设置分页
func WithPage(page, limit *int) PagingParams {
	return func(db *gorm.DB) *gorm.DB {
		if page == nil {
			page = new(int)
			*page = 1
		}
		// 如果page为-1,则不进行分页
		if *page == -1 {
			return db
		}
		if limit == nil || *limit < 0 {
			limit = new(int)
			*limit = 10
		}
		offset := (*page - 1) * (*limit)
		db = db.Offset(offset).Limit(*limit)
		return db
	}
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
	return e.MyDB.Model(&model.MsgModel{}).
		Preload("FromUser").
		Preload("ToUser").
		Preload("Exhibitor").
		Preload("Activity").
		Preload("Attachments")
}

func (e *msgRepo) Insert(msg *model.MsgModel) (*model.MsgModel, error) {
	if err := e.modelDB().Create(&msg).Error; err != nil {
		return nil, err
	}
	newMsg, _ := e.SelectByID(msg.ID)
	return newMsg, nil
}

func (e *msgRepo) update(msg *model.MsgModel) error {
	return e.modelDB().Save(&msg).Error
}

func (e *msgRepo) UpdateStatus(id int64, status uint8) error {
	msg, err := e.SelectByID(id)
	if err != nil {
		return err
	}
	msg.Status = status
	if err := e.update(msg); err != nil {
		return err
	}
	return nil
}

func NewMsgRepo(dbConn *db.DBConn) RepoInterface {
	return &msgRepo{
		MyDB: dbConn.MySQLConn,
	}
}
