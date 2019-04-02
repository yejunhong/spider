package model

type CartoonChapter struct{
	Id int64
	ResourceNo string
	UniqueId string
	ListUniqueId string
	IsFree int64
	Status int64
	ResourceUrl string
	ResourceName string
	ResourceImgUrl string
	DownloadImgUrl string
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
 * 获取漫画章节列表
 * @param unique_id 漫画章节
 * @return []CartoonChapter{}
 *
 */
 func (model *Model) GetChaptersFindByListUniqueId(list_unique_id string) []CartoonChapter{
	var cartoonChapters []CartoonChapter = []CartoonChapter{}
	model.Db.Where("list_unique_id = ? AND status = 0 AND is_free = 0", list_unique_id).Find(&cartoonChapters) // 执行sql
	return cartoonChapters
}