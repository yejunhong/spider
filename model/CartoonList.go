package model

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
	model.Db.Where("resource_no = ?", no).Find(&cartoon_list)
	return cartoon_list
}