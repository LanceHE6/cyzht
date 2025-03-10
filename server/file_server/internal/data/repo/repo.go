package repo

import (
	"gorm.io/gorm"
)

type Repo struct {
	UserAvatarRepo     UserAvatarRepoInterface
	ActivityAvatarRepo ActivityAvatarRepoInterface
}

func NewRepository(db *gorm.DB) *Repo {
	return &Repo{
		UserAvatarRepo:     NewUserAvatarRepo(db),
		ActivityAvatarRepo: NewActivityAvatarRepo(db),
	}
}
