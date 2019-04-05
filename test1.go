package main
 
import (
    "spider/model"
    "spider/service"
)

func main(){
    
    var models *model.Model = &model.Model{Db: model.InitDb()}
    var browser service.NodeBrowser = service.NodeBrowser{
        Service: service.Service{models},
    }
    
    browser.CreateBrowserClient() // 创建浏览器客户端
   
    var resource service.Resource = service.Spider{Models: models, Browser: browser}
    // spider.ChapterList(4)
    resource.SpiderContentByResourceId(4)
}
