package model

import (
	// "fmt"
)

type CartoonResource struct{
	Id int64
	UniqueId string
	ResourceNo string
	ResourceUrl string
	ResourceName string
	ConfigName string
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
 func (model *Model) CreateCartoon(resource_url, resource_name, config_name string) CartoonResource{
	var cartoon CartoonResource = CartoonResource{
		ResourceUrl: resource_url,
		ResourceName: resource_name,
		ConfigName: config_name,
	}
	model.Db.Create(&cartoon)
	return cartoon
}

/**
 *
 * 通过id 获取漫画资源书籍
 * @return CartoonResource{}
 *
 */
func (model *Model) GetCartoonById(id int64) CartoonResource{
	var cartoon CartoonResource = CartoonResource{}
	model.Db.Where("id = ?", id).First(&cartoon)
	return cartoon
}