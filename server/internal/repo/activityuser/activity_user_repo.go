package activityuser

import (
	"github.com/jinzhu/gorm"
	"server/internal/config"
	"server/internal/db"
	"server/internal/model"
	"server/pkg/rpc/file_server/api/v1/file_server"
)

type RepoInterface interface {
	Insert(uid, aid int64) error
	SelectByUID(uid int64) (*[]model.ActivityUserModel, error)
	SelectByAID(aid int64) (*[]model.ActivityUserModel, error)
	DeleteByAID(aid int64) error
	Delete(uid, aid int64) error
	update(activityUser *model.ActivityUserModel) error
}

type activityUserRepo struct {
	C             *config.Config
	MyDB          *gorm.DB
	FileRpcServer file_server.FileServiceClient
}

func (a *activityUserRepo) modelDB() *gorm.DB {
	return a.MyDB.Model(&model.ActivityUserModel{})
}
func (a *activityUserRepo) Insert(uid, aid int64) error {
	// 如果已加入则忽略
	if a.modelDB().Where("user_id = ? and activity_id = ?", uid, aid).First(&model.ActivityUserModel{}).Error == nil {
		return nil
	}
	au := model.ActivityUserModel{
		UserID:     uid,
		ActivityID: aid,
	}
	return a.modelDB().Create(&au).Error
}

func (a *activityUserRepo) SelectByUID(uid int64) (*[]model.ActivityUserModel, error) {
	var activityUser []model.ActivityUserModel
	err := a.modelDB().Preload("Activity").
		Joins("JOIN activity ON activity_user.activity_id = activity.id").
		Where("activity_user.user_id = ? AND activity.is_deleted = ?", uid, false).
		Find(&activityUser).Error
	// 拼接展会图标地址
	for i, au := range activityUser {
		if au.Activity.Avatar != "" {
			activityUser[i].Activity.Avatar = a.C.Server.FileServer.StaticURL + au.Activity.Avatar
		}
	}
	return &activityUser, err
}

func (a *activityUserRepo) SelectByAID(aid int64) (*[]model.ActivityUserModel, error) {
	var activityUser []model.ActivityUserModel
	err := a.modelDB().Where("activity_id = ?", aid).Find(&activityUser).Error
	// 拼接展会图标地址
	for i, au := range activityUser {
		if au.Activity.Avatar != "" {
			activityUser[i].Activity.Avatar = a.C.Server.FileServer.StaticURL + au.Activity.Avatar
		}
	}
	return &activityUser, err
}

func (a *activityUserRepo) DeleteByAID(aid int64) error {
	//TODO implement me
	panic("implement me")
}

func (a *activityUserRepo) Delete(uid, aid int64) error {
	return a.modelDB().Where("user_id = ? and activity_id = ?", uid, aid).Delete(&model.ActivityUserModel{}).Error
}

func (a *activityUserRepo) update(activityUser *model.ActivityUserModel) error {
	//TODO implement me
	panic("implement me")
}

func NewActivityUserRepo(c *config.Config, dbConn *db.DBConn, fileRpcServer file_server.FileServiceClient) RepoInterface {
	return &activityUserRepo{
		MyDB:          dbConn.MySQLConn,
		C:             c,
		FileRpcServer: fileRpcServer,
	}
}
