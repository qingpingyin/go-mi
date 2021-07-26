package models

import "time"

type TraceType uint

const (
	//登录
	TraceTypeLogin  = iota
	//登出
	TraceTypeOut
	//编辑
	TraceTypeEdit
	//删除
	TraceTypeDel
)
type Trace struct{
	Id uint `json:"id" gorm:"primaryKey;not null;autoIncrement;comment:'主键'"`
	Uid uint `json:"uid" gorm:"type:bigint(20);not null;default:0;comment:'用户id'"`
	Ip int `json:"ip" gorm:"type:int(10);default:0;comment:'ip地址'"`
	Ext string `json:"ext" gorm:"type:varchar(1000);default:'';comment:'扩展信息'"`
	CreatedAt  time.Time `json:"create_at" gorm:"not null;comment:'注册时间'"`
	Type  uint  `json:"type" gorm:"type:tinyint(4);not null;default:0;comment:'类型(0:登录1:退出2:修改3:删除)'"`
}

