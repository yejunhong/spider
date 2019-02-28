package main
 
import (
    "fmt"
    "sync"
    "time"
    "math/rand"
)

func main(){
    /*
    var waitGroup sync.WaitGroup
    waitGroup.Add(1)
    for i := 0; i <= 10; i++ {
        go func(){
            defer waitGroup.Done()
            fmt.Println(i)
        }()
    }
    waitGroup.Wait()*/

    var waitNotice = make(chan int, 10) // 等等下一次任务
    var task = make(chan int) // 任务

    // var maxTask = 10 // 最大任务数量
    // var currentTask = 0 // 当前执行数量

    var task_list = make(chan int, 101)
    var end = make(chan int, 1)

    var waitGroup sync.WaitGroup

    go func(){ // 任务

        for i := 0; i <= 100; i++ { // 任务列表
            task_list <- i
        }

        for i := 0; i <= 10; i++ { // 执行10个任务
            b := <-task_list
            task <- b
        }

        for {
            select {
                case <- waitNotice:
                    if len(task_list) > 0 {
                        b := <-task_list
                        task <- b
                    } else {
                        end <- 1
                    }
                        
            }
        }
    }()
    
    

    TastFor:
        for{
            select{
                case t := <- task:
                    waitGroup.Add(1)
                    go func(){
                        defer waitGroup.Done()
                        time.Sleep(time.Second * time.Duration(rand.Intn(10)))
                        fmt.Println(t)
                        waitNotice <- 1 // 告诉队列 我执行完了 可以执行下一个任务
                    }()
                case <- end:
                    fmt.Println("退出队列")
                    break TastFor
                default: // 检测chan 是否已经满
            }
        }
    fmt.Println("等待执行完毕")
    waitGroup.Wait()
    fmt.Println("执行完成")
}
