package grpc
 
import (
    "google.golang.org/grpc"
    "fmt"
    "context"
)

type NodeBrowser struct{
    GrpcBrowserClient BrowserClient
}

/* 双向流
// runRouteChat receives a sequence of route notes, while sending notes for various locations.
func runRouteChat(client pb.RouteGuideClient) {
	notes := []*pb.RouteNote{
		{Location: &pb.Point{Latitude: 0, Longitude: 1}, Message: "First message"},
		{Location: &pb.Point{Latitude: 0, Longitude: 2}, Message: "Second message"},
		{Location: &pb.Point{Latitude: 0, Longitude: 3}, Message: "Third message"},
		{Location: &pb.Point{Latitude: 0, Longitude: 1}, Message: "Fourth message"},
		{Location: &pb.Point{Latitude: 0, Longitude: 2}, Message: "Fifth message"},
		{Location: &pb.Point{Latitude: 0, Longitude: 3}, Message: "Sixth message"},
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.RouteChat(ctx)
	if err != nil {
		log.Fatalf("%v.RouteChat(_) = _, %v", client, err)
	}
	waitc := make(chan struct{})
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				// read done.
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("Failed to receive a note : %v", err)
			}
			log.Printf("Got message %s at point(%d, %d)", in.Message, in.Location.Latitude, in.Location.Longitude)
		}
	}()
	for _, note := range notes {
		if err := stream.Send(note); err != nil {
			log.Fatalf("Failed to send a note: %v", err)
		}
	}
	stream.CloseSend()
	<-waitc
}*/

/**
 *
 * 创建grpc客户端
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
 func(browser *NodeBrowser) CrawlList(url, user_config_name string) *ListReply{
    // 客户端向grpc服务端发起请求
    var result, err = browser.GrpcBrowserClient.CrawlList(context.Background(), &Request{
        Url: url,
        ConfigName: user_config_name,
    })
   
    if err != nil{
        panic(err)
    }
    return result
}

/**
 *
 * 爬取数据章节
 * @param url string 爬取的远程url
 * @param user_config_name string 数据源使用的配置
 *
 */
 func(browser *NodeBrowser) CrawlChapter(url, user_config_name string) *ChapterReply{
    // 客户端向grpc服务端发起请求
    var result, err = browser.GrpcBrowserClient.CrawlChapter(context.Background(), &Request{
        Url: url,
        ConfigName: user_config_name,
    })
   
    if err != nil{
        panic(err)
    }
    return result
}

/**
 *
 * 爬取数据章节内容
 * @param url string 爬取的远程url
 * @param user_config_name string 数据源使用的配置
 *
 */
 func(browser *NodeBrowser) CrawlChapterContent(url, user_config_name string) *ChapterContentReply{
    // 客户端向grpc服务端发起请求
    var result, err = browser.GrpcBrowserClient.CrawlChapterContent(context.Background(), &Request{
        Url: url,
        ConfigName: user_config_name,
    })
   
    if err != nil{
        panic(err)
    }
    return result
}