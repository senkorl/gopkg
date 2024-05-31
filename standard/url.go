package standard

import (
	"fmt"
	"net/url"
)

func ValidateUrl() {
	str := "https://lme-test-public.lmeapp.com/251698719460201pEEJ.png"
	_, err := url.ParseRequestURI(str)
	if err != nil {
		fmt.Println(1)
		return
	} else {
		fmt.Println(2)
	}
	return
}
