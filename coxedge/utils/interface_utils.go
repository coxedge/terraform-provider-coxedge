package utils

func ConvertListInterfaceToStringArray(orig interface{}) []string {
	interfaceList := orig.([]interface{})
	stringList := make([]string, len(interfaceList))
	for i, s := range interfaceList {
		stringList[i] = s.(string)
	}
	return stringList
}
