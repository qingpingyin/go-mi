package models

import "time"

type Device struct {
	Id uint `json:"id" gorm:"primaryKey;not null;autoIncrement;comment:'主键'"`
	Client string  `json:"client" gorm:"type:varchar(50);not null;default:'';comment:'客户端'"`
	CreatedAt  time.Time `json:"create_at "gorm:"not null;comment:'注册时间'"`
	Uid uint `json:"uid" gorm:"type:bigint(20);not null;default:0;comment:'用户id'"`
	Ext string `json:"ext" gorm:"type:varchar(1000);default:'';comment:'扩展信息'"`
	Ip int `json:"ip" gorm:"type:int(10);not null;default:0;comment:'ip地址'"`
	Model string `json:"model" gorm:"type:varchar(50);default:'';comment:'设备信息'"`
}
