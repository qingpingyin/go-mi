package models

import (
	"time"
)

type Users struct {
	Id uint `json:"id" gorm:"primaryKey;not null;autoIncrement;comment:'主键'"`
	Email string `json:"email" gorm:"type:varchar(50);unique;not null;default:'';comment:'邮箱'"`
	CreatedAt time.Time `json:"create_at" gorm:"not null;comment:'注册时间'"`
	Ext string `json:"ext" gorm:"type:varchar(1000);default:'';comment:'扩展信息'"`
	UpdatedAt  time.Time `json:"update_at" gorm:"not null;comment:'更新时间'"`
	NikeName string `json:"nike_name" gorm:"type:varchar(50);not null;default:'';comment:'昵称'"`
	RealName string `json:"real_name" gorm:"type:varchar(50);not null;default:'';comment:'真实姓名'"`
	Password string `json:"-" gorm:"type:varchar(50);not null;default:'';comment:'密码'"`
	Mobile string `json:"mobile" gorm:"type:varchar(20);unique;not null;unique;default:'';comment:'手机号'"`
	Salt string `json:"-" gorm:"type:char(4);not null;comment:'盐值'"`
	Status uint `json:"status" gorm:"type:tinyint(4);not null;default:0;comment:'状态（0：未审核,1:通过 2删除）'"`
}

// 根据条件获取用户详情
func GetUserByWhere(where ...interface{}) (au Users,err error) {
	var u Users
	err = Db.First(&u, where...).Error
	return u,err
}
//根据条件所有查询用户详情
func GetUsersBy(where ...interface{}) (au []Users) {
	Db.Find(&au, where...)
	return
}

func (u *Users) SingIn(t *Trace, d *Device) error{
	tx := Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Create(&u).Error; err != nil {
		tx.Rollback()
		return err
	}
	t.Uid=u.Id
	if err := tx.Create(&t).Error; err != nil {
		tx.Rollback()
		return err
	}
	d.Uid=u.Id
	if err := tx.Create(&d).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}