package model

import "gorm.io/gorm"

type Share struct {
	Id        int            `json:"id" gorm:"column:id;type:int(11);not null;primaryKey;autoIncrement"`
	Fid       string         `json:"fid" gorm:"column:fid;type:varchar(255);default:null"`
	Uid       int            `json:"uid" gorm:"column:uid;type:int(11);default:null"`
	HashId    string         `json:"hash_id" gorm:"column:hash_id;type:varchar(255);default:null"`
	Expire    int            `json:"expire" gorm:"column:expire;type:bigint(11);default:null"`
	Type      int            `json:"type" gorm:"column:type;type:int(11);default:null"`
	Password  string         `json:"password" gorm:"column:password;type:varchar(10);default:null"`
	CreatedAt LocalTime      `json:"created_at" gorm:"column:created_at;type:datetime;default:null"`
	UpdatedAt LocalTime      `json:"updated_at" gorm:"column:updated_at;type:datetime;default:null"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at;type:datetime;default:null"`
}
