package main
 
import (
    "fmt"
    "spider/model"
    "spider/lib"
    "sync"
	Drive "spider/grpc"
)

var browser Drive.NodeBrowser
var models model.Model

func main(){

    browser = Drive.NodeBrowser{}
    models = model.Model{Db: model.InitDb()}

    browser.CreateBrowserClient() // 创建浏览器客户端
    
    // 通过资源获取漫画列表
    /*var cartoon = models.GetCartoonById(2)
    GetList(cartoon.ResourceUrl, cartoon);*/

    // GetChapter(2)
    GetChapterContent(2)
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

type TaskInfo struct {
    Id int64
    UniqueId string
    ListUniqueId string
    ResourceUrl string
}
func GetChapter(id int64){

    var cartoon = models.GetCartoonById(id)

    var task = make(chan TaskInfo, 20) // 最大任务数量
    var wait = make(chan int, 10) // 等待执行
    var end = make(chan int, 1) // 任务投放结束

    go func(){
        var cartoonList = models.GetCartoonListByNo(cartoon.ResourceNo)
        for _, val := range cartoonList {
            wait <- 1 // 等待执行后进行下一个任务
            task <- TaskInfo{Id: val.Id, UniqueId: val.UniqueId, ResourceUrl: val.ResourceUrl}
        }
        end <- 1 // 任务投放结束
    }()

    var waitGroup sync.WaitGroup
    
    EndFor:
    for{
        select{
            case taskInfo := <-task: 

                waitGroup.Add(1)

                fmt.Println(taskInfo)

                go func(){
                    
                    defer waitGroup.Done() // 
                    
                    fmt.Println("请求页面：", taskInfo.ResourceUrl)

                    var chapter_data *Drive.ChapterReply = browser.CrawlChapter(taskInfo.ResourceUrl, cartoon.ConfigName) // 浏览器拉取列表数据
                    var data []map[string]interface{}
                    // 获取服务端返回的结果
                    var cartoon_is_free string = "0" // 是否收费 0免费

                    for _, v := range chapter_data.Data {
                        if v.IsFree == "1" && cartoon_is_free == "0" {
                            cartoon_is_free = "1"
                        }
                        data = append(data, map[string]interface{}{
                            "resource_no": cartoon.ResourceNo,
                            "unique_id": lib.MD5(taskInfo.UniqueId + cartoon.ResourceNo + v.ResourceName),
                            "list_unique_id": taskInfo.UniqueId,
                            "conent": "",
                            "is_free": v.IsFree,
                            "status": 0,
                            "resource_url": v.ResourceUrl,
                            "resource_name": v.ResourceName,
                            "resource_img_url": v.ResourceImgUrl,
                            "cdate": lib.Time(),
                        })
                    }
                    if chapter_data.Detail != nil {
                        models.UpdateCartoonListById(taskInfo.Id, map[string]interface{}{"is_free": cartoon_is_free, "is_end": chapter_data.Detail.IsEnd, "status": 1})
                    }
                    models.BatchInsert("cartoon_chapter", data, []string{"is_free", "resource_url", "resource_name", "resource_img_url"})
                    <-wait // 告诉队列可以新增任务了
                }()
            case <- end:
                fmt.Println("退出队列")
                break EndFor
            default:
        }
    }

    fmt.Println("等待执行完毕")
    waitGroup.Wait() //
    fmt.Println("执行完成")
}

func GetChapterContent(id int64){

    var cartoon = models.GetCartoonById(id)

    var task = make(chan TaskInfo, 5) // 最大任务数量
    var wait = make(chan int, 10) // 等待执行
    var end = make(chan int, 1) // 任务投放结束

    go func(){
        var cartoonList = models.GetCartoonChapterListByNo(cartoon.ResourceNo)
        for _, val := range cartoonList {
            wait <- 1 // 等待执行后进行下一个任务
            task <- TaskInfo{Id: val.Id, UniqueId: val.UniqueId, ListUniqueId: val.ListUniqueId, ResourceUrl: val.ResourceUrl}
        }
        end <- 1 // 任务投放结束
    }()

    var waitGroup sync.WaitGroup
    
    EndFor:
    for{
        select{
            case taskInfo := <-task: 

                waitGroup.Add(1)

                fmt.Println(taskInfo)
                fmt.Println("请求页面：", taskInfo.ResourceUrl)
                go func(){
                    defer waitGroup.Done() // 

                    var chapter_data *Drive.ChapterContentReply = browser.CrawlChapterContent(taskInfo.ResourceUrl, cartoon.ConfigName) // 浏览器拉取列表数据
                    var data []map[string]interface{}
                    // 获取服务端返回的结果
                    // 清除现有数据
                    models.DeleteChapterContentByChapterUniqueId(taskInfo.UniqueId)

                    for _, v := range chapter_data.Data {
                        data = append(data, map[string]interface{}{
                            "resource_no": cartoon.ResourceNo,
                            "list_unique_id": taskInfo.ListUniqueId,
                            "chapter_unique_id": taskInfo.UniqueId,
                            "resource_url": v.ResourceImgUrl,
                            "cdate": lib.Time(),
                        })
                    }
                    models.UpdateCartoonChapterById(taskInfo.Id, map[string]interface{}{"status": 1})
                    models.BatchInsert("cartoon_chapter_content", data, []string{})
                    <-wait // 告诉队列可以新增任务了
                }()
            case <- end:
                fmt.Println("退出队列")
                break EndFor
            default:
        }
    }

    fmt.Println("等待执行完毕")
    waitGroup.Wait() //
    fmt.Println("执行完成")

}