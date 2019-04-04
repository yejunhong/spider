package service
 
import (
    "fmt"
    Drive "spider/grpc"
    "spider/model"
    "strconv"
    // "os/exec"
)

type Spider struct{
    Models *model.Model
    Browser NodeBrowser
}

/**
 *
 * 根据资源id
 * @param resourceId int64 资源Id
 *
 */
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
            request <- &Drive.Request{Url: "end", ConfigName: ""}
            fmt.Println("执行完成")
    }
}

/**
 *
 * 根据书籍Id 获取章节列表
 * @param resourceId int64 书籍Id
 *
 */
 func (spider *Spider) ChapterList(resourceId int64) {
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
                            Url: v.ResourceUrl, 
                            ConfigName: resource.ConfigName}
        }
        isend = true
    }()

    Loop:
        for {
            select{
                case <-spiderRequset.End:
                    if isend == true && len(spiderRequset.CartoonList) == 0 {
                        spiderEnd <- 1 // 中断程序
                        return
                    } 
                    fmt.Println("next url", <- next)
                case <-spiderEnd:
                    request <- &Drive.Request{Url: "end", ConfigName: "", Id: "0"}
                    break Loop // 中断循环
                default:
            }
        }
    fmt.Println("爬虫程序结束")
}

/**
 *
 * 根据书籍Id 获取章节列表
 * @param resourceId int64 书籍Id
 *
 */
 func (spider *Spider) ContentList(resourceId int64) {
    var request chan *Drive.Request = make(chan *Drive.Request, 5)
    var end chan int = make(chan int, 1)

    var resource model.CartoonResource = spider.Models.GetCartoonById(resourceId)
   
    var spiderRequset *SpiderRequset = &SpiderRequset{
        End: end,
        Request: request,
        CartoonResource: resource,
        CartoonChapter: make(map[string]model.CartoonChapter),
    }
    go spider.Browser.Content(spiderRequset)
    var next chan int = make(chan int, 5)
    var spiderEnd chan int = make(chan int, 1)
    var isend bool = false // 是否结束程序
    go func() { // 协程 发送爬虫信息
        var cartoonInfo = spider.Models.GetSqlCartoonListByNoStatus(resource.ResourceNo, 1)
        for _, info := range cartoonInfo {
            var cartoonChapter = spider.Models.GetChaptersFindByListUniqueId(info.UniqueId, 0)
            fmt.Println("编号：", info.ResourceNo, "-书籍数量：", len(cartoonChapter), "-名称：", info.ResourceName, "-", info.UniqueId)
            for _, v := range cartoonChapter {
                var IdStr string = strconv.FormatInt(v.Id,10)
                spiderRequset.CartoonChapter[IdStr] = v
            }
            for _, v := range cartoonChapter {
                next <- 1
                request <- &Drive.Request{
                        Id: strconv.FormatInt(v.Id,10), 
                        Url: v.ResourceUrl, 
                        ConfigName: resource.ConfigName}
            }
        }
        isend = true
    }()

    Loop:
        for {
            select{
                case <-spiderRequset.End:
                    <-next
                    if isend == true && len(spiderRequset.CartoonChapter) == 0 {
                        spiderEnd <- 1 // 中断程序
                        return
                    } 
                    // fmt.Println("next url", <- next)
                case <-spiderEnd:
                    request <- &Drive.Request{Url: "end", ConfigName: "", Id: "0"}
                    break Loop // 中断循环
                default:
            }
        }
    fmt.Println("爬虫程序结束")
}