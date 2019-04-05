package service
 
import (
    "fmt"
    Drive "spider/grpc"
    "spider/model"
    "strconv"
    // "os/exec"
    // "time"
)

type Spider struct{
    Models *model.Model
    Browser NodeBrowser
}

// 资源 爬虫
type Resource interface { 
    SpiderBookByResourceId(resourceId int64) // 根据资源爬取书籍列表
    SpiderChapterByResourceId(resourceId int64) // 根据资源爬取书籍章节列表
    SpiderContentByResourceId(resourceId int64) // 根据资源爬取书籍章节内容
}

// 书籍 爬虫
type Book interface { 
    SpiderChapterByBookId(bookId int64) // 根据书籍Id爬取章节列表
    SpiderContentByBookId(bookId int64) // 根据书籍Id爬取章节内容
}

/**
 *
 * 根据资源id 爬虫书籍列表
 * @param resourceId int64 资源Id
 *
 */
func (spider *Spider) SpiderBookResourceId(resourceId int64){
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
    <-spiderRequset.End
    request <- &Drive.Request{Url: "end", ConfigName: "", Id: "0"}
    fmt.Println("执行完成")
}

/**
 *
 * 根据资源Id 获取章节列表
 * @param resourceId int64 资源Id
 *
 */
 func (spider *Spider) SpiderChapterByResourceId(resourceId int64) {
    var request chan *Drive.Request = make(chan *Drive.Request, 5)
    var end chan int = make(chan int, 1)

    var resource model.CartoonResource = spider.Models.GetCartoonById(resourceId)
    var spiderRequset *SpiderRequset = &SpiderRequset{
        End: end,
        Request: request,
        CartoonResource: resource,
        CartoonList: make(map[string]model.CartoonList),
    }
    go spider.Browser.Chapter(spiderRequset)
    var next chan int = make(chan int, 5)
    var spiderEnd chan int = make(chan int, 1)
    var isend bool = false // 是否结束程序
    go func() { // 协程 发送爬虫信息
        var cartoonInfo = spider.Models.GetCartoonListByNoStatus(resource.ResourceNo, 0)
        for _, v := range cartoonInfo {
            var IdStr string = strconv.FormatInt(v.Id,10)
            spiderRequset.CartoonList[IdStr] = v
        }
        for _, v := range cartoonInfo {
            next <- 1
            request <- &Drive.Request{
                            Id: strconv.FormatInt(v.Id,10), 
                            Url: v.ResourceUrl, ConfigName: resource.ConfigName}
        }
        isend = true
    }()
    for {
        select{
            case <-spiderRequset.End:
                if isend == true && len(spiderRequset.CartoonList) == 0 {
                    spiderEnd <- 1 // 中断程序
                } 
            case <-spiderEnd:
                request <- &Drive.Request{Url: "end", ConfigName: "", Id: "0"}
                goto Loop // 中断循环
        }
    }
    Loop:
        fmt.Println("爬虫程序结束")
}

/**
 *
 * 根据书籍Id 获取章节列表
 * @param resourceId int64 书籍Id
 *
 */
 func (spider *Spider) SpiderContentByResourceId(resourceId int64) {
    var request chan *Drive.Request = make(chan *Drive.Request, 3)
    var end chan int = make(chan int, 1)
    var resource model.CartoonResource = spider.Models.GetCartoonById(resourceId)
    var spiderRequset *SpiderRequset = &SpiderRequset{
        End: end,
        Request: request,
        CartoonResource: resource,
        CartoonChapter: make(map[string]model.CartoonChapter),
    }
    go spider.Browser.Content(spiderRequset)
    var next chan int = make(chan int, 3)
    var spiderEnd chan int = make(chan int, 1)
    var isend bool = false // 是否结束程序
    go func() { // 协程 发送爬虫信息
        var cartoonInfo = spider.Models.GetSqlCartoonListByNoStatus(resource.ResourceNo, 1)
        for _, info := range cartoonInfo {
            var cartoonChapter = spider.Models.GetChaptersFindByListUniqueId(info.UniqueId, 0)
            fmt.Println("编号：", info.ResourceNo, "-名称：", info.ResourceName, "-", info.UniqueId, "-书籍章节：", len(cartoonChapter))
            for _, v := range cartoonChapter {
                var IdStr string = strconv.FormatInt(v.Id,10)
                spiderRequset.CartoonChapter[IdStr] = v
            }
            for _, v := range cartoonChapter {
                next <- 1
                request <- &Drive.Request{
                        Id: strconv.FormatInt(v.Id,10), 
                        Url: v.ResourceUrl, ConfigName: resource.ConfigName}
            }
        }
        isend = true
    }()
    for {
        select{
            case <-spiderRequset.End:
                <-next
                if isend == true && len(spiderRequset.CartoonChapter) == 0 {
                    spiderEnd <- 1 // 中断程序
                } 
            case <-spiderEnd:
                request <- &Drive.Request{Url: "end", ConfigName: "", Id: "0"}
                goto Loop // 退出循环
        }
    }
    Loop:
        fmt.Println("爬虫程序结束")
}

