package req


type ReqAddress struct {
	ReceiverName string `json:"receiver_name"  binding:"required" `
	ReceiverMobile string `json:"receiver_mobile" binding:"required,mobile" `
	ReceiverProvince string `json:"receiver_province"  binding:"required" `
	ReceiverCity string `json:"receiver_city"  binding:"required" `
	ReceiverDistrict string `json:"receiver_district"  binding:"required" `
	ReceiverAddress string `json:"receiver_address"  binding:"required" `
}
