package models

import "time"

type Notice struct {
	Id uint `json:"id" gorm:"primaryKey;not null;autoIncrement;comment:'主键'"`
	CreatedAt time.Time `json:"create_at" gorm:"comment:'创建时间'"`
	Ext string `json:"ext" gorm:"type:text;default:'';comment:'主题内容'"`
	OperationType uint `json:"operation_type" gorm:"type:int(10);comment:'(1:绑定2:解绑)'" `
	UpdatedAt  time.Time `json:"update_at" gorm:"comment:'更新时间'"`
}

// 根据条件获取用户详情
func GetNoticeByWhere(where ...interface{}) (au Notice,err error) {
	var notice Notice
	err = Db.First(&notice, where...).Error
	return notice,err
}


