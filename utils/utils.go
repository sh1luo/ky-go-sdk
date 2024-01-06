package utils

import (
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// 取运行目录
func C程序_取运行目录() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	return res

}
func W文本_取随机字符串(字符串长度 int) string {
	var strByte = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	var strByteLen = len(strByte)
	bytes := make([]byte, 字符串长度)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 字符串长度; i++ {
		bytes[i] = strByte[r.Intn(strByteLen-1)]
	}
	if bytes[0] == strByte[strByteLen-1] { //第一位不能是0 防止意外
		bytes[0] = strByte[strByteLen-2]
	}

	return string(bytes)
}
func W文本_取随机字符串_数字(字符串长度 int) string {
	var strByte = []byte("1234567890")
	var strByteLen = len(strByte)
	bytes := make([]byte, 字符串长度)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 字符串长度; i++ {
		bytes[i] = strByte[r.Intn(strByteLen-1)]
	}
	if bytes[0] == strByte[strByteLen-1] { //第一位不能是0 防止意外
		bytes[0] = strByte[strByteLen-2]
	}

	return string(bytes)
}

// 文本取出中间文本
func W文本_取出中间文本(内容 string, 左边文本 string, 右边文本 string) string {
	左边位置 := strings.Index(内容, 左边文本)
	if 左边位置 == -1 {
		return ""
	}
	左边位置 = 左边位置 + len(左边文本)
	内容 = string([]byte(内容)[左边位置:])

	var 右边位置 int
	if 右边文本 == "" {
		右边位置 = len(内容)
	} else {
		右边位置 = strings.Index(内容, 右边文本)
		if 右边位置 == -1 {
			return ""
		}
	}
	内容 = string([]byte(内容)[:右边位置])
	return 内容
}
