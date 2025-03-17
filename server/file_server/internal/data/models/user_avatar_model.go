package models

type UserAvatarModel struct {
	BaseModel
	FileName string `json:"file_name"`
	FileType string `json:"file_type"`
	FileSize int64  `json:"file_size"`
	FileURL  string `json:"file_url"`
}

func (*UserAvatarModel) TableName() string {
	return "user_avatar"
}
