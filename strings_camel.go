package zconv

import (
	"strings"
	"unicode"
)

// ToCamel converts a string to camel
func ToCamel(s string) string {
	return camel(s, "", StringsSep())
}

func ToDotCamel(s string) string {
	return camel(s, StringsSepDot, StringsSep())
}

func ToSpaceCamel(s string) string {
	return camel(s, StringsSepSpace, StringsSep())
}

func camel(str string, sep string, SkipSep map[string]byte) string {
	var newStr string

	nextUpper := true
	for _, s := range str {
		sStr := string(s)

		// 遇到分隔符，跳过，下个字母重置为大写
		if _, ok := SkipSep[sStr]; ok {
			nextUpper = true
			continue
		}

		// 保留原有大写不变
		if unicode.IsUpper(s) == true {
			nextUpper = true
		}

		// 分割符，下个字母为大写 && 不为首字母
		if (nextUpper == true) && (len(newStr) > 0) {
			newStr += sep
		}

		// 判断大小写
		if nextUpper == true {
			newStr += strings.ToUpper(sStr)
		} else {
			newStr += strings.ToLower(sStr)
		}
		nextUpper = false
	}

	return newStr
}
