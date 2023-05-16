package model

type File struct {
	Id         int       `json:"id" gorm:"column:id;type:int(11);not null;primaryKey;autoIncrement"`
	Pid        int       `json:"pid" gorm:"column:pid;type:int(11);not null"`
	Uid        int       `json:"uid" gorm:"column:uid;type:int(11);not null"`
	Name       string    `json:"name" gorm:"column:name;type:varchar(255);not null"`
	Thumb      string    `json:"thumb" gorm:"column:thumb;type:varchar(255);not null"`
	Size       float64   `json:"size" gorm:"column:size;type:double;not null"`
	Type       string    `json:"type" gorm:"column:type;type:varchar(255);not null"`
	Path       string    `json:"path" gorm:"column:path;type:varchar(255);not null"`
	Favorite   int       `json:"favorite" gorm:"column:favorite;type:int(1);not null"`
	DeleteFlag int       `json:"delete_flag" gorm:"column:delete_flag;type:int(1);not null"`
	CreatedAt  LocalTime `json:"created_at" gorm:"column:created_at;not null"`
	UpdatedAt  LocalTime `json:"updated_at" gorm:"column:updated_at;not null"`
}
