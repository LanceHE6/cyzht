package activity

import (
	"context"
	"errors"
	"github.com/jinzhu/gorm"
	"server/internal/config"
	"server/internal/db"
	"server/internal/model"
	"server/pkg/logger"
	"server/pkg/rpc/file_server/api/v1/file_server"
	"time"
)

type RepoInterface interface {
	// Insert 插入并返回id
	Insert(activity *model.ActivityModel) (int64, error)
	// SelectByID 依id查询
	SelectByID(id int64) (*model.ActivityModel, error)
	// DeleteByID 删除
	DeleteByID(id int64) error
	// Search 查询
	// 使用option模式
	// 使用示例: Search(WithName(&name), WithPage(&page,&limit))
	Search(params ...PagingParams) (*[]model.ActivityModel, int, error)
	// update 更新
	update(activity *model.ActivityModel) error
	// UpdateIcon 更新展会头像地址
	UpdateIcon(id int64, avatarUrl string) error
	// UploadIcon 上传展会头像并更新头像地址
	UploadIcon(id int64, filename string, data []byte) error
}

type activityRepo struct {
	C             *config.Config
	MyDB          *gorm.DB
	FileRpcServer file_server.FileServiceClient
}

func (a *activityRepo) modelMyDB() *gorm.DB {
	return a.MyDB.Model(&model.ActivityModel{})
}

type PagingParams func(*gorm.DB) *gorm.DB

// Search 查询
// 使用option模式
// 使用示例: Search(WithName(&name), WithPage(&page,&limit))
func (a *activityRepo) Search(params ...PagingParams) (*[]model.ActivityModel, int, error) {
	db := a.modelDB().Preload("Creator") // 预加载创建者信息
	for _, param := range params {
		db = param(db)
	}
	var activities []model.ActivityModel
	// 获取总数 需清除原有的分页,防止影响总数
	countDB := db.Offset(-1).Limit(-1)
	var total int
	if err := countDB.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := db.Find(&activities).Error; err != nil {
		return nil, 0, err
	}

	return &activities, total, nil
}

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

// WithName 根据名称搜索
func WithName(name *string) PagingParams {
	return func(db *gorm.DB) *gorm.DB {
		if name != nil {
			db = db.Where("name LIKE ?", "%"+*name+"%")
		}
		return db
	}
}

// WithID 根据ID搜索
func WithID(ID *int64) PagingParams {
	return func(db *gorm.DB) *gorm.DB {
		if ID != nil {
			db = db.Where("id = ?", *ID)
		}
		return db
	}
}

// WithCreator 根据创建者搜索
func WithCreator(creator *string) PagingParams {
	return func(db *gorm.DB) *gorm.DB {
		// TODO 实现根据用户名搜索
		if creator != nil {
			db = db.Preload("user").Where("creator = ?", *creator)
		}
		return db
	}
}

// WithLocation 根据举办地点搜索
func WithLocation(location *string) PagingParams {
	return func(db *gorm.DB) *gorm.DB {
		if location != nil {
			db = db.Where("location LIKE ?", "%"+*location+"%")
		}
		return db
	}
}

// WithKeyword 根据关键字搜索(关键字范围: 活动名称,活动地点,活动描述)
func WithKeyword(keyword *string) PagingParams {
	return func(db *gorm.DB) *gorm.DB {
		if keyword != nil {
			db = db.Where("name LIKE ? or introduce LIKE ? or location LIKE ?",
				"%"+*keyword+"%",
				"%"+*keyword+"%",
				"%"+*keyword+"%",
			)
		}
		return db
	}
}

// WithIsInProgress 根据活动是否进行中搜索
func WithIsInProgress(isInProgress *bool) PagingParams {
	return func(db *gorm.DB) *gorm.DB {
		if isInProgress != nil {
			// 根据当前时间判断活动是否正在进行中
			if *isInProgress {
				db = db.Where("start_at <= ? and end_at >= ?", time.Now(), time.Now())
			} else {
				db = db.Where("start_at > ? or end_at < ?", time.Now(), time.Now())
			}
		}
		return db
	}
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
	err := a.modelMyDB().Where("id = ? AND is_deleted = ?", id, false).First(&activity).Error
	if err != nil {
		return nil, err
	}

	return &activity, nil
}

func (a *activityRepo) modelDB() *gorm.DB {
	return a.MyDB.Model(&model.ActivityModel{})
}

func (a *activityRepo) Insert(activity *model.ActivityModel) (int64, error) {
	err := a.modelDB().Create(&activity).Error
	if err != nil {
		return -1, err
	}
	// 若插入成功activity会自动赋值id
	return activity.ID, nil
}

func (a *activityRepo) update(activity *model.ActivityModel) error {
	return a.modelDB().Save(&activity).Error
}

// UploadIcon
//
//	@Description: 上传活动图标至文件服务器
//	@receiver u userRepo
//	@param id int64 用户id
//	@param filename string 文件名
//	@param data []byte 文件数据
//	@return error 错误信息
func (a *activityRepo) UploadIcon(id int64, filename string, data []byte) error {
	// 获取文件后缀
	rep, err := a.FileRpcServer.UploadActivityIcon(context.Background(), &file_server.UploadAvatarOrIconRequest{
		Id:          id,
		FileContent: data,
		FileName:    filename,
	})
	if err != nil || rep.FileUrl == "" {
		if err != nil {
			logger.Logger.Errorf("文件上传服务器失败: %s", err.Error())
		}
		return errors.New("文件上传服务器失败")
	}
	return a.UpdateIcon(id, rep.FileUrl)
}

// UpdateIcon
//
//	@Description: 更新数据库图标地址
//	@receiver u
//	@param id
//	@param fileURL
//	@return error
func (a *activityRepo) UpdateIcon(id int64, avatarURL string) error {
	return a.modelMyDB().Where("id = ?", id).Update("icon", avatarURL).Error
}

func NewActivityRepo(c *config.Config, dbConn *db.DBConn, fileRpcServer file_server.FileServiceClient) RepoInterface {
	return &activityRepo{
		C:             c,
		MyDB:          dbConn.MySQLConn,
		FileRpcServer: fileRpcServer,
	}
}
