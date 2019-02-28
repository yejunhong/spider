package main
 
import (
    "fmt"
    "time"
)


func main(){
    
    var task = make(chan int, 10) // 最大任务数量
    var wait = make(chan int, 10) // 最大任务数量
    var wait = make(chan int, 10) // 最大任务数量

    go func(){
        for i := 0; i <= 100; i++ {
            wait <- i
            task <- i
        }
        end <- 1
    }()

    for{
        select{
            case <-task: {
                go func(){
                    time.Sleep(time.Second * 1)
                    <-wait // 告诉队列可以新增任务了
                }()
            }
            default:
        }
    }


}
