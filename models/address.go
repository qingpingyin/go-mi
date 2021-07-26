package models

type Address struct {
	Id uint `json:"id" gorm:"primaryKey;not null;autoIncrement;comment:'主键'"`
	Uid uint `json:"uid" gorm:"type:bigint(20);not null;comment:'userid'"`
	ReceiverName string `json:"receiver_name" gorm:"type:varchar(50);default:'';comment:'收货人姓名'" binding:"required" `
	ReceiverMobile string `json:"receiver_mobile" gorm:"type:varchar(20);default:'';comment:'收货手机'" binding:"required,mobile" `
	ReceiverProvince string `json:"receiver_province" gorm:"type:varchar(50);default:'';comment:'省份'" binding:"required" `
	ReceiverCity string `json:"receiver_city" gorm:"type:varchar(50);default:'';comment:'市'" binding:"required" `
	ReceiverDistrict string `json:"receiver_district" gorm:"type:varchar(50);default:'';comment:'区'" binding:"required" `
	ReceiverAddress string `json:"receiver_address" gorm:"type:varchar(100);default:'';comment:'详细地址'" binding:"required" `
}

// 根据条件获取用户详情
func GetAddressByWhere(where ...interface{}) (address Address,err error) {
	err = Db.First(&address, where...).Error
	return
}
func GetAddressBy(where ...interface{})(ads []Address,err error){
	err = Db.Find(&ads, where...).Error
	return
}

func (a *Address)CreateAddress()error{
	err := Db.Save(a).Error
	return err
}
func (a *Address)Count(uid int)(count int64){
	Db.Model(a).Where("uid=?",uid).Count(&count)
	return
}
func (a *Address)DeleteAddressById()error{
	err := Db.Delete(a).Error
	return err
}