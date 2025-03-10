package models

type UserAvatarModel struct {
	AvatarModel
}

func (*UserAvatarModel) TableName() string {
	return "user_avatar"
}
