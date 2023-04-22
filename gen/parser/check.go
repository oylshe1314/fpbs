package parser

import "regexp"

func CheckPackageName(packageName string) bool {
	if len(packageName) == 0 {
		return false
	}
	exp, err := regexp.Compile("^[_A-Za-z][_0-9A-Za-z]*$")
	if err != nil {
		return false
	}

	return exp.MatchString(packageName)
}

func CheckStructName(structName string) bool {
	return CheckPackageName(structName)
}

func CheckFieldName(fieldName string) bool {
	return CheckPackageName(fieldName)
}
