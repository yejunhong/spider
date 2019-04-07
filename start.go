package main
 
import (
    "fmt"
    "spider/model"
    "spider/web/api"
    "os"
    "os/exec"
)
var models *model.Model
func main(){
    models = &model.Model{
        Db: model.InitDb("202.43.91.26", "caiji", "caijishiwo7788dd", "caiji"),
        Db61: model.InitDb("103.232.190.61", "xiaoshuo", "Manhua778899dd+-", "xiaoshuo"),
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