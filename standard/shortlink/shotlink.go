package main

import (
	"crypto/sha1"
	"encoding/binary"
	"fmt"
)

const (
	charset      = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	base         = 62
	maxShortCode = 62 * 62 * 62 * 62 * 62 * 62 // 62^6 = 56800235584
)

func main() {
	// 示例用法
	url := "https://www.example.com/very/long/url"
	id := uint64(123456)
	salt := "your-secret-salt"

	shortCode := GenerateShortCode(url, id, salt)
	fmt.Printf("Generated 6-digit short code: %s\n", shortCode)
}

// 生成6位短码
func GenerateShortCode(url string, id uint64, salt string) string {
	num := customHash(url, id, salt)
	return decimalToBase62(num)
}

// 自定义哈希算法
func customHash(url string, id uint64, salt string) uint64 {
	// 拼接数据
	data := []byte(fmt.Sprintf("%s%d%s", url, id, salt))

	// 使用SHA-1哈希
	hasher := sha1.New()
	hasher.Write(data)
	hashBytes := hasher.Sum(nil)

	// 取前6字节转换为uint64（实际使用前8字节中的6字节）
	var hashPart [8]byte
	copy(hashPart[2:], hashBytes[:6]) // 保留前2字节为0，避免溢出

	// 转换为0~62^6范围的数值
	num := binary.BigEndian.Uint64(hashPart[:])
	return num % maxShortCode
}

// 十进制转62进制（固定6位）
func decimalToBase62(n uint64) string {
	// 创建6位字符数组
	result := make([]byte, 6)

	// 从低位到高位计算
	for i := 5; i >= 0; i-- {
		remainder := n % base
		result[i] = charset[remainder]
		n = n / base
	}

	// 如果n未归零说明超过62^6，但经过mod运算不可能发生
	return string(result)
}
