package model

import (
	"gorm.io/gorm"
	"weapp/app/utils"
)

type File struct {
	Id        int            `json:"id" gorm:"column:id;type:int(11);not null;primaryKey;autoIncrement"`
	Pid       int            `json:"pid" gorm:"column:pid;type:int(11);not null"`
	Uid       int            `json:"uid" gorm:"column:uid;type:int(11);not null"`
	Name      string         `json:"name" gorm:"column:name;type:varchar(255);not null"`
	Thumb     string         `json:"thumb" gorm:"column:thumb;type:varchar(255);not null"`
	Size      float64        `json:"size" gorm:"column:size;type:double;not null"`
	Type      string         `json:"type" gorm:"column:type;type:varchar(255);not null"`
	Path      string         `json:"path" gorm:"column:path;type:varchar(255);not null"`
	Favorite  int            `json:"favorite" gorm:"column:favorite;type:int(1);not null"`
	CreatedAt LocalTime      `json:"created_at" gorm:"column:created_at;type:datetime;default:null"`
	UpdatedAt LocalTime      `json:"updated_at" gorm:"column:updated_at;type:datetime;default:null"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at;type:datetime;default:null"`
}

// GetDirAndFile 遍历文件夹及子文件夹、文件
func GetDirAndFile(files []File, dirs []File) []File {
	for _, v := range files {
		dirs = append(dirs, v)
		if v.Type == "folder" {
			var children []File
			Db.Model(&File{}).Where("pid = ?", v.Id).Find(&children) // 查询v.ID下的文件
			subdirs := GetDirAndFile(children, dirs)
			dirs = append(dirs, subdirs...) // 将子文件夹、文件合并到dirs
		}
	}
	return dirs
}

// SaveDirAndFile 保存文件夹及子文件夹、文件
func SaveDirAndFile(info map[string]int, files []File, dirs []File) []File {
	for _, v := range files {
		if v.Type == "folder" {
			// 查询子文件夹下的文件
			var subFile []File

			Db.Where("pid = ?", v.Id).Find(&subFile) // 查询pid下的文件

			// 获取创建文件夹时生成的ID
			var file File
			file.Pid = info["pid"]
			file.Uid = info["uid"]
			file.Type = "folder"
			file.Name = v.Name
			file.CreatedAt = LocalTime(utils.GetDateTime())
			file.UpdatedAt = LocalTime(utils.GetDateTime())

			// 创建文件夹生成的ID
			Db.Create(&file)

			info["pid"] = file.Id
			subdirs := SaveDirAndFile(info, subFile, dirs)
			dirs = append(dirs, subdirs...)
		} else {
			v.Id = 0
			v.Pid = info["pid"]
			v.Uid = info["uid"]
			v.CreatedAt = LocalTime(utils.GetDateTime())
			v.UpdatedAt = LocalTime(utils.GetDateTime())
			dirs = append(dirs, v)
		}
	}
	return dirs
}
