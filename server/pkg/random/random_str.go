package random

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type StrType int

const (
	Number StrType = iota
	Letter
	NumberAndLetter
)

// CreateRandomStr 生成指定类型和长度的随机字符串 长度1-20
//
// 1 纯数字
//
// 2 纯字母
//
// 3 数字和字母混合
func CreateRandomStr(length int, strType ...StrType) string {
	codeType := Number
    if len(strType) > 0 {
        codeType = strType[0]
    }
	if length < 1 || length > 20 {
		length = 6
	}
	switch codeType {
	// 纯数字验证码
	case 1:
		format := "%0" + strconv.Itoa(length) + "v"
		return fmt.Sprintf(format, rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	// 纯字母验证码
	case 2:
		var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
		var str strings.Builder
		for i := 0; i < length; i++ {
			index := rand.Intn(len(letters))
			str.WriteRune(letters[index])
		}
		return str.String()
	// 数字+字母验证码
	case 3:
		var digits = []rune("0123456789")
		var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
		var sb strings.Builder
		for i := 0; i < length; i++ {
			// 随机选择是添加数字还是字母
			if rand.Intn(10) < 6 { // 让数字多一些，60%几率添加数字
				index := rand.Intn(len(digits))
				sb.WriteRune(digits[index])
			} else {
				index := rand.Intn(len(letters))
				sb.WriteRune(letters[index])
			}
		}
		return sb.String()
		// 纯数字验证码
	default:
		format := "%0" + strconv.Itoa(length) + "v"
		return fmt.Sprintf(format, rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(9223372036854775807)+1)
	}
}
