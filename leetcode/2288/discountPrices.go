package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func discountPrices(sentence string, discount int) string {
	words := strings.Split(sentence, " ")
	for i, word := range words {
		if strings.Contains(word, "$") && isNumeric(word[1:]) {
			price, _ := strconv.Atoi(word[1:])
			discountedPrice := (1 - float64(discount)/100) * float64(price)
			words[i] = fmt.Sprintf("$%.2f", discountedPrice)
		}
	}
	return strings.Join(words, " ")
}

func isNumeric(s string) bool {
	matched, _ := regexp.MatchString("^[0-9]+$", s)
	return matched
}

func main() {
	sentence := "there are $1 $2 and 5$ candies in the shop"
	discount := 50
	s := discountPrices(sentence, discount)
	fmt.Println(s)
}
