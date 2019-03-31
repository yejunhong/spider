package service
 
import (
    "fmt"
    Drive "spider/grpc"
    "spider/model"
    // "os/exec"
)

type Spider struct{
    Models *model.Model
    Browser NodeBrowser
}

func (spider *Spider) Book(resourceId int64){
    var request chan *Drive.Request = make(chan *Drive.Request, 1)
    var end chan int = make(chan int, 1)
    var resource model.CartoonResource = spider.Models.GetCartoonById(resourceId)
    var spiderRequset *SpiderRequset = &SpiderRequset{
        End: end,
        Request: request,
        CartoonResource: resource,
    }
    go spider.Browser.Book(spiderRequset)
    go func() {
        request <- &Drive.Request{Url: resource.ResourceUrl, ConfigName: resource.ConfigName}
    }()
    select{
        case <-spiderRequset.End:
            fmt.Println("执行完成")
    }
}

/**
 *
 * 根据书籍Id 获取章节列表
 * @param bookId int64 书籍Id
 *
 */
func (spider *Spider) Chapter(bookId int64) {
    var request chan *Drive.Request = make(chan *Drive.Request, 1)
    var end chan int = make(chan int, 1)
    var cartoonInfo = spider.Models.GetCartoonInfoById(bookId)
    var resource model.CartoonResource = spider.Models.GetCartoonByResourceNo(cartoonInfo.ResourceNo)
   
    var spiderRequset *SpiderRequset = &SpiderRequset{
        End: end,
        Request: request,
        CartoonResource: resource,
        CartoonList: cartoonInfo,
    }
    go spider.Browser.Chapter(spiderRequset)
    go func() {
        request <- &Drive.Request{Url: cartoonInfo.ResourceUrl, ConfigName: resource.ConfigName}
    }()
    select{
        case <-spiderRequset.End:
            fmt.Println("执行完成")
    }
}

/**
 *
 * 根据书籍Id 获取章节列表
 * @param bookId int64 书籍Id
 *
 */
 func (spider *Spider) ChapterList(resourceId int64) {
    var request chan *Drive.Request = make(chan *Drive.Request, 1)
    var end chan int = make(chan int, 1)

    var resource model.CartoonResource = spider.Models.GetCartoonById(resourceId)
    
   
    var spiderRequset *SpiderRequset = &SpiderRequset{
        End: end,
        Request: request,
        CartoonResource: resource,
    }
    go spider.Browser.Chapter(spiderRequset)
    var next chan int = make(chan int, 1)
    go func() {
        var cartoonInfo = spider.Models.GetCartoonListByNo(resource.ResourceNo)
        for _, v := range cartoonInfo {
            spiderRequset.CartoonList = v
            request <- &Drive.Request{Url: v.ResourceUrl, ConfigName: resource.ConfigName}
            next <- 1
        }
    }()

    for {
        select{
            case <-spiderRequset.End:
                n := <- next
                fmt.Println("执行完成1", n)
            default:
        }
    }
    
}

/**
 *
 * 根据书籍章节Id获取章节内容
 * @param chapterId int64 书籍章节
 *
 */
func (spider *Spider) Content(chapterId int64){
    var request chan *Drive.Request = make(chan *Drive.Request, 1)
    var end chan int = make(chan int, 1)
    // 关闭浏览器
    var chapterInfo = spider.Models.GetCartoonChapterInfoById(chapterId)
    var resource model.CartoonResource = spider.Models.GetCartoonByResourceNo(chapterInfo.ResourceNo)
    var spiderRequset *SpiderRequset = &SpiderRequset{
        End: end,
        Request: request,
        CartoonResource: resource,
        CartoonChapter: chapterInfo,
    }
    go spider.Browser.Content(spiderRequset)
    go func() {
        request <- &Drive.Request{Url: chapterInfo.ResourceUrl, ConfigName: resource.ConfigName}
    }()
    select{
        case <-spiderRequset.End:
            fmt.Println("执行完成")
    }
    
}
