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
	BookType int64
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
 * 通过id 获取漫画资源书籍
 * @return CartoonResource{}
 *
 */
 func (model *Model) GetCartoonListByNoStatus(no string, status int64) []CartoonList{
	var cartoon_list []CartoonList = []CartoonList{}
	model.Db.Where("resource_no = ? AND status = ?", no, status).Find(&cartoon_list)
	return cartoon_list
}

/**
 *
 * 通过id 获取漫画资源书籍
 * @return CartoonResource{}
 *
 */
 func (model *Model) GetSqlCartoonListByNoStatus(no string, status int64) []CartoonList{
	var cartoon_list []CartoonList = []CartoonList{}
	model.Db.Raw(`SELECT * FROM cartoon_list list
			LEFT JOIN (
				SELECT list_unique_id FROM cartoon_chapter WHERE resource_no = ? AND status = ? GROUP BY list_unique_id
			) chapter ON (chapter.list_unique_id = list.unique_id)
			WHERE list.resource_no = ? AND chapter.list_unique_id IS NULL`, 
		no, status, no).Find(&cartoon_list)
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

type ResCartoonList struct {
	CartoonList
	CdateText string
}
/**
 *
 * 获取漫画资列表
 * @param resource_no 资源编号
 * @param resource_name 模糊查询 资源名称
 * @param show_num 显示行数
 * @param start_num 开始行数
 * @return []CartoonResource{}, count
 *
 */
 func (model *Model) GetCartoons(resource_no string, resource_name string, show_num int64, start_num int64) ([]ResCartoonList, int64){
	
	// 分页资源
	var CartoonsDb *gorm.DB = model.Db

	if resource_name != "" { // 检索资源名称
		CartoonsDb = CartoonsDb.Where("resource_name LIKE ?", "%" + resource_name + "%")
	}

	if resource_no != "" { // 检索资源名称
		CartoonsDb = CartoonsDb.Where("resource_no = ?", resource_no)
	}

	var cartoons []ResCartoonList = []ResCartoonList{}
	var count int64
	CartoonsDb.Model(&CartoonList{}).Count(&count)
	CartoonsDb.Table("cartoon_list").Limit(show_num).Offset(start_num).Find(&cartoons) // 执行sql

	return cartoons, count
}

/**
 *
 * 通过UniqueId 获取漫画资源书籍
 * @return CartoonResource{}
 *
 */
 func (model *Model) GetCartoonInfoByUniqueId(UniqueId int64) CartoonList{
	var cartoon_list CartoonList = CartoonList{}
	model.Db.Where("unique_id = ?", UniqueId).First(&cartoon_list)
	return cartoon_list
}

type BookCount struct {
	Number int64
	ResourceNo string
}
/**
 *
 * 通过ListUniqueId 获取漫画资源书籍
 * @param list_unique_id 漫画ID
 * @return []chaptersCount{}
 *
 */
 func (model *Model) GetBookByResourceNoCount(ResourceNo []interface{}) []BookCount{
	var bookCount []BookCount = []BookCount{}
	model.Db.Table("cartoon_list").Raw(`SELECT count(resource_no) number, resource_no FROM cartoon_list 
					WHERE resource_no IN (?) GROUP BY resource_no`, ResourceNo).Find(&bookCount)
	return bookCount
}