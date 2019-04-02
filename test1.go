package main
 
import (
    "spider/model"
    "spider/service"
    // "os/exec"
)

func main(){
    var models *model.Model = &model.Model{Db: model.InitDb()}
    var browser service.NodeBrowser = service.NodeBrowser{
        Service: service.Service{models},
    }

    browser.CreateBrowserClient() // 创建浏览器客户端
    var spider service.Spider = service.Spider{Models: models, Browser: browser}
    // spider.Book(4)
    spider.ChapterList(4)
}
