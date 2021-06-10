package models

import "time"

type ProductImg struct {
	Id uint `json:"id" gorm:"primaryKey;not null;autoIncrement;comment:'主键'"`
	Pid uint `json:"pid" gorm:"type:bigint(20);not null;comment:'外键'"`
	ImgUrl string `json:"img_url" gorm:"type:text;not null;comment:'商品轮播图'"`
	CreatedAt time.Time `json:"create_at" gorm:"comment:'上架时间'"`
	UpdatedAt  time.Time `json:"update_at" gorm:"comment:'更新时间'"`
}