/**
 *
 * 根据书籍Id 爬取章节列表
 * @param BookId int64 书籍Id
 *
 */
 func (spider *Spider) SpiderChapterByBookId(BookId int64) {
    var request chan *Drive.Request = make(chan *Drive.Request, 5)
    var end chan int = make(chan int, 1)
    var bookInfo = spider.Models.GetCartoonInfoById(BookId)
    var resource model.CartoonResource = spider.Models.GetCartoonByResourceNo(bookInfo.ResourceNo)
    var spiderRequset *SpiderRequset = &SpiderRequset{
        End: end,
        Request: request,
        CartoonResource: resource,
        CartoonList: make(map[string]model.CartoonList),
    }
    go spider.Browser.Chapter(spiderRequset)
    go func() { // 协程 发送爬虫信息
        var IdStr string = strconv.FormatInt(bookInfo.Id,10)
        spiderRequset.CartoonList[IdStr] = bookInfo
        request <- &Drive.Request{Id: IdStr, Url: bookInfo.ResourceUrl, ConfigName: resource.ConfigName}
    }()
    <-spiderRequset.End
    request <- &Drive.Request{Url: "end", ConfigName: "", Id: "0"} // 通知grpc 断开连接
    fmt.Println("爬虫程序结束")
}

/**
 *
 * 根据书籍Id 获取章节内容
 * @param BookId int64 书籍Id
 *
 */
 func (spider *Spider) SpiderContentByBookId(BookId int64) {
    var request chan *Drive.Request = make(chan *Drive.Request, 3)
    var end chan int = make(chan int, 1)
    var bookInfo = spider.Models.GetCartoonInfoById(BookId)
    var resource model.CartoonResource = spider.Models.GetCartoonByResourceNo(bookInfo.ResourceNo)
    var spiderRequset *SpiderRequset = &SpiderRequset{
        End: end,
        Request: request,
        CartoonResource: resource,
        CartoonChapter: make(map[string]model.CartoonChapter),
    }
    go spider.Browser.Content(spiderRequset)
    var next chan int = make(chan int, 3)
    var spiderEnd chan int = make(chan int, 1)
    go func() { // 协程 发送爬虫信息
        var cartoonChapter = spider.Models.GetChaptersFindByListUniqueId(bookInfo.UniqueId, 0)
        fmt.Println("编号：", bookInfo.ResourceNo, "-名称：", bookInfo.ResourceName, "-", bookInfo.UniqueId, "-书籍章节：", len(cartoonChapter))
        for _, v := range cartoonChapter {
            var IdStr string = strconv.FormatInt(v.Id,10)
            spiderRequset.CartoonChapter[IdStr] = v
        }
        for _, v := range cartoonChapter {
            next <- 1
            request <- &Drive.Request{
                    Id: strconv.FormatInt(v.Id,10), 
                    Url: v.ResourceUrl, ConfigName: resource.ConfigName}
        }
    }()
    for {
        select{
            case <-spiderRequset.End:
                <-next
                if len(spiderRequset.CartoonChapter) == 0 {
                    spiderEnd <- 1 // 中断程序
                } 
            case <-spiderEnd:
                request <- &Drive.Request{Url: "end", ConfigName: "", Id: "0"}
                goto Loop // 退出循环
        }
    }
    Loop:
        fmt.Println("爬虫程序结束")
}