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

    models = &model.Model{Db: model.InitDb()}
    // go CommandNode()
    api.HttpRun(models, "4321")
    // fmt.Println(models.GetCartoonResources("酷", 1, 0))
}

func CommandNode(){
    cmd := exec.Command("node", "./node/server.ts")  
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    err := cmd.Run()
    if err != nil {
        fmt.Println("cmd.Output: ", err)
        return
    }
    //阻塞直到该命令执行完成，该命令必须是被Start方法开始执行的
}