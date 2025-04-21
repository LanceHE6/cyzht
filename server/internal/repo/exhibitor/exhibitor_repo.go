package exhibitor

import (
	"github.com/jinzhu/gorm"
	"server/internal/db"
	"server/internal/model"
)

type RepoInterface interface {
	// Insert 插入
	Insert(exhibitor *model.ExhibitorModel) (*model.ExhibitorModel, error)
	// SelectByID 依id查询
	SelectByID(id int64) (*model.ExhibitorModel, error)
	// SelectByAID 根据展会id查询
	SelectByAID(aid int64) (*[]model.ExhibitorModel, error)
	// DeleteByID 删除
	DeleteByID(id int64) error
	// update 更新
	update(exhibitor *model.ExhibitorModel) error
}

type exhibitorRepo struct {
	MyDB *gorm.DB
}

func (e *exhibitorRepo) SelectByAID(aid int64) (*[]model.ExhibitorModel, error) {
	var exhibitor []model.ExhibitorModel
	err := e.modelDB().Preload("Creator").Where("activity_id = ?", aid).Find(&exhibitor).Error
	return &exhibitor, err
}

func (e *exhibitorRepo) DeleteByID(id int64) error {
	return e.modelDB().Delete(&model.ExhibitorModel{}, "id = ?", id).Error
}

func (e *exhibitorRepo) SelectByID(id int64) (*model.ExhibitorModel, error) {
	var exhibitor model.ExhibitorModel
	err := e.modelDB().Where("id = ?", id).First(&exhibitor).Error
	if err != nil {
		return nil, err
	}
	return &exhibitor, nil
}

func (e *exhibitorRepo) modelDB() *gorm.DB {
	return e.MyDB.Model(&model.ExhibitorModel{})
}

func (e *exhibitorRepo) Insert(exhibitor *model.ExhibitorModel) (*model.ExhibitorModel, error) {
	err := e.modelDB().Create(&exhibitor).Error
	if err != nil {
		return nil, err
	}
	return exhibitor, nil
}

func (e *exhibitorRepo) update(exhibitor *model.ExhibitorModel) error {
	return e.modelDB().Save(&exhibitor).Error
}

func NewExhibitorRepo(dbConn *db.DBConn) RepoInterface {
	return &exhibitorRepo{
		MyDB: dbConn.MySQLConn,
	}
}
