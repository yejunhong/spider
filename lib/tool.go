package lib

import (
	"time"
	"crypto/md5"
	"encoding/hex"
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