package zconv

// ToCamel 转换大驼峰，转换为 "ToCamel"
func ToCamel(s string) string {
	return ToCustomCamel(s, "", true)
}

// ToDotCamel 转换以【.】连接的大驼峰，转换为 "To.Camel"
func ToDotCamel(s string) string {
	return ToCustomCamel(s, StringsSepDot, true)
}

// ToSpaceCamel 转换以【 】连接的大驼峰，转换为 "To Camel"
func ToSpaceCamel(s string) string {
	return ToCustomCamel(s, StringsSepSpace, true)
}
