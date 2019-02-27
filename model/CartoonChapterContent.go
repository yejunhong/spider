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