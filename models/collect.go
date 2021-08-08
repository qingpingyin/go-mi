package models

import "time"

type Collect struct {
	Id uint `json:"id" gorm:"primaryKey;not null;autoIncrement;comment:'主键'"`
	Pid uint `json:"pid" gorm:"type:bigint(20);not null;comment:'商品id'"`
	Uid uint `json:"uid" gorm:"type:bigint(20);not null;comment:'用户id'"`
	CreatedAt time.Time `json:"create_at" gorm:"not null;comment:'注册时间'"`
	UpdatedAt  time.Time `json:"update_at" gorm:"not null;comment:'更新时间'"`
}

func (c *Collect)CreateCollect()error{
	return Db.Create(&c).Error
}

func GetCollectByWhere(where ...interface{}) (Collect ,error) {
	var c Collect
	err = Db.First(&c, where...).Error
	return c,err
}
func GetCollectBy(where ...interface{}) (c []Collect) {
	Db.Find(&c, where...)
	return
}
func GetCollectCountBy(where ...interface{})(count int64){
	Db.Model(&Collect{}).Where(where[0],where[1:]...).Count(&count)
	return
}
func DelCollect(where ...interface{})error{
	return Db.Where(where[0],where[1:]...).Delete(&Collect{}).Error
}