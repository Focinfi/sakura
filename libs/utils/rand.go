package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var numberRunes = []rune("0123456789")
var charRunes = []rune("abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")

// RandNumberString retuns rand string of number in length of the given length
func RandNumberString(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = numberRunes[rand.Intn(len(numberRunes))]
	}
	return string(b)
}

// RandCharString retuns rand string of number in length of the given length
func RandCharString(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = charRunes[rand.Intn(len(charRunes))]
	}
	return string(b)
}

// RandNumberStringSlice generates codes for a ticket and
func RandNumberStringSlice(count uint, length uint) []string {
	codes := make([]string, count, count)
	for i := 0; i < int(count); i++ {
		codes[i] = RandNumberString(int(length))
	}

	return codes
}

// RandPhoneNumber random chinese phone number prefixed with 13
func RandPhoneNumber() string {
	return fmt.Sprintf("13%v", RandNumberString(9))
}
