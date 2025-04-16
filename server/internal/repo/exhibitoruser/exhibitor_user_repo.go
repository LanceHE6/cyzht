package exhibitoruser

import (
	"github.com/jinzhu/gorm"
	"server/internal/db"
	"server/internal/model"
)

type RepoInterface interface {
	Insert(uid, eid int64) error
	SelectByUIDAndAID(uid, aid int64) (*[]model.ExhibitorUserModel, error)
	SelectByAID(aid int64) (*[]model.ExhibitorUserModel, error)
	DeleteByAID(aid int64) error
	Delete(uid, aid int64) error
	update(exhibitorUser *model.ExhibitorUserModel) error
}

type exhibitorUserRepo struct {
	MyDB *gorm.DB
}

func (e exhibitorUserRepo) modelDB() *gorm.DB {
	return e.MyDB.Model(&model.ExhibitorUserModel{}).Preload("Exhibitor").Preload("Exhibitor.Activity")
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

func (e exhibitorUserRepo) SelectByUIDAndAID(uid, aid int64) (*[]model.ExhibitorUserModel, error) {
	var eu []model.ExhibitorUserModel
	return &eu, e.modelDB().Joins("JOIN exhibitor ON exhibitor_user.exhibitor_id = exhibitor.id").
		Where("exhibitor_user.user_id = ? AND exhibitor.activity_id = ?", uid, aid).Find(&eu).Error
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
	return e.modelDB().Where("user_id = ? and exhibitor_id = ?", uid, aid).Delete(&model.ExhibitorUserModel{}).Error
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
