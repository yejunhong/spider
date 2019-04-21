package main
 
import (
    "fmt"
    "spider/model"
    "spider/lib"
    "spider/web/api"
    "os"
    "os/exec"
)
var models *model.Model
func main(){

    var config = lib.LoadConfig()
 
    models = &model.Model{
        Db: model.InitDb(config.Db_caiji.Host, config.Db_caiji.User, config.Db_caiji.Pass, config.Db_caiji.Name),
        Db61: model.InitDb(config.Db_xiaoshuo.Host, config.Db_xiaoshuo.User, config.Db_xiaoshuo.Pass, config.Db_xiaoshuo.Name),
        DbManhua: model.InitDb(config.Db_manhua.Host, config.Db_manhua.User, config.Db_manhua.Pass, config.Db_manhua.Name),
    }
    go CommandNode()
    api.HttpRun(models, "4321")
}

// 阻塞直到该命令执行完成，该命令必须是被Start方法开始执行的
func CommandNode(){
    cmd := exec.Command("./spider.sh")  
    fmt.Println("创建Grpc 爬虫浏览器")
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    err := cmd.Run()
    
    if err != nil {
        fmt.Println("cmd.Output: ", err)
        return
    }
}