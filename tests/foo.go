package tests

import (
	"crypto/sha1"
	b64 "encoding/base64"
	"fmt"
	"math"
	"strconv"

	"github.com/golang-module/carbon/v2"
	"github.com/spf13/cast"
)

func GetToken() {
	id := "XI8SCSqdrC55atSwjLlhM5A9FodQKU2J"
	time := "1632728136"
	time = cast.ToString(carbon.Now().Timestamp())
	sign := "b6f35f391626ebd517e826170effdb94deae2fa0afab7daafc22a03bcfe2fb46"
	h := sha1.New()
	h.Write([]byte(sign + time))
	bs := fmt.Sprintf("%x", h.Sum(nil))
	fmt.Println(bs)
	fmt.Println(b64.StdEncoding.EncodeToString([]byte(id + "." + time + "." + bs)))
	return
}

type User struct {
	Name string
	Age  uint32
}

func Quote() {
	x := 1
	p := &x         // p, of type *int, points to x
	fmt.Println(*p) // "1"
	fmt.Println(p)

	*p = 2         // equivalent to x = 2
	fmt.Println(x) // "2"

	user := &User{
		Name: "xiaoMing",
		Age:  18,
	}

	fmt.Println(user.Name)

	user1 := User{
		Name: "xiaoWANG",
		Age:  21,
	}
	fmt.Println(user1.Name)

}

type Celsius float64    // 摄氏温度
type Fahrenheit float64 // 华氏温度

const (
	AbsoluteZeroC Celsius = -273.15 // 绝对零度
	FreezingC     Celsius = 0       // 结冰点温度
	BoilingC      Celsius = 100     // 沸水温度
)

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

func Bar() {
	fmt.Printf("%g\n", BoilingC-FreezingC) // "100" °C
	boilingF := CToF(BoilingC)
	fmt.Printf("%g\n", boilingF-CToF(FreezingC)) // "180" °F
	//fmt.Printf("%g\n", boilingF-FreezingC)       // compile error: type mismatch

	var ce Celsius
	var fa Fahrenheit
	fmt.Println(ce == 0) // "true"
	fmt.Println(fa >= 0) // "true"
	fmt.Println(fa)
	fmt.Println(ce)
	fmt.Println(ce == Celsius(fa)) // "true"!
	//fmt.Println(c == f)          // compile error: type mismatch

	c := FToC(212.0)
	fmt.Println(c.String()) // "100°C"
	fmt.Printf("%v\n", c)   // "100°C"; no need to call String explicitly
	fmt.Printf("%s\n", c)   // "100°C"
	fmt.Println(c)          // "100°C"
	fmt.Printf("%g\n", c)   // "100"; does not call String
	fmt.Println(float64(c)) // "100"; does not call String
}

func FormatFloat(num float64, decimal int) (float64, error) {
	// 默认乘1
	d := float64(1)
	if decimal > 0 {
		// 10的N次方
		d = math.Pow10(decimal)
	}
	// math.trunc作用就是返回浮点数的整数部分
	// 再除回去，小数点后无效的0也就不存在了
	res := strconv.FormatFloat(math.Trunc(num*d)/d, 'f', -1, 64)
	return strconv.ParseFloat(res, 64)
}

func Foo() {
	var (
		a1, a2 []uint32
	)
	a1 = append(a1, uint32(4))
	a2 = append(a2, uint32(7))
	fmt.Printf("%v", a1)
	fmt.Printf("%v", a2)
}

func Bar11() {
	a, b := 100, 77
	fmt.Println(a % b)
}
