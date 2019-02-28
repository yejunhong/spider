package model

import (
	// "fmt"
)

type CartoonList struct{
	Id int64
	ResourceNo string
	UniqueId string
	Title string
	Tags string
	Author string
	Detail string
	Status int64
	ResourceUrl string
	ResourceName string
	ResourceImgUrl string
	DownloadImgUrl string
	IsFree int64
	IsEnd int64
	Cdate int64
}

/**
 *
 * 通过id 获取漫画资源书籍
 * @return CartoonResource{}
 *
 */
 func (model *Model) GetCartoonInfoById(id int64) CartoonList{
	var cartoon_list CartoonList = CartoonList{}
	model.Db.Where("id = ?", id).First(&cartoon_list)
	return cartoon_list
}

/**
 *
 * 通过id 获取漫画资源书籍
 * @return CartoonResource{}
 *
 */
 func (model *Model) GetCartoonListByNo(no string) []CartoonList{
	var cartoon_list []CartoonList = []CartoonList{}
	model.Db.Where("resource_no = ? AND status = 0", no).Find(&cartoon_list)
	return cartoon_list
}

/**
 *
 * 通过id 修改漫画是否完结状态
 * @param id int64 漫画id
 * @param udata string 漫画是否完结
 *
 */
 func (model *Model) UpdateCartoonListById(id int64, udata map[string]interface{}){
	model.Db.Table("cartoon_list").Where("id = ?", id).Updates(udata)
}
