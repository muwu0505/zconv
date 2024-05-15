package zconv

import (
	"strings"
	"unicode"
)

// ToCustomCamel 自定义驼峰
// ori 原字符串
// sep 分隔符
// nextUpper 首字母是否大写
func ToCustomCamel(ori string, sep string, firstUpper bool) (dst string) {
	nextUpper := firstUpper
	for _, s := range ori {
		// 仅保留字母，其余全部认为是分隔符，跳过，下个字母重置为大写
		if unicode.IsLetter(s) == false {
			nextUpper = true
			continue
		}

		// 保留原有大写不变
		if unicode.IsUpper(s) == true {
			nextUpper = true
		}

		// 首字母按传入的 nextUpper 判断大小写
		if len(dst) <= 0 {
			nextUpper = firstUpper
		}

		// 分割符，下个字母为大写 && 不为首字母
		if (nextUpper == true) && (len(dst) > 0) {
			dst += sep
		}

		// 判断大小写
		if nextUpper == true {
			dst += strings.ToUpper(string(s))
		} else {
			dst += strings.ToLower(string(s))
		}

		// 下个字母重置为小写
		nextUpper = false
	}

	return
}
