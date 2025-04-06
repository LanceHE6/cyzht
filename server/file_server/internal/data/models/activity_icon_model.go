package models

type ActivityIconModel struct {
	BaseModel
	FileName string `json:"file_name"`
	FileType string `json:"file_type"`
	FileSize int64  `json:"file_size"`
	FileURL  string `json:"file_url"`
}

func (*ActivityIconModel) TableName() string {
	return "activity_icon"
}
