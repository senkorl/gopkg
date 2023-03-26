package hreg

import (
	"fmt"
	"regexp"
	"strings"
)

func GetCrawlUsername(url string) string {
	re := regexp.MustCompile("(http|https):\\/\\/[\\w\\-_]+(\\.[\\w\\-_]+)+([\\w\\-\\.,@?^=%&:/~\\+#]*[\\w\\-\\@?^=%&/~\\+#])?")
	result := re.FindAllStringSubmatch(url, -1)
	if len(result) == 0 {
		return ""
	}
	slice := Explode(url, "/")
	if len(slice) == 5 {
		return slice[len(slice)-2]
	}
	if len(slice) == 4 {
		return slice[len(slice)-1]
	}

	return ""
}

// Explode 字符串用符号隔开的数组
func Explode(str, sep string) []string {
	if len(str) == 0 {
		return []string{}
	}
	return strings.Split(str, sep)
}

var (
	regIdNumber = regexp.MustCompile(`([1-9])([0-7]\d{4})(19[0-9][0-9]|20[0-3][0-9])\d{4}(\d{3})(\d|X|x)`)
	regPhone    = regexp.MustCompile("(1\\d{2})(\\d{4})(\\d{4})")
)

const (
	replaceIdNumberStr = "$1$2********$4$5"
	replacePhoneStr    = "$1****$3"
	tempStr            = "#@$12$236+96"
)

func Check() {
	s := []string{"15659411565", "350122199607014615"}
	args := fmt.Sprintf("%v", s)
	//args = regPhone.ReplaceAllString(args, replacePhoneStr)
	//fmt.Printf("%v", args)
	args = regIdNumber.ReplaceAllString(args, replaceIdNumberStr)
	fmt.Printf("%v", args)
}
