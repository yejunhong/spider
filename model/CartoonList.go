package model

import (
	"time"
)

type CartoonList struct{
	Id int64
	Unique_id string
	Title string
	Tags string
	Author string
	Detail string
	Status int64
	Resource_url string
	Resource_name string
	Resource_img_url string
	Download_img_url string
	Cdate int64
}


/**
 *
 * 创建一条漫画资源信息
 * @param resource_url 资源地址
 * @param resource_name 资源名称
 * @param config_name 资源使用的配置
 * @return CartoonResource{}
 *
 */
 func (model *Model) CreateCartoonList(resource_url, resource_name, config_name string) CartoonList{
	var cartoon_list CartoonList = CartoonList{
		Unique_id string
		Title string
		Tags string
		Author string
		Detail string
		Status int64
		Resource_url string
		Resource_name string
		Resource_img_url string
		Download_img_url string
		Cdate: time.Now().Unix()
	}
	model.Db.Create(&cartoon_list)
	return cartoon_list
}