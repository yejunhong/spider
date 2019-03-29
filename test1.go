package main
 
import (
    "fmt"
    Drive "spider/grpc"
)

func main(){

    var browser Drive.NodeBrowser = Drive.NodeBrowser{}

    browser.CreateBrowserClient() // 创建浏览器客户端

    var request chan *Drive.Request = make(chan *Drive.Request, 1)
    go browser.Book(request)
    
    request <- &Drive.Request{Url: "https://www.kuaikanmanhua.com/tag/0", ConfigName: "kuaikanmanhua"}
    select{}
    fmt.Println("执行完成")
}
