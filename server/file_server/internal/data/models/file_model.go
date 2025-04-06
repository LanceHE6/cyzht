package models

type FileModel struct {
	BaseModel
	FileName string `json:"file_name"`
	FileType string `json:"file_type"`
	FileSize int64  `json:"file_size"`
	FileURL  string `json:"file_url"`
}

func (*FileModel) TableName() string {
	return "file"
}
