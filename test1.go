package main
 
import (
    "fmt"
    "spider/lib"
    "spider/model"
    "spider/service"
)

func main(){
    
    // lib.DonwloadFile("./1.png", "http://f2.kkmh.com/image/190305/0UX9Ax9Z0.webp-fe.w360.webp.m.i1")
    var config = lib.LoadConfig()
 
    var models *model.Model = &model.Model{
        Db: model.InitDb(config.Db_caiji.Host, config.Db_caiji.User, config.Db_caiji.Pass, config.Db_caiji.Name),
        // Db61: model.InitDb(config.Db_xiaoshuo.Host, config.Db_xiaoshuo.User, config.Db_xiaoshuo.Pass, config.Db_xiaoshuo.Name),
        // DbManhua: model.InitDb(config.Db_manhua.Host, config.Db_manhua.User, config.Db_manhua.Pass, config.Db_manhua.Name),
    }
    fmt.Println(1)
    var services = service.Service{models}
    /*
    var bookList = models.GetCartoonListByNoStatus("C001", 1)
    for _, v := range bookList {
        fmt.Println("书籍名称：", v.ResourceName)
        // services.DownloadBookIdImg(v.Id, v.ResourceImgUrl)
        // services.DownloadBookIdChaptersImg(v.UniqueId)
        services.DownloadBookIdContentImg(v.UniqueId)
        fmt.Println("下载书籍完毕。")
    }
    return*/
    
    var browser service.NodeBrowser = service.NodeBrowser{Service: services}
    
    browser.CreateBrowserClient() // 创建浏览器客户端
   
    var resource service.Spider = service.Spider{Models: models, Browser: browser}
    // resource.SpiderBookByResourceId(6)
    resource.SpiderChapterByResourceId(7)
    // resource.SpiderContentByResourceId(7)
}
