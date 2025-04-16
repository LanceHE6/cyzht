package model

// ExhibitorUserModel 用户加入参展商绑定模型
type ExhibitorUserModel struct {
	BaseModel
	ExhibitorID int64           `gorm:"column:exhibitor_id;type:bigint;not null" json:"-"`
	Exhibitor   *ExhibitorModel `gorm:"foreignKey:ExhibitorID;references:ID" json:"exhibitor"`
	UserID      int64           `gorm:"column:user_id;type:bigint;not null" json:"-"`
	User        *UserModel      `gorm:"foreignKey:UserID;references:ID" json:"user"`
	Role        uint8           `gorm:"column:role;type:tinyint;default:1" json:"role"`
}

func (*ExhibitorUserModel) TableName() string {
	return "exhibitor_user"
}
