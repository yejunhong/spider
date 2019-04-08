package lib

import (
	"time"
	"crypto/md5"
	"encoding/hex"
	"net/http"
    "strings"
    "os"
    "bytes"
    "io"
    "io/ioutil"
)

/**
 *
 * 返回当前时间戳
 * @return int64
 *
 */
func Time() int64{
	return time.Now().Unix()
}

/**
 *
 * 返回当前时间戳 转换后的日期格式
 * @param t int64 时间戳
 * @return string
 *
 */
func DateTime(t int64) string {
	//时间戳转日期
	//设置时间戳 使用模板格式化为日期字符串
	return time.Unix(t, 0).Format("2006-01-02 15:04:05") 
}

/**
 *
 * md5加密
 * @param encrypt string 需要加密的字符串
 * @return "加密后的字符串"
 *
 */
func MD5(encrypt string) string{
	h := md5.New() 
    h.Write([]byte(encrypt)) // 需要加密的字符串为 123456 
    return hex.EncodeToString(h.Sum(nil) ) // 输出加密结果 
}

/**
 *
 * 根据路径 创建|写入文件内容
 * @param path string 存放路径
 * @param text string 需要写入的内容
 * @return n int, err error
 *
 */
func WriteFile(path string, text string) (n int, err error) {
    var isExist bool = PathExists(path)
    if isExist == true {
        return
    } else {
        var paths []string = strings.Split(path, "/")
        var name string = strings.Join(paths[:len(paths) - 1], "/")
        os.MkdirAll(name, os.ModePerm)
    }
	out, _ := os.Create(path)
    defer out.Close()
    n, err = out.Write([]byte(text))
    return 
}

/**
 *
 * 根据远程路径下载文件
 * @param path string 存放路径
 * @param url string 远程文件地址
 * @return n int64, err error
 *
 */
func DonwloadFile(path string, url string) (n int64, err error) {
    var isExist bool = PathExists(path)
    if isExist == true {
        return
    } else {
        var paths []string = strings.Split(path, "/")
        var name string = strings.Join(paths[:len(paths) - 1], "/")
        os.MkdirAll(name, os.ModePerm)
    }
	out, err := os.Create(path)
	defer out.Close()
	resp, err := http.Get(url)
	defer resp.Body.Close()
	pix, err := ioutil.ReadAll(resp.Body)
	n, err = io.Copy(out, bytes.NewReader(pix))
	return
}

/**
 *
 * 根据文件路径 检测是否存在
 * @param path string 路径
 * @return bool
 *
 */
func PathExists(path string) bool {
    _, err := os.Stat(path)
    if err == nil {
        return true
    }
    if os.IsExist(err) {
        return true
    }
    return false
}