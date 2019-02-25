package main
 
import (
    "fmt"
    "spider/model"
	Drive "spider/grpc"
)

func main(){

    var browser Drive.NodeBrowser = Drive.NodeBrowser{}
    var models model.Model = model.Model{Db: model.InitDb()}
    
    browser.CreateBrowserClient()

    var cartoon = models.GetCartoonById(1)

    var list_data *Drive.ListReply = browser.CrawlList(cartoon.ResourceUrl, cartoon.ConfigName) // 浏览器拉取列表数据
  
    // 获取服务端返回的结果
    for _, v := range list_data.Data {
        cartoon_list
        fmt.Println(v.ResourceUrl)
    }
    
}