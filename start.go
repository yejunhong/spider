package main
 
import (
    "google.golang.org/grpc"
    "fmt"
	"context"
	Drive "spider/grpc"
)

func main(){
	// 创建一个grpc连接器
    conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
    if err != nil{
        fmt.Println(err)
    }
    // 当请求完毕后记得关闭连接,否则大量连接会占用资源
    defer conn.Close()
 
    // 创建grpc客户端
    c := Drive.NewGreeterClient(conn)
 
    name := "我是客户端,正在请求服务端!!!"
    // 客户端向grpc服务端发起请求
    result, err := c.Ping(context.Background(), &Drive.pingRequest{})
    fmt.Println(name)
    if err != nil{
        fmt.Println("请求失败!!!")
        return
    }
    // 获取服务端返回的结果
    fmt.Println(result.Message)
}