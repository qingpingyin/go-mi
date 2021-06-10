package models

import (
	"time"
)

type Categories struct {
	Id uint `json:"id" gorm:"primaryKey;not null;autoIncrement;comment:'主键'"`
	CategoriesName string `json:"categories_name" gorm:"type:varchar(20);not null;comment:'商品类别'"`
	ParentId int `json:"parent_id" gorm:"type:bigint(20);not null;comment:'父分类id'"`
	IsNav bool `json:"is_nav" gorm:"type:tinyint(4);not null;default:0;comment:'状态（0：不是,1:是）'"`
	CreatedAt time.Time `json:"create_at" gorm:"not null;comment:'创建时间'"`
	UpdatedAt  time.Time `json:"update_at" gorm:"not null;comment:'更新时间'"`
	Product []Product `json:"product" gorm:"foreignKey:Cid;references:Id"`
	Children []Categories `json:"children" gorm:"-"`
}
//根据条件查询 单个Categories
func GetCategoriesByWhere(where ...interface{})(Categories,error){
	var cate Categories
	err := Db.Preload("Product").First(&cate,where...).Error
	return cate,err
}
//根据条件 查询所有Categories
func GetCategoriesBy(where ...interface{}) ([]Categories,error) {
	var list [] Categories
	err := Db.Find(&list,where...).Error
	return list,err
}

func GetAllCate (where ...interface{})(res []Categories){
	Db.Model(Categories{}).Find(&res,where...)
	return res
}