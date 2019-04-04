package service
 
import (
    "spider/model"
    "spider/lib"
	Drive "spider/grpc"
)

type Service struct{
    Models *model.Model
}

/**
 *
 * 爬取书籍列表
 * @param book 书籍信息
 * @param cartoon 资源配置信息
 *
 */
func (service *Service) RecordBook(book *Drive.Res, cartoon model.CartoonResource){
    // 获取服务端返回的结果
    var data []map[string]interface{}
    for _, v := range book.Data {
        data = append(data, map[string]interface{}{
            "resource_no": cartoon.ResourceNo,
            "unique_id": lib.MD5(cartoon.ResourceNo + v.ResourceName),
            "tags": v.Tags,
            "author": v.Author,
            "detail": v.Detail,
            "status": 0,
            "resource_url": v.ResourceUrl,
            "resource_name": v.ResourceName,
            "resource_img_url": v.ResourceImgUrl,
            "cdate": lib.Time(),
        })
    }
    if len(data) > 0 {
        service.Models.BatchInsert("cartoon_list", data, []string{"tags", "author", "detail", "resource_url", "resource_name", "resource_img_url"})
    }
}

/**
 *
 * 爬取书籍章节列表
 * @param chapter 章节列表
 * @param cartoon 资源配置信息
 * @param cartoonInfo 书籍信息
 * @return []map[string]interface{}
 *
 */
func (service *Service) RecordChapter(
                        chapter *Drive.ResChapter, 
                        cartoon model.CartoonResource, 
                        cartoonInfo model.CartoonList) {

    var cartoon_is_free string = "0"
    var data []map[string]interface{}
    for _, v := range chapter.Data {
        if v.IsFree == "1" && cartoon_is_free == "0" {
            cartoon_is_free = "1"
        }
        data = append(data, map[string]interface{}{
            "resource_no": cartoon.ResourceNo,
            "unique_id": lib.MD5(cartoonInfo.UniqueId + cartoon.ResourceNo + v.ResourceName),
            "list_unique_id": cartoonInfo.UniqueId,
            "conent": "",
            "is_free": v.IsFree,
            "status": 0,
            "resource_url": v.ResourceUrl,
            "resource_name": v.ResourceName,
            "resource_img_url": v.ResourceImgUrl,
            "cdate": lib.Time(),
        })
    }
    if len(data) > 0 {
        service.Models.UpdateCartoonListById(cartoonInfo.Id, map[string]interface{}{"status": 1})
        service.Models.BatchInsert("cartoon_chapter", data, []string{"is_free", "resource_url", "resource_name", "resource_img_url"})
    }
}

/**
 *
 * 爬取书籍章节内容
 * @param content 章节内容
 * @param cartoon 资源配置信息
 * @param cartoonChapter 书籍章节信息
 *
 */
func (service *Service) RecordContent(
                        content *Drive.ResContent, 
                        cartoon model.CartoonResource, 
                        cartoonChapter model.CartoonChapter) {
    if cartoon.BookType == 1 {
        var data []map[string]interface{}
        // 获取服务端返回的结果
        // 清除现有数据
        service.Models.DeleteChapterContentByChapterUniqueId(cartoonChapter.UniqueId)
        for _, v := range content.Data {
            data = append(data, map[string]interface{}{
                "resource_no": cartoon.ResourceNo,
                "list_unique_id": cartoonChapter.ListUniqueId,
                "chapter_unique_id": cartoonChapter.UniqueId,
                "resource_url": v.ResourceImgUrl,
                "cdate": lib.Time(),
            })
        }
        if len(data) > 0 {
            service.Models.UpdateCartoonChapterById(cartoonChapter.Id, map[string]interface{}{"status": 1})
            service.Models.BatchInsert("cartoon_chapter_content", data, []string{})
        }
    }
    
    if cartoon.BookType == 2 {
        service.Models.UpdateCartoonChapterById(cartoonChapter.Id, map[string]interface{}{
            "status": 1,
            "content": content.Data[0].ResourceImgUrl})
    }

}   