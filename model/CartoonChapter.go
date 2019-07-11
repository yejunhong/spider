package model

import (
	"github.com/jinzhu/gorm"
)
type CartoonChapter struct{
	Id int64
	ResourceNo string
	UniqueId string
	ListUniqueId string
	IsFree int64
	Status int64
	Content string
	Detail string
	ResourceUrl string
	ResourceName string
	ResourceImgUrl string
	DownloadImgUrl string
	Sort int64
	BookType int64
	Cdate int64
}

/**
 *
 * 通过id 获取漫画资源书籍
 * @return CartoonResource{}
 *
 */
 func (model *Model) GetCartoonChapterInfoById(id int64) CartoonChapter{
	var cartoon_chapter_info CartoonChapter = CartoonChapter{}
	model.Db.Where("id = ?", id).First(&cartoon_chapter_info)
	return cartoon_chapter_info
}

/**
 *
 * 通过id 获取漫画资源书籍
 * @return CartoonResource{}
 *
 */
 func (model *Model) GetCartoonChapterListByNo(no string) []CartoonChapter{
	var cartoon_chapter []CartoonChapter = []CartoonChapter{}
	model.Db.Where("resource_no = ? AND is_free = 0 AND status = 0", no).Limit(20000).Find(&cartoon_chapter)
	return cartoon_chapter
}

/**
 *
 * 通过id 修改漫画涨价是否完结状态
 * @param id int64 漫画id
 * @param udata string 漫画是否完结
 *
 */
func (model *Model) UpdateCartoonChapterById(id int64, udata map[string]interface{}){
	model.Db.Table("cartoon_chapter").Where("id = ?", id).Updates(udata)
}

/**
 *
 * 通过id 修改漫画涨价是否完结状态
 * @param id int64 漫画id
 * @param udata string 漫画是否完结
 *
 */
func (model *Model) UpdateCartoonChapterByIds(ids []int64, udata map[string]interface{}){
	model.Db.Table("cartoon_chapter").Where("id IN (?)", ids).Updates(udata)
}

/**
 *
 * 获取漫画章节列表
 * @param unique_id 漫画章节
 * @return []CartoonChapter{}
 *
 */
 func (model *Model) GetChaptersFindByListUniqueId(list_unique_id string, status int64) []CartoonChapter{
	var cartoonChapters []CartoonChapter = []CartoonChapter{}
	var CartoonsDb *gorm.DB = model.Db.Where("list_unique_id = ?", list_unique_id)
	if status != -1 {
		CartoonsDb = CartoonsDb.Where("status = ?", status)
	}
	CartoonsDb.Find(&cartoonChapters) // 执行sql
	return cartoonChapters
}

type ChaptersCount struct {
	Number int64
	NotNumber int64
	ListUniqueId string
}
/**
 *
 * 通过ListUniqueId 获取漫画资源书籍
 * @param list_unique_id 漫画ID
 * @return []chaptersCount{}
 *
 */
 func (model *Model) GetChaptersFindByListUniqueIdCount(list_unique_id []interface{}) []ChaptersCount{
	var chaptersCount []ChaptersCount = []ChaptersCount{}
	model.Db.Table("cartoon_chapter").Raw(`SELECT count(list_unique_id) number, COUNT(IF(status = 0, true, null)) not_number, list_unique_id FROM cartoon_chapter 
					WHERE list_unique_id IN (?) GROUP BY list_unique_id`, 
					list_unique_id).Find(&chaptersCount)
	return chaptersCount
}

/**
 *
 * 获取漫画章节列表
 * @param unique_id 漫画章节
 * @return []CartoonChapter{}
 *
 */
 func (model *Model) GetChaptersImgByListUniqueId(list_unique_id string) []CartoonChapter{
	var cartoonChapters []CartoonChapter = []CartoonChapter{}
	model.Db.Where("list_unique_id = ? AND download_img_url IS NULL", list_unique_id).Find(&cartoonChapters) // 执行sql
	return cartoonChapters
}