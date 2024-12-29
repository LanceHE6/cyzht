package repo

import (
	"gorm.io/gorm"
)

type Repo struct {
	UserAvatarRepo UserAvatarRepoInterface
}

func NewRepository(db *gorm.DB) *Repo {
	return &Repo{
		UserAvatarRepo: NewUserAvatarRepo(db),
	}
}
