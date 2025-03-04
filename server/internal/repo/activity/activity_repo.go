package activity

import (
	"github.com/jinzhu/gorm"
	"server/internal/db"
	"server/internal/model"
	"time"
)

type RepoInterface interface {
	// Insert 插入
	Insert(activity *model.ActivityModel) error
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
}

type activityRepo struct {
	MyDB *gorm.DB
}

type PagingParams func(*gorm.DB) *gorm.DB

// Search 查询
// 使用option模式
// 使用示例: Search(WithName(&name), WithPage(&page,&limit))
func (a *activityRepo) Search(params ...PagingParams) (*[]model.ActivityModel, int, error) {
	db := a.modelDB()
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
			db = db.Preloads("user").Where("creator = ?", *creator)
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

func NewActivityRepo(dbConn *db.DBConn) RepoInterface {
	return &activityRepo{
		MyDB: dbConn.MySQLConn,
	}
}
