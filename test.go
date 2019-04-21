package main
import (
  "fmt"
  "net/http"
  "io/ioutil"
)
func main() {
    var url = "http://store.cqhdx.com/img1/v_112178_7feb38";
    UploadImg(url)
}

func UploadImg(imgUrl string) {
    resp, err :=   http.Get("http://upload.manhua008.com/Img/Index/load?url=" + imgUrl)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }
    fmt.Println(string(body))
}