package req


type OrderReq struct {
	Uid uint `json:"uid"`
	AddressId uint `json:"address_id"`
	Pids []int `json:"pids"`
}

type UserReq struct {
	Uid uint `json:"uid" binding:"required"`
	NikeName string `json:"nike_name" binding:"required"`
	RealName string `json:"real_name" binding:"required"`
}

type CollectReq struct {
	Uid uint `json:"uid" binding:"required"`
	Pid uint `json:"pid" binding:"required"`
}