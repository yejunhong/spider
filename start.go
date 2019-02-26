package main
 
import (
    "fmt"
    "spider/model"
    "spider/lib"
	Drive "spider/grpc"
)

var browser Drive.NodeBrowser
var models model.Model

func main(){

    browser = Drive.NodeBrowser{}
    models = model.Model{Db: model.InitDb()}

    browser.CreateBrowserClient() // 创建浏览器客户端
    GetList();
    // GetChapter();
}


func GetList(){
    var cartoon = models.GetCartoonById(1)
    var list_data *Drive.ListReply = browser.CrawlList(cartoon.ResourceUrl, cartoon.ConfigName) // 浏览器拉取列表数据

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
        // fmt.Println(v)
    }
    
    models.BatchInsert("cartoon_list", data, []string{"tags", "author", "detail", "resource_url", "resource_name", "resource_img_url"})

}

func GetChapter(){
    var cartoon = models.GetCartoonById(1)
    var cartoonList = models.GetCartoonInfoById(6760)
    
    var chapter_data *Drive.ChapterReply = browser.CrawlChapter(cartoonList.ResourceUrl, cartoon.ConfigName) // 浏览器拉取列表数据

    var data []map[string]interface{}
    // 获取服务端返回的结果

    for _, v := range chapter_data.Data {
        data = append(data, map[string]interface{}{
            "resource_no": cartoon.ResourceNo,
            "unique_id": lib.MD5(cartoon.ResourceNo + v.ResourceName),
            "list_unique_id": cartoon.UniqueId,
            "conent": "",
            "is_free": v.IsFree,
            "status": 0,
            "resource_url": v.ResourceUrl,
            "resource_name": v.ResourceName,
            "resource_img_url": v.ResourceImgUrl,
            "cdate": lib.Time(),
        })
        // fmt.Println(v)
    }
    fmt.Println(data)
    return
    models.BatchInsert("cartoon_chapter", data, []string{"is_free", "resource_url", "resource_name", "resource_img_url"})

}