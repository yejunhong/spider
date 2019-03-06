package model

import (
	// "fmt"
	"github.com/jinzhu/gorm"
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

/**
 *
 * 获取漫画资列表
 * @param resource_name 模糊查询 资源名称
 * @param show_num 显示行数
 * @param start_num 开始行数
 * @return []CartoonResource{}
 *
 */
 func (model *Model) GetCartoons(resource_name string, show_num int64, start_num int64) []CartoonList{
	
	// 分页资源
	var CartoonsDb *gorm.DB = model.Db.Limit(show_num).Offset(start_num)

	if resource_name != "" { // 检索资源名称
		CartoonsDb = CartoonsDb.Where("resource_name LIKE ?", "%" + resource_name + "%")
	}

	var cartoons []CartoonList = []CartoonList{}
	CartoonsDb.Find(&cartoons) // 执行sql

	return cartoons
}