package main
 
import (
    "fmt"
    Drive "spider/grpc"
    "spider/model"
    "spider/service"
)

func main(){
    var models *model.Model = &model.Model{Db: model.InitDb()}
    var browser service.NodeBrowser = service.NodeBrowser{
        Service: service.Service{models},
    }

    browser.CreateBrowserClient() // 创建浏览器客户端
    
    var request chan *Drive.Request = make(chan *Drive.Request, 1)

    // var cartoon = models.GetCartoonById(1)
    // var cartoonList = models.GetCartoonListByNo(cartoon.ResourceNo)
    // var cartoonList = models.GetCartoonChapterListByNo(cartoon.ResourceNo)
    var resource model.CartoonResource = models.GetCartoonById(1)
    go browser.Book(&service.SpiderRequset{
        Request: request,
        CartoonResource: resource,
    })
    
    request <- &Drive.Request{Url: resource.ResourceUrl, ConfigName: resource.ConfigName}
    select{}
    fmt.Println("执行完成")
}
