package utils

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"fmt"
	"math/rand"
	"strings"
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

// 62个字符, 需要6bit做索引(2 ^ 6 = 64)
var charTable = [...]rune{
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k',
	'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v',
	'w', 'x', 'y', 'z', '0', '1', '2', '3', '4', '5', '6',
	'7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H',
	'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S',
	'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
}

// ShortID returns a short id
func ShortID(url string) string {
	shortList := make([]string, 0, 4)
	sumData := md5.Sum([]byte(url))
	for i := 0; i < 4; i++ {
		part := sumData[i*4 : i*4+4]
		partUint := binary.BigEndian.Uint32(part)
		partUint &= 0x3fffffff
		shortIDBuf := &bytes.Buffer{}
		for j := 0; j < 6; j++ {
			index := partUint & 0x3d
			shortIDBuf.WriteRune(charTable[index])
			partUint = partUint >> 5
		}
		shortList = append(shortList, shortIDBuf.String()[:2])
	}
	return strings.Join(shortList, "")
}
