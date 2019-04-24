package main
import (
  "fmt"
  "net/http"
  "io/ioutil"
  "encoding/json"
)
func main() {
    var url = "https://t1.hddhhn.com/uploads/tu/201512/277/98.jpg";
    var img = UploadImg(url)
    fmt.Println(img)
}

type Data struct {
    Img string
    Md5 string
    Img2 string
}
type ImgUpload struct {
    Code int64
    Message string
    Data Data
}
func UploadImg(imgUrl string) *ImgUpload {
    resp, err := http.Get("http://upload.manhua118.com/Img/Index/load?url=" + imgUrl)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }
    img := &ImgUpload{}
    var errs = json.Unmarshal(body, img)
    if errs != nil {
        panic(errs)
    }
    return img
}