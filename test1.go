package main
 
import (
    "fmt"
    Drive "spider/grpc"
    "spider/model"
    "spider/service"
)

func main(){
    var browser service.NodeBrowser = service.NodeBrowser{
        Service: service.Service{&model.Model{Db: model.InitDb()}},
    }

    browser.CreateBrowserClient() // 创建浏览器客户端
    
    var request chan *Drive.Request = make(chan *Drive.Request, 1)
    go browser.Book(&service.SpiderRequset{Request: request})
    
    request <- &Drive.Request{Url: "https://www.kuaikanmanhua.com/tag/0", ConfigName: "kuaikanmanhua"}
    select{}
    fmt.Println("执行完成")
}
