package service
 
import (
    "fmt"
    "os"
    "net/http"
    "io/ioutil"
    "encoding/json"
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
        donwloadFile("cover/" + cartoon.ResourceNo, v.ResourceImgUrl)
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
            "content": "",
            "is_free": v.IsFree,
            "status": 0,
            "resource_url": v.ResourceUrl,
            "resource_name": v.ResourceName,
            "resource_img_url": v.ResourceImgUrl,
            "book_type": cartoon.BookType,
            "cdate": lib.Time(),
        })
        donwloadFile("chapter/" + cartoon.ResourceNo, v.ResourceImgUrl)
    }
    if len(data) > 0 {
        var updateInfo = map[string]interface{}{"status": 1}
        if chapter.Detail.Tags != "" {
            updateInfo["tags"] = chapter.Detail.Tags
        }
        if chapter.Detail.Detail != ""  {
            updateInfo["detail"] = chapter.Detail.Detail
        }
        if chapter.Detail.IsEnd != ""  {
            updateInfo["is_end"] = chapter.Detail.IsEnd
        }
        service.Models.UpdateCartoonListById(cartoonInfo.Id, updateInfo)
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
            donwloadFile("content/" + cartoon.ResourceNo + "/" + cartoonChapter.ListUniqueId + "/" + cartoonChapter.UniqueId, v.ResourceImgUrl)
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

/**
 *
 * 下载文件
 * @param p string 路径
 * @param url string 下载url
 *
 */
 func donwloadFile(p string, url string) {
    if url != "" {
        var path = "/Volumes/book/" + p + "/" + lib.MD5(url) + ".jpg"
        lib.DonwloadFile(path, url)
    }
}

type Data struct {
	Img string
	Md5 string
	Img2 string
}
type ImgUpload struct {
	Code int64
	Message string
	Data Data
}
func UploadImg(imgUrl string) *ImgUpload {

    img := &ImgUpload{}
    resp, err_ := http.Get("http://upload.manhua118.com/Img/Index/load?url=" + imgUrl)
    if err_ != nil {
        return img
    }
    defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return img
	}
	var errs = json.Unmarshal(body, img)
	if errs != nil {
	    return img
	}
	return img
}

/**
 *
 * 下载书籍图片
 * @param bookId 书籍ID
 * @param resourceUrl 书籍封面
 *
 */
 func (service *Service) DownloadBookIdImg(bookId int64, resourceUrl string){
    var UploadFileContent = UploadImg(resourceUrl)
    service.Models.UpdateCartoonListById(bookId, map[string]interface{}{"download_img_url": UploadFileContent.Data.Img})
}

/**
 *
 * 下载书籍章节图片
 * @param UniqueId 书籍唯一ID
 *
 */
 func (service *Service) DownloadBookIdChaptersImg(UniqueId string){
    var chapters = service.Models.GetChaptersImgByListUniqueId(UniqueId)
    var count = len(chapters)
    fmt.Println("\r\n同步章节图片")
    for k, img := range chapters {
        var UploadFileContent = UploadImg(img.ResourceImgUrl)
        var filePath string = UploadFileContent.Data.Img
        if filePath != "" {
            service.Models.UpdateCartoonChapterById(img.Id, map[string]interface{}{"download_img_url": filePath})
        }
        fmt.Printf("\r图片同步：%d/%d-处理总数：%s", (k + 1), count, filePath)
        os.Stdout.Sync()
    }
}

/**
 *
 * 下载书籍章节内容图片
 * @param UniqueId 书籍唯一ID
 *
 */
func (service *Service) DownloadBookIdContentImg(UniqueId string){
    var content = service.Models.GetContentsImgFindByListUniqueId(UniqueId)
    var count = len(content)
    fmt.Println("\r\n同步章节内容图片")
    for k, img := range content {
        var UploadFileContent = UploadImg(img.ResourceUrl)
        var filePath string = UploadFileContent.Data.Img
        if filePath != "" {
            /*go func(Id int64, file string){
                service.Models.UpdateCartoonContentById(Id, map[string]interface{}{"download_img_url": file})
            }(img.Id, filePath)*/
            service.Models.UpdateCartoonContentById(img.Id, map[string]interface{}{"download_img_url": filePath})
        }
        fmt.Printf("\r图片同步：%d/%d-处理总数：%s", (k + 1), count, filePath)
        os.Stdout.Sync()
    }
}