package main
import (
  "fmt"
  "net/http"
  "io/ioutil"
  "encoding/json"
)
func main() {
    var photos = []map[string]string{}
    photos = append(photos, map[string]string{
        "url": "222",
        "name": "333",
    })
    var user = map[string]interface{}{
        "thumbnail": "kongyixueyuan",
        "files": []map[string]string{},
    }
    user["photos"] = photos

    jsonStr, _ := json.Marshal(user)
    fmt.Println(string(jsonStr))
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