package service
 
import (
    "spider/model"
    "spider/lib"
	Drive "spider/grpc"
)

type Service struct{
    Browser Drive.NodeBrowser
    Models *model.Model
}

func (service *Service) InitService(){
    service.Browser = Drive.NodeBrowser{}
    service.Browser.CreateBrowserClient() // 创建浏览器客户端
}

/**
 *
 * 爬取书籍列表
 * @param cartoon
 * @param url 需要爬取的链接
 * @param test 测试爬虫
 * @return []map[string]interface{}
 *
 */
func (service *Service) CrawlBookList(cartoon model.CartoonResource, url string, test bool) []map[string]interface{}{

    if url == "" {
        url = cartoon.ResourceUrl
    }

    var list_data *Drive.ListReply = service.Browser.CrawlList(url, cartoon.ConfigName) // 浏览器拉取列表数据
    var data []map[string]interface{}
    // 获取服务端返回的结果
    for _, v := range list_data.Data {
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
    
    if test == false {
        service.Models.BatchInsert("cartoon_list", data, []string{"tags", "author", "detail", "resource_url", "resource_name", "resource_img_url"})
        if list_data.Next != "" && len(data) > 0 { // 是否下一页
            service.CrawlBookList(cartoon, list_data.Next, test)
        }
    }

    return data
}

/**
 *
 * 爬取书籍章节列表
 * @param cartoon 资源配置信息
 * @param cartoonInfo 书籍信息
 * @param url 需要爬取的链接
 * @param test 测试爬虫
 * @return []map[string]interface{}
 *
 */
func (service *Service) CrawlBookChapter(cartoon model.CartoonResource, cartoonInfo model.CartoonList, url string, test bool) []map[string]interface{}{
    
    if url == "" {
        url = cartoonInfo.ResourceUrl
    }

    var chapter_data *Drive.ChapterReply = service.Browser.CrawlChapter(url, cartoon.ConfigName) // 浏览器拉取列表数据
    var data []map[string]interface{}
    var cartoon_is_free string = "0" // 是否收费 0免费

    for _, v := range chapter_data.Data {
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
    if test == false {
        if chapter_data.Detail != nil {
            service.Models.UpdateCartoonListById(cartoonInfo.Id, map[string]interface{}{"is_free": cartoon_is_free, "is_end": chapter_data.Detail.IsEnd, "status": 1})
        }
        service.Models.BatchInsert("cartoon_chapter", data, []string{"is_free", "resource_url", "resource_name", "resource_img_url"})
        if chapter_data.Next != "" && len(data) > 0 { // 是否下一页
            service.CrawlBookChapter(cartoon, cartoonInfo, chapter_data.Next, test)
        }
    }
    return data
}

/**
 *
 * 爬取书籍章节被人
 * @param cartoon 资源配置信息
 * @param cartoonChapter 书籍章节信息
 * @param url 需要爬取的链接
 * @param test 测试爬虫
 * @return []map[string]interface{}
 *
 */
func (service *Service) CrawlBookChapterContent(cartoon model.CartoonResource, cartoonChapter model.CartoonChapter, url string, test bool) []map[string]interface{}{
    
    if url == "" {
        url = cartoonChapter.ResourceUrl
    }

    var chapter_data *Drive.ChapterContentReply = service.Browser.CrawlChapterContent(url, cartoon.ConfigName) // 浏览器拉取列表数据
    var data []map[string]interface{}
    // 获取服务端返回的结果
    // 清除现有数据
    service.Models.DeleteChapterContentByChapterUniqueId(cartoonChapter.UniqueId)
    for _, v := range chapter_data.Data {
        data = append(data, map[string]interface{}{
            "resource_no": cartoon.ResourceNo,
            "list_unique_id": cartoonChapter.ListUniqueId,
            "chapter_unique_id": cartoonChapter.UniqueId,
            "resource_url": v.ResourceImgUrl,
            "cdate": lib.Time(),
        })
    }
    if test == false {
        service.Models.UpdateCartoonChapterById(cartoonChapter.Id, map[string]interface{}{"status": 1})
        service.Models.BatchInsert("cartoon_chapter_content", data, []string{})
        if chapter_data.Next != "" && len(data) > 0 { // 是否下一页
            service.CrawlBookChapterContent(cartoon, cartoonChapter, chapter_data.Next, test)
        }
    }
    return data
}