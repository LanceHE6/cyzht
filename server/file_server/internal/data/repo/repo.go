package repo

import (
	"gorm.io/gorm"
)

type Repo struct {
	UserAvatarRepo   UserAvatarRepoInterface
	ActivityIconRepo ActivityIconRepoInterface
	ImageRepo        ImageRepoInterface
}

func NewRepository(db *gorm.DB) *Repo {
	return &Repo{
		UserAvatarRepo:   NewUserAvatarRepo(db),
		ActivityIconRepo: NewActivityIconRepo(db),
		ImageRepo:        NewImageRepo(db),
	}
}
