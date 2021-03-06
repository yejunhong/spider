package model

import (
	// "fmt"
	"github.com/jinzhu/gorm"
)

type CartoonResource struct{
	Id int64
	ResourceNo string
	ResourceUrl string
	ResourceName string
	ConfigName string
	BookType int64
}

/**
 *
 * 获取漫画资源列表
 * @param resource_name 模糊查询 资源名称
 * @param show_num 显示行数
 * @param start_num 开始行数
 * @return []CartoonResource{}, count
 *
 */
func (model *Model) GetCartoonResources(resource_name string, show_num int64, start_num int64) ([]CartoonResource, int64){
	
	// 分页资源
	var CartoonResourcesDb *gorm.DB = model.Db.Limit(show_num).Offset(start_num)

	if resource_name != "" { // 检索资源名称
		CartoonResourcesDb = CartoonResourcesDb.Where("resource_name LIKE ?", "%" + resource_name + "%")
	}

	var cartoon []CartoonResource = []CartoonResource{}
	var count int64

	CartoonResourcesDb.Model(&CartoonResource{}).Count(&count)
	CartoonResourcesDb.Find(&cartoon) // 执行sql

	return cartoon, count
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

/**
 *
 * 通过id 获取漫画资源书籍
 * @return CartoonResource{}
 *
 */
 func (model *Model) GetCartoonByResourceNo(ResourceNo string) CartoonResource{
	var cartoon CartoonResource = CartoonResource{}
	model.Db.Where("resource_no = ?", ResourceNo).First(&cartoon)
	return cartoon
}

/**
 *
 * 通过资源对象 创建|修改信息
 * @return CartoonResource{}
 *
 */
func (model *Model) SetCartoonResource(resource CartoonResource) CartoonResource{
	var cartoon CartoonResource = CartoonResource{}
	model.Db.Where(CartoonResource{ResourceNo: resource.ResourceNo}).Assign(resource).FirstOrCreate(&cartoon)
	return cartoon
}