package main
 
import (
    "fmt"
    "net/http"
    "strings"
    "os"
    "bytes"
    "io"
    "io/ioutil"
)

func main(){
    var path = "./static/upload/book/2/1.txt"
    WriteFile(path, "http://img.ufo666.cn/upload/BookCoverH//20190222051500163.jpg")
    fmt.Println(1)
}