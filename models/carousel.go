package models


//轮播图
type Carousel struct {
	Id uint `json:"id" gorm:"primaryKey;not null;autoIncrement;comment:'主键'"`
	Pid uint `json:"pid" gorm:"bigint(20);not null;comment:'商品id'" `
	ImgUrl string `json:"img_url" gorm:"type:varchar(200);not null;comment:'图片地址'"`
	IsPlay bool `json:"is_play" gorm:"type:tinyint(4);default:0;commit:'0不播放,1播放'" `
}
//查询轮播图所有数据
func GetCarouselList()([]Carousel,error){
	var carousels []Carousel

	err := Db.Find(&carousels,"is_play = true").Error
	return carousels,err
}