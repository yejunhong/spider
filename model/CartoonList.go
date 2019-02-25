package model

import (
	// "time"
	"fmt"
	// "strings"
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
 func (model *Model) CreateCartoonList(data []map[string]interface{}) {
	
	// var field []string
	var values []interface{}
	
	for key, val := range data {
		
		var value []interface{}

		for k, v := range val {
			if key == 0 {
				fmt.Println(k)
				
			}

			value = append(value, v)
		}
		
		values = append(values, 1)
	}
	fmt.Println(values)
	/*
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
		Cdate: time.Now().Unix()
	}

	model.Db.Exec(fmt.Sprintf(`INSERT INTO t_goods_selected_attributes(goods_no, goods_attributes_no, source, addtime) VALUES %s `,
								strings.Join(sql_selected_attributes, ",")))

	model.Db.Create(&cartoon_list)
	return cartoon_list*/
}