package exhibitoruser

import (
	"github.com/jinzhu/gorm"
	"server/internal/db"
	"server/internal/model"
)

type RepoInterface interface {
	Insert(uid, eid int64) error
	SelectByUID(uid int64) (*[]model.ExhibitorUserModel, error)
	SelectByAID(aid int64) (*[]model.ExhibitorUserModel, error)
	DeleteByAID(aid int64) error
	Delete(uid, aid int64) error
	update(exhibitorUser *model.ExhibitorUserModel) error
}

type exhibitorUserRepo struct {
	MyDB *gorm.DB
}

func (e exhibitorUserRepo) modelDB() *gorm.DB {
	return e.MyDB.Model(&model.ExhibitorUserModel{}).Preload("Exhibitor.Activity")
}

func (e exhibitorUserRepo) Insert(uid, eid int64) error {
	// 如果已加入则忽略
	if e.modelDB().Where("user_id = ? and exhibitor_id = ?", uid, eid).First(&model.ExhibitorUserModel{}).Error == nil {
		return nil
	}
	eu := model.ExhibitorUserModel{
		UserID:      uid,
		ExhibitorID: eid,
	}
	return e.modelDB().Create(&eu).Error
}

func (e exhibitorUserRepo) SelectByUID(uid int64) (*[]model.ExhibitorUserModel, error) {
	var exhibitorUser []model.ExhibitorUserModel
	return &exhibitorUser, e.modelDB().Where("user_id = ?", uid).Find(&exhibitorUser).Error
}

func (e exhibitorUserRepo) SelectByAID(aid int64) (*[]model.ExhibitorUserModel, error) {
	//TODO implement me
	panic("implement me")
}

func (e exhibitorUserRepo) DeleteByAID(aid int64) error {
	//TODO implement me
	panic("implement me")
}

func (e exhibitorUserRepo) Delete(uid, aid int64) error {
	//TODO implement me
	panic("implement me")
}

func (e exhibitorUserRepo) update(exhibitorUser *model.ExhibitorUserModel) error {
	//TODO implement me
	panic("implement me")
}

func NewExhibitorUserRepo(conn *db.DBConn) RepoInterface {
	return &exhibitorUserRepo{
		MyDB: conn.MySQLConn,
	}
}
