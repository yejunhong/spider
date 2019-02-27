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
	model.Db.Where("resource_no = ?", no).Find(&cartoon_chapter)
	return cartoon_chapter
}

