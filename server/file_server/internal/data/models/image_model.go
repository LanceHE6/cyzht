package models

type FileType uint8

const (
	FileTypeImage FileType = iota + 1
	FileTypeDocument
	FileTypeVideo
	FileTypeAudio
)

// FileModel 文件存储表
type FileModel struct {
	BaseModel
	FileName  string   `gorm:"column:file_name;type:varchar(255);not null"`
	FileType  FileType `gorm:"column:file_type;type:tinyint;not null"`
	MimeType  string   `gorm:"column:mime_type;type:varchar(100);not null"`
	FileSize  int64    `gorm:"column:file_size;type:bigint;not null"`
	FileURL   string   `gorm:"column:file_url;type:varchar(512);not null;unique"` // 添加唯一约束
	ThumbURL  string   `gorm:"column:thumb_url;type:varchar(512)"`
	Width     int      `gorm:"column:width;type:int"`
	Height    int      `gorm:"column:height;type:int"`
	Duration  int      `gorm:"column:duration;type:int"`
	Hash      string   `gorm:"column:hash;type:varchar(64);not null;uniqueIndex"` // 添加唯一索引
	OwnerID   int64    `gorm:"column:owner_id;type:bigint"`
	OwnerType string   `gorm:"column:owner_type;type:varchar(50)"`
}

func (*FileModel) TableName() string {
	return "file"
}
