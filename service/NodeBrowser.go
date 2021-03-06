package service
 
import (
    "google.golang.org/grpc"
    "fmt"
    "context"
    "sync"
    "io"
    Drive "spider/grpc"
    "spider/model"
)

type NodeBrowser struct{
    Service Service
    GrpcBrowserClient Drive.BrowserClient
}

type SpiderRequset struct {
    End chan int // 程序结束
    Request chan *Drive.Request
    CartoonResource model.CartoonResource
    CartoonList map[string]model.CartoonList
    CartoonChapter map[string]model.CartoonChapter
}

/** 
 *
 * 创建grpc客户端双向流
 *
 */
func(browser *NodeBrowser) CreateBrowserClient() *grpc.ClientConn{
    // 创建一个grpc连接器
    var conn *grpc.ClientConn
    var err error
    conn, err = grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())
    if err != nil{
        fmt.Println(err)
    }
    // 当请求完毕后记得关闭连接,否则大量连接会占用资源
    // defer conn.Close()
    browser.GrpcBrowserClient = Drive.NewBrowserClient(conn)
    return conn
}

/**
 *
 * 爬取数据列表
 * @param url string 爬取的远程url
 * @param user_config_name string 数据源使用的配置
 *
 */
func(browser *NodeBrowser) Book(spiderRequset *SpiderRequset) {
    // 客户端向grpc服务端发起请求
    // ctx, cancel := context.WithTimeout(context.Background(), 20 * time.Second)
    ctx, cancel := context.WithCancel(context.Background())
	stream, err := browser.GrpcBrowserClient.Book(ctx)
	if err != nil {
        panic(err)
    }
    
    defer func(){ // 函数结束后执行
        stream.CloseSend()
        cancel()
    }()

    var waitGroup sync.WaitGroup
	waitGroup.Add(1) // 增加计算器
	go func() {
        for {
            data, err := stream.Recv()
            if err == io.EOF {
                fmt.Println("关闭-爬虫书籍列表-Grpc双向流连接", err)
                waitGroup.Done()
                return
            }
            if err != nil {
                fmt.Println("Failed to receive a note :", err)
                return
            }
            browser.Service.RecordBook(data, spiderRequset.CartoonResource)
            if data.Next == false {
                spiderRequset.End <- 1
            }
        }
    }()
    go func(){
        for{
            select{ // 发送需要爬取的url，及配置
                case res := <-spiderRequset.Request:
                    err := stream.Send(res)
                    
                    if err != nil {
                        fmt.Println(err)
                    }
                    fmt.Println("发生成功")
            }
        }
    }()
    waitGroup.Wait()
}

/**
 *
 * 爬取数据章节
 * @param url string 爬取的远程url
 * @param user_config_name string 数据源使用的配置
 *
 */
func(browser *NodeBrowser) Chapter(spiderRequset *SpiderRequset){
    // 客户端向grpc服务端发起请求
    ctx, cancel := context.WithCancel(context.Background())
	stream, err := browser.GrpcBrowserClient.Chapter(ctx)
	if err != nil {
        panic(err)
    }
    
    defer func(){ // 函数结束后执行
        if r := recover(); r!= nil {
            fmt.Println("Recover from r : ",r)
        } else {
            stream.CloseSend()
            cancel()
        }
    }()

    var waitGroup sync.WaitGroup
	waitGroup.Add(1) // 增加计算器
	go func() {
        for {
            data, err := stream.Recv()
            if err == io.EOF {
                fmt.Println("关闭-爬虫章节-Grpc双向流连接", err)
                waitGroup.Done()
                return
            }
            if err != nil {
                fmt.Println("Failed to receive a note ", err)
                return
            }
            if val, ok := spiderRequset.CartoonList[data.Id]; ok {
                go browser.Service.RecordChapter(data, spiderRequset.CartoonResource, val)
                if data.Next == false { 
                    delete(spiderRequset.CartoonChapter, data.Id)
                }
            }
            
            if data.Next == false {
                spiderRequset.End <- 1
            }
        }
    }()
    go func(){
        for{
            select{ // 发送需要爬取的url，及配置
                case res := <-spiderRequset.Request:
                    // fmt.Println("抓取章节-编号：", res.Id, "-书籍Url：", res.Url)
                    err := stream.Send(res)
                    if err != nil {
                        fmt.Println(err)
                    }
            }
        }
    }()
    waitGroup.Wait()
    fmt.Println("退出爬虫")
}

/**
 *
 * 爬取数据章节内容
 * @param url string 爬取的远程url
 * @param user_config_name string 数据源使用的配置
 *
 */
func(browser *NodeBrowser) Content(spiderRequset *SpiderRequset){

    // 客户端向grpc服务端发起请求
    ctx, cancel := context.WithCancel(context.Background())
	stream, err := browser.GrpcBrowserClient.Content(ctx)
	if err != nil {
        panic(err)
    }
    
    defer func(){ // 函数结束后执行
        if r := recover(); r!= nil {
            fmt.Println("Recover from r : ",r)
        } else {
            stream.CloseSend()
            cancel()
        }
    }()

    var waitGroup sync.WaitGroup
	waitGroup.Add(1) // 增加计算器
	go func() {
        for {
            data, err := stream.Recv()
            if err == io.EOF {
                fmt.Println("关闭-爬虫内容-Grpc双向流连接", err)
                waitGroup.Done()
                return
            }
            if err != nil {
                fmt.Println("Failed to receive a note", err)
                return
            }
            // fmt.Println(data.Id)
            if val, ok := spiderRequset.CartoonChapter[data.Id]; ok {
                browser.Service.RecordContent(data, spiderRequset.CartoonResource, val)
                delete(spiderRequset.CartoonChapter, data.Id)
            }
            if data.Next == false {
                spiderRequset.End <- 1
            }
        }
    }()
    go func(){
        for{
            select{ // 发送需要爬取的url，及配置
                case res := <-spiderRequset.Request:
                    // fmt.Println("抓取内容-编号：", res.Id, "-书籍Url：", res.Url)
                    err := stream.Send(res)
                    if err != nil {
                        fmt.Println(err)
                    }
                    // fmt.Println("发生成功")
            }
        }
    }()
    waitGroup.Wait()
}