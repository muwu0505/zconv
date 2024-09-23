package zconv

// ToLowerCamel 转换小驼峰，转换为 "toCamel"
func ToLowerCamel(s string) string {
	return ToCustomCamel(s, "", false)
}

// ToLowerDotCamel 转换以"."连接的小驼峰，转换为 "to.Camel"
func ToLowerDotCamel(s string) string {
	return ToCustomCamel(s, StringsSepDot, false)
}

// ToLowerSpaceCamel 转换以" "连接的小驼峰，转换为 "to Camel"
func ToLowerSpaceCamel(s string) string {
	return ToCustomCamel(s, StringsSepSpace, false)
}
