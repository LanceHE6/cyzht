package models

type ActivityAvatarModel struct {
	AvatarModel
}

func (*ActivityAvatarModel) TableName() string {
	return "activity_avatar"
}
