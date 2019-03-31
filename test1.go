package main
 
import (
    "fmt"
    Drive "spider/grpc"
    "spider/model"
    "spider/service"
    "os/exec"
)

func main(){
    var models *model.Model = &model.Model{Db: model.InitDb()}
    var browser service.NodeBrowser = service.NodeBrowser{
        Service: service.Service{models},
    }

    browser.CreateBrowserClient() // 创建浏览器客户端
    
    var request chan *Drive.Request = make(chan *Drive.Request, 1)
    var end chan int = make(chan int, 1)
    // pkill Chromium
    // 关闭浏览器
    // cmd := exec.Command("pkill", "Chromium")
    // cmd.Run()
    // var cartoon = models.GetCartoonById(1)
    // var cartoonList = models.GetCartoonListByNo(cartoon.ResourceNo)
    // var cartoonList = models.GetCartoonChapterListByNo(cartoon.ResourceNo)
    var resource model.CartoonResource = models.GetCartoonById(1)
    var spiderRequset *service.SpiderRequset = &service.SpiderRequset{
        End: end,
        Request: request,
        CartoonResource: resource,
    }
    go browser.Book(spiderRequset)
    
    go func() {
        request <- &Drive.Request{Url: resource.ResourceUrl, ConfigName: resource.ConfigName}
    }()
    
    select{
        case <-spiderRequset.End:
            fmt.Println("执行完成")
    }
    
}
