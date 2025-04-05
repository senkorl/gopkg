package main

import (
	"crypto/md5"
	"encoding/binary"
	"fmt"
	"hash/crc32"
	"strings"
)

// Base62 字符集
const base62Chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

// Base62 编码
func base62Encode(num uint64) string {
	var sb strings.Builder
	for num > 0 {
		remainder := num % 62
		sb.WriteByte(base62Chars[remainder])
		num /= 62
	}
	result := sb.String()
	// 反转字符串以获得正确顺序
	runes := []rune(result)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// 计算 CRC32
func crc32Hash(url string) uint32 {
	return crc32.ChecksumIEEE([]byte(url))
}

// 计算 MD5 并取前 8 字节
func md5Hash(url string) uint64 {
	hash := md5.Sum([]byte(url))
	return binary.BigEndian.Uint64(hash[:8]) // 取前 8 字节
}

// 生成短链
func generateShortURL(url string) string {
	hashValue := md5Hash(url)                        // 生成哈希
	crcValue := crc32Hash(url)                       // 计算 CRC32 校验码
	combined := (hashValue << 32) | uint64(crcValue) // 组合哈希和校验
	return base62Encode(combined)[:8]                // 取前 8 位
}

func main() {
	url := "https://example.com/very-long-url-paths"
	shortURL := generateShortURL(url)
	fmt.Println("Short URL:", shortURL)
}
