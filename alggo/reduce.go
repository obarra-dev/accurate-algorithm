package alggo

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func flattenLevelOne(arr []interface{}) string {
	var f []interface{}
	for _, e := range arr {
		switch v := e.(type) {
		case int:
			f = append(f, e)
		case []interface{}:
			for _, s := range v {
				f = append(f, s)
			}
		}
	}

	return formatFlatten(f)
}

func formatFlatten(f []interface{}) string {
	var sb strings.Builder
	sb.WriteString("[")
	for _, e := range f {
		switch v := e.(type) {
		case int:
			sb.WriteString(strconv.Itoa(v))
			sb.WriteString(",")
		case []interface{}:
			sb.WriteString(formatFlatten(v))
		}
	}
	sb.WriteString("]")

	return sb.String()
}

func mostRepeated(sentence string) string {
	n := strings.ToUpper(sentence)
	words := strings.Fields(n)

	counter := make(map[string]uint)
	for _, v := range words {
		counter[v]++
	}

	var mr = struct {
		word   string
		amount uint
	}{
		word:   "",
		amount: 0,
	}

	for k, v := range counter {
		if v > mr.amount {
			mr.word = k
			mr.amount = v
		}
	}

	return mr.word
}

func isPalindrome(str string) bool {
	length := len(str)
	for i := 0; i < length; i++ {
		if str[i] != str[length-i-1] {
			return false
		}
	}

	return true
}

func findFirstDuplicated(numbers []int) int {
	length := len(numbers)
	for i := 0; i < length; i++ {
		index := numbers[i]
		if index < 0 {
			index = -index
		}
		if numbers[index-1] > 0 {
			numbers[index-1] = -numbers[index-1]
		} else {
			return index
		}
	}

	return 0
}

func calculateSqrtSorted(number []float64) []float64 {
	fmt.Println("omar rules method")
	ss := make([]float64, len(number))
	j := len(number) - 1
	for i, v := range number {
		if v < 0 && math.Abs(v) > math.Abs(number[j]) {
			ss[j] = v * v
		} else {
			ss[i] = v * v
		}
		j--
	}

	return ss
}

func jumpingOnClouds(clouds []int) int {
	var jump int
	length := len(clouds)
	for i := 0; i < length; i++ {
		if i+2 < length && clouds[i+2] == 0 {
			jump++
			i++
		} else if i+1 < length && clouds[i+1] == 0 {
			jump++
		}
	}
	return jump
}

func salesByMatch(socks []int) int {
	storage := make([]int, 100)
	for _, v := range socks {
		storage[v]++
	}

	var amountPairs int
	for _, v := range storage {
		amountPairs += v / 2
	}

	return amountPairs
}

func detectNetwork(cardNumber string) string {
	var networkRules = []struct {
		network string
		iin     []string
		length  []int
	}{
		{"American Express", []string{"34", "37"}, []int{15}},
		{"Diners Club", []string{"38", "39"}, []int{14}},
		{"Visa", []string{"4"}, []int{13, 16, 19}},
		{"MasterCard", []string{"51", "52", "53", "54", "55"}, []int{16}},
	}

	validateIIN := func(iin []string, cardNumber string) bool {
		for _, v := range iin {
			if strings.HasPrefix(cardNumber, v) {
				return true
			}
		}
		return false
	}

	validateLength := func(length []int, cardNumber string) bool {
		for _, v := range length {
			if v == len(cardNumber) {
				return true
			}
		}
		return false
	}

	for _, v := range networkRules {
		if validateIIN(v.iin, cardNumber) && validateLength(v.length, cardNumber) {
			return v.network
		}
	}

	return "no match"
}
