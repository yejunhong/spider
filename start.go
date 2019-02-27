package main
 
import (
    "fmt"
    "spider/model"
    "spider/lib"
	Drive "spider/grpc"
)

var browser Drive.NodeBrowser
var models model.Model

var i = 0
func test(){
    i = i + 1
    if i < 5 {
        fmt.Println(1)
        test()
    }
    
}

func main(){

    browser = Drive.NodeBrowser{}
    models = model.Model{Db: model.InitDb()}

    browser.CreateBrowserClient() // 创建浏览器客户端
    
    // 通过资源获取漫画列表
    var cartoon = models.GetCartoonById(2)
    GetList(cartoon.ResourceUrl, cartoon);

    // GetChapter(2)
    //GetChapterContent(1)
}

func GetList(resource_url string, cartoon model.CartoonResource){

    fmt.Println("爬取：", resource_url)

    var list_data *Drive.ListReply = browser.CrawlList(resource_url, cartoon.ConfigName) // 浏览器拉取列表数据
    var data []map[string]interface{}
    // 获取服务端返回的结果
    for _, v := range list_data.Data {
        data = append(data, map[string]interface{}{
            "cartoon_id": cartoon.Id,
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
    
    models.BatchInsert("cartoon_list", data, []string{"tags", "author", "detail", "resource_url", "resource_name", "resource_img_url"})
    if list_data.Next != "" && len(data) > 0 { // 是否下一页
        GetList(list_data.Next, cartoon)
    }

}

func GetChapter(id int64){
    var cartoon = models.GetCartoonById(id)
    var cartoonList = models.GetCartoonListByNo(cartoon.ResourceNo)

    for _, val := range cartoonList {
        
        fmt.Println("请求页面：", val.ResourceUrl)
        
        var chapter_data *Drive.ChapterReply = browser.CrawlChapter(val.ResourceUrl, cartoon.ConfigName) // 浏览器拉取列表数据
        var data []map[string]interface{}
        // 获取服务端返回的结果

        var cartoon_is_free string = "0" // 是否收费 0免费

        for _, v := range chapter_data.Data {

            if v.IsFree == "1" && cartoon_is_free == "0" {
                cartoon_is_free = "1"
            }

            data = append(data, map[string]interface{}{
                "resource_no": cartoon.ResourceNo,
                "unique_id": lib.MD5(val.UniqueId + cartoon.ResourceNo + v.ResourceName),
                "list_unique_id": val.UniqueId,
                "conent": "",
                "is_free": v.IsFree,
                "status": 0,
                "resource_url": v.ResourceUrl,
                "resource_name": v.ResourceName,
                "resource_img_url": v.ResourceImgUrl,
                "cdate": lib.Time(),
            })
        }
        // 修改状态信息
        models.UpdateCartoonListById(val.Id, map[string]interface{}{"is_free": cartoon_is_free, "is_end": chapter_data.Detail.IsEnd})
        models.BatchInsert("cartoon_chapter", data, []string{"is_free", "resource_url", "resource_name", "resource_img_url"})
        
    }
}

func GetChapterContent(id int64){
    var cartoon = models.GetCartoonById(id)
    var cartoonList = models.GetCartoonChapterListByNo(cartoon.ResourceNo)

    for _, val := range cartoonList {
        
        fmt.Println("请求页面：", val.ResourceUrl)
        
        var chapter_data *Drive.ChapterContentReply = browser.CrawlChapterContent(val.ResourceUrl, cartoon.ConfigName) // 浏览器拉取列表数据
        var data []map[string]interface{}
        // 获取服务端返回的结果
        // 清除现有数据
        models.DeleteChapterContentByChapterUniqueId(val.UniqueId)
        for _, v := range chapter_data.Data {
            data = append(data, map[string]interface{}{
                "resource_no": cartoon.ResourceNo,
                "list_unique_id": val.ListUniqueId,
                "chapter_unique_id": val.UniqueId,
                "resource_url": v.ResourceImgUrl,
                "cdate": lib.Time(),
            })
            // fmt.Println(v)
        }
        models.BatchInsert("cartoon_chapter_content", data, []string{})
        
    }
}