package cate

import (
	"MI/models"
)

func CateTree(pid int)[]models.Categories{
	categories := models.GetAllCate("parent_id=?", pid)
	for i, v := range categories {
		 if list,err := models.GetAllProductBy(1,6,"cid=?",int(v.Id));err == nil{
		 	categories[i].Product =list
		 }
		categories[i].Children = CateTree(int(v.Id))
	}
	return categories
}
