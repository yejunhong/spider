package main
 
import (
    "spider/model"
    "spider/service"
)

func main(){
    
    var models *model.Model = &model.Model{
        Db: model.InitDb("202.43.91.26", "caiji", "caijishiwo7788dd", "caiji"),
        // Db61: model.InitDb("103.232.190.61", "xiaoshuo", "Manhua778899dd+-", "xiaoshuo"),
    }
    var browser service.NodeBrowser = service.NodeBrowser{
        Service: service.Service{models},
    }
    
    browser.CreateBrowserClient() // 创建浏览器客户端
   
    var resource service.Spider = service.Spider{Models: models, Browser: browser}
    // spider.ChapterList(4)
    resource.SpiderContentByResourceId(4)
}
