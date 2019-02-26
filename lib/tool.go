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