package models

type Cart struct {
	Id int64 `json:"id"`
	Uid int `json:"uid"`
	CartItem  []CartItem `json:"cart_item"`
}

type CartItem struct {
	Cid int64 `json:"cid"`
	Product Product `json:"product"`
	Num int `json:"num"`
}
type Item struct {
	Uid string`json:"uid" binding:"required"`
	Pid string `json:"pid" binding:"required"`
	Num string  `json:"num" binding:"required"`
}

