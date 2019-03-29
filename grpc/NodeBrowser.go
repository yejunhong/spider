package grpc
 
import (
    "google.golang.org/grpc"
    "fmt"
    "context"
    "sync"
    "time"
    "io"
    "spider/model"
)

type NodeBrowser struct{
    GrpcBrowserClient BrowserClient
}

/** 
 *
 * 创建grpc客户端双向流
 *
 */
func(browser *NodeBrowser) CreateBrowserClient(){
    // 创建一个grpc连接器
    var conn, err = grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())
    if err != nil{
        fmt.Println(err)
    }
    // 当请求完毕后记得关闭连接,否则大量连接会占用资源
    // defer conn.Close()
    browser.GrpcBrowserClient = NewBrowserClient(conn)
}

/**
 *
 * 爬取数据列表
 * @param url string 爬取的远程url
 * @param user_config_name string 数据源使用的配置
 *
 */
 func(browser *NodeBrowser) Book(request chan *Request) {
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
            fmt.Println(err)
            if err == io.EOF {
                fmt.Println(err)
                return
            }
            if err != nil {
                fmt.Println("Failed to receive a note : %v", err)
                return
            }
            model.RecordBook(data)
        }
        waitGroup.Done()
    }()
    go func(){
        for{
            select{ // 发送需要爬取的url，及配置
                case res := <-request:
                    err := stream.Send(res)
                    if err != nil {
                        fmt.Println(err)
                    }
                    fmt.Println("发生成功")
                default:
            }
        }
    }()
    waitGroup.Wait()
    fmt.Println("抓取结束")
}

/**
 *
 * 爬取数据章节
 * @param url string 爬取的远程url
 * @param user_config_name string 数据源使用的配置
 *
 */
 func(browser *NodeBrowser) Chapter(request chan *Request){
    // 客户端向grpc服务端发起请求
    ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	stream, err := browser.GrpcBrowserClient.Chapter(ctx)
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
                in, err := stream.Recv()
                if err == io.EOF {
                    // read done.
                    return
                }
                if err != nil {
                    fmt.Println("Failed to receive a note : %v", err)
                    return
                }
                
                fmt.Println(in)
            }
        waitGroup.Done()
    }()
    go func(){
        for{
            select{ // 发送需要爬取的url，及配置
                case res := <-request:
                    err := stream.Send(res)
                    if err != nil {
                        fmt.Println(err)
                    }
                    fmt.Println("发生成功")
                default:
            }
        }
    }()
    waitGroup.Wait()
    fmt.Println("抓取结束")
}

/**
 *
 * 爬取数据章节内容
 * @param url string 爬取的远程url
 * @param user_config_name string 数据源使用的配置
 *
 */
 func(browser *NodeBrowser) Content(request chan *Request){
    // 客户端向grpc服务端发起请求
    ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	stream, err := browser.GrpcBrowserClient.Chapter(ctx)
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
            in, err := stream.Recv()
            if err == io.EOF {
                // read done.
                return
            }
            if err != nil {
                fmt.Println("Failed to receive a note : %v", err)
                return
            }
            fmt.Println(in)
        }
        waitGroup.Done()
    }()
    go func(){
        for{
            select{ // 发送需要爬取的url，及配置
                case res := <-request:
                    err := stream.Send(res)
                    if err != nil {
                        fmt.Println(err)
                    }
                    fmt.Println("发生成功")
                default:
            }
        }
    }()
    waitGroup.Wait()
    fmt.Println("抓取结束")
}