package service
 
import (
    "fmt"
    "os"
    "time"
    "net/http"
    "io/ioutil"
    "encoding/json"
    "spider/model"
    "spider/lib"
    "sync"
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
    fmt.Println("抓取章节数量：", len(chapter.Data))
    for _, v := range chapter.Data {
        if v.IsFree == "1" && cartoon_is_free == "0" {
            cartoon_is_free = "1"
        }
        var UniqueId = lib.MD5(cartoonInfo.UniqueId + cartoon.ResourceNo + v.ResourceName)
        data = append(data, map[string]interface{}{
            "resource_no": cartoon.ResourceNo,
            "unique_id": UniqueId,
            "list_unique_id": cartoonInfo.UniqueId,
            "content": "",
            "is_free": v.IsFree,
            "status": 0,
            "detail": v.Detail,
            "resource_url": v.ResourceUrl,
            "resource_name": v.ResourceName,
            "resource_img_url": v.ResourceImgUrl,
            "book_type": cartoon.BookType,
            "cdate": lib.Time(),
            "sort": v.Sort,
        })
        donwloadFile("chapter/" + cartoon.ResourceNo + "/" + cartoonInfo.UniqueId, v.ResourceImgUrl)
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
        service.Models.BatchInsert("cartoon_chapter", data, []string{"is_free", "resource_url", "resource_name", "resource_img_url", "sort"})
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
        if len(content.Data) >= 1 {
            service.Models.UpdateCartoonChapterById(cartoonChapter.Id, map[string]interface{}{
                "status": 1,
                "content": content.Data[0].ResourceImgUrl})
        } 
    }

}   

/**
 *
 * 下载文件
 * @param p string 路径
 * @param url string 下载url
 *
 */
 func donwloadFile(p string, url string) string {
    var md5Url string = lib.MD5(url)
    if url != "" {
        var path = "/data/book/" + p + "/" + md5Url + ".jpg"
        lib.DonwloadFile(path, url)
    }
    return md5Url
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

    client := &http.Client{
        Timeout: time.Second * 10,
    }
    // var md5Url = donwloadFile("tmp", imgUrl)
    var uploadUrl = `http://upload.manhua118.com/Img/Index/load?url=http://ye153259.viphk1.ngrok.org/img/` + imgUrl + ".jpg"
    // fmt.Println(uploadUrl)
    resp, err_ := client.Get(uploadUrl)
    if err_ != nil {
        fmt.Println("请求超时", uploadUrl)
        return img
    }
    defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
        fmt.Println("获取body失败", uploadUrl)
		return img
	}
	var errs = json.Unmarshal(body, img)
	if errs != nil {
        fmt.Println("获取json参数错误", uploadUrl)
	    return img
    }
    // fmt.Println(imgUrl)
    // os.Remove("/data/book/" + imgUrl + ".jpg") // 上传完后 删除图片
	return img
}

/**
 *
 * 下载书籍图片
 * @param bookId 书籍ID
 * @param resourceUrl 书籍封面
 *
 */
 func (service *Service) DownloadBookIdImg(bookId int64, resourceNo string, resourceUrl string){
    var UploadFileContent = UploadImg("cover/" + resourceNo + "/" + lib.MD5(resourceUrl))
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
        var imgUrl = "chapter/" + img.ResourceNo +"/"+UniqueId+ "/" + lib.MD5(img.ResourceImgUrl)
        var UploadFileContent = UploadImg(imgUrl)
        var filePath string = UploadFileContent.Data.Img
        if filePath != "" {
            service.Models.UpdateCartoonChapterById(img.Id, map[string]interface{}{"download_img_url": filePath})
            // os.Remove(imgUrl + ".jpg") // 删除文件test.txt
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

    var waitGroup sync.WaitGroup
    
    var i chan int = make(chan int, 15)
    for k, img := range content {
        waitGroup.Add(1) // 增加计算器
        i <- 1
        go func(img model.CartoonChapterContent, k int) {
            var p = "content/" + img.ResourceNo + "/" + img.ListUniqueId + "/" + img.ChapterUniqueId + "/" + lib.MD5(img.ResourceUrl)
            var UploadFileContent = UploadImg(p)
            var filePath string = UploadFileContent.Data.Img
            if filePath != "" {
                service.Models.UpdateCartoonContentById(img.Id, map[string]interface{}{"download_img_url": filePath})
                fmt.Printf("\r图片同步：%d/%d-处理总数：%s", (k + 1), count, filePath)
                os.Stdout.Sync()
            }
            <- i
            waitGroup.Done()
        }(img, k)
    }
    waitGroup.Wait()
} 
/**
 *
 * 下载书籍章节内容图片
 * @param UniqueId 书籍唯一ID
 *

func (service *Service) DownloadBookIdContentImg(UniqueId string){
    var content = service.Models.GetContentsImgFindByListUniqueId(UniqueId)
    var count = len(content)
    fmt.Println("\r\n同步章节内容图片")
    var c chan int = make(chan int, 2)
    for k, img := range content {
        // fmt.Println("下载地址：", img.ResourceUrl)
        c <- 1
        go func(img model.CartoonChapterContent){
            var p = "content/" + img.ResourceNo + "/" + img.ListUniqueId + "/" + img.ChapterUniqueId + "/" + lib.MD5(img.ResourceUrl)
            var UploadFileContent = UploadImg(p)
            var filePath string = UploadFileContent.Data.Img
            if filePath != "" {
                service.Models.UpdateCartoonContentById(img.Id, map[string]interface{}{"download_img_url": filePath})
                fmt.Printf("\r图片同步：%d/%d-处理总数：%s", (k + 1), count, filePath)
                os.Stdout.Sync()
            }
            <-c
        }(img)
    }
} */