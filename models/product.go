package models

import "time"

type Product struct {
	Id uint `json:"id" gorm:"primaryKey;not null;autoIncrement;comment:'主键'"`
	Cid uint `json:"cid" gorm:"type:bigint(20);not null;comment:'外键'"`
	MarketPrice float64 `json:"market_price" gorm:"type:double(10,2);comment:'原价'"`
	ShopPrice  float64 `json:"shop_price" gorm:"type:double(10,2);not null;comment:'商城价'"`
	Title string `json:"title" gorm:"type:varchar(50);not null;comment:'商品标题'"`
	Image string `json:"image" gorm:"type:text;not null;comment:'缩略图'"`
	ImgUrl string `json:"img_url" gorm:"type:text;not null;comment:'封面图'"`
	Description string `json:"description" gorm:"type:text;not null;comment:'商品描述'"`
	Flag bool `json:"flag" gorm:"type:tinyint(4);not null;default:1;comment:'0 下架 1 正常'"`
	Inventory uint `json:"inventory" gorm:"type:bigint(20);not null;comment:'库存'"`
	CreatedAt time.Time `json:"create_at" gorm:"comment:'上架时间'"`
	UpdatedAt  time.Time `json:"update_at" gorm:"comment:'更新时间'"`
	ProductImg []ProductImg `json:"product_img" gorm:"foreignKey:Pid;references:Id"`
}

func GetAllProductBy(page ,pageSize int,where ...interface{})([]Product, error){
	var products []Product
	offset := GetOffset(page,pageSize)
	err := Db.Model(&Product{}).Preload("ProductImg").Limit(pageSize).Offset(offset).Find(&products,where...).Error
	return products,err
}
func GetProductByWhere(where ...interface{})(Product,error){
	var product Product
	err := Db.Preload("ProductImg").First(&product, where...).Error
	return product,err
}