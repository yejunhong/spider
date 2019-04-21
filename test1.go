package main
 
import (
    "spider/lib"
    "spider/model"
    "spider/service"
)

func main(){
    
    var config = lib.LoadConfig()
 
    var models *model.Model = &model.Model{
        Db: model.InitDb(config.Db_caiji.Host, config.Db_caiji.User, config.Db_caiji.Pass, config.Db_caiji.Name),
        Db61: model.InitDb(config.Db_xiaoshuo.Host, config.Db_xiaoshuo.User, config.Db_xiaoshuo.Pass, config.Db_xiaoshuo.Name),
        // DbManhua: model.InitDb(config.Db_manhua.Host, config.Db_manhua.User, config.Db_manhua.Pass, config.Db_manhua.Name),
    }
    var browser service.NodeBrowser = service.NodeBrowser{
        Service: service.Service{models},
    }
    
    browser.CreateBrowserClient() // 创建浏览器客户端
   
    var resource service.Spider = service.Spider{Models: models, Browser: browser}
    // resource.SpiderChapterByResourceId(5)
    resource.SpiderContentByResourceId(5)
}
