package models


type Order struct {
	OrderId string `json:"id" gorm:"type:varchar(50);primaryKey;not null;autoIncrement;comment:'主键'"`
	Uid uint `json:"uid" gorm:"type:bigint(20);not null;comment:'userid'"`
	Aid uint `json:"aid" gorm:"type:bigint(20);not null;comment:'地址id'"`
	PayStatus uint `json:"pay_status" gorm:"type:int(10);default:1;not null;comment:'支付状态(1:待支付,2:已支付)'"`
	Payment float64 `json:"payment" gorm:"varchar(50);not null;comment:'支付金额'"`
	PaymentType uint `json:"payment_type" gorm:"type:int(10);default:1;not null;comment:'支付方式'"`
	CreatedAt  int64 `json:"create_at" gorm:"not null;comment:'注册时间'"`
	UpdatedAt  int64 `json:"update_at" gorm:"not null;comment:'更新时间'"`
	OrderItem []OrderItem `json:"order_item" gorm:"foreignKey:OrderId"`
}

type OrderItem struct {
	Id uint `json:"id" gorm:"primaryKey;not null;autoIncrement;comment:'主键'"`
	OrderId string `json:"order_id" gorm:"type:varchar(50);not null;comment:'订单id'"`
	Pid int `json:"order_id" gorm:"int(10);not null;comment:'商品id'"`
	Num int `json:"num" gorm:"int(10);not null;comment:'商品数量 '"`
	Title string `json:"title" gorm:"type:varchar(50);not null;comment:'商品标题'"`
	Price  float64 `json:"price" gorm:"type:double(10,2);not null;comment:'单价'"`
	TotalPrice float64 `json:"total_price" gorm:"type:double(10,2);not null;comment:'总价格'"`
	ImgUrl string `json:"img_url" gorm:"type:text;not null;comment:'封面图'"`
}

func (o *Order)Create()error{
	tx := Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Create(&o).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func GetAllOrderBy(page,pageSize int,where ...interface{})([]Order,error){
	var orders []Order
	offset := GetOffset(page,pageSize)
	err := Db.Preload("OrderItem").Limit(pageSize).Offset(offset).Find(&orders,where...).Error
	return orders,err
}
func GetAllOrderByWhere(where ...interface{})(Order,error){
	var order Order
	err := Db.Preload("OrderItem").First(&order,where...).Error
	return order,err
}

func GetOrderCountBy(where ...interface{})(count int64){
	Db.Model(&Order{}).Where(where[0],where[1:]...).Count(&count)
	return
}