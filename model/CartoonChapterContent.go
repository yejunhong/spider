package model

type CartoonChapterContent struct{
	Id int64
	ResourceNo string
	ListUniqueId string
	ChapterUniqueId string
	ResourceUrl string
	DownloadImgUrl string
	Cdate int64
}

/**
 *
 * 通过章节id 删除章节内容数据
 *
 */
func (model *Model) DeleteChapterContentByChapterUniqueId(ChapterUniqueId string){
	model.Db.Where("chapter_unique_id = ?", ChapterUniqueId).Delete(CartoonChapterContent{})
}

/**
 *
 * 获取漫画章节 内容
 * @param unique_id 漫画章节id
 * @return []CartoonChapterContent{}
 *
 */
func (model *Model) GetContentsFindByChapterUniqueId(unique_id string) []CartoonChapterContent{
	var cartoonChapters []CartoonChapterContent = []CartoonChapterContent{}
	model.Db.Where("chapter_unique_id = ?", unique_id).Find(&cartoonChapters) // 执行sql
	return cartoonChapters
}

/**
 *
 * 获取漫画所有内容
 * @param list_unique_id 漫画d
 * @return []CartoonChapterContent{}
 *
 */
func (model *Model) GetContentsFindByChapterListUniqueId(list_unique_id string) []CartoonChapterContent{
	var cartoonChapters []CartoonChapterContent = []CartoonChapterContent{}
	model.Db.Where("list_unique_id = ?", list_unique_id).Find(&cartoonChapters) // 执行sql
	return cartoonChapters
}

/**
 *
 * 获取漫画没有下载过的内容
 * @param list_unique_id 漫画d
 * @return []CartoonChapterContent{}
 *
 */
func (model *Model) GetContentsImgFindByListUniqueId(list_unique_id string) []CartoonChapterContent{
	var cartoonChapters []CartoonChapterContent = []CartoonChapterContent{}
	model.Db.Where("list_unique_id = ? AND download_img_url IS NULL", list_unique_id).Find(&cartoonChapters) // 执行sql
	return cartoonChapters
}

/**
 *
 * 通过id 修改漫画内容状态
 * @param id int64 漫画id
 * @param imgPath string 图片地址
 *
 */
func (model *Model) UpdateCartoonContentById(id int64, udata map[string]interface{}){
	model.Db.Table("cartoon_chapter_content").Where("id = ?", id).Updates(udata)
}