package repo

import (
	"gorm.io/gorm"
)

type Repo struct {
	FileRepo FileRepoInterface
}

func NewRepository(db *gorm.DB) *Repo {
	return &Repo{
		FileRepo: NewFileRepo(db),
	}
}
