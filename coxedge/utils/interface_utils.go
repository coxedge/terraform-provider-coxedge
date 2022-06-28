/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
package utils

func ConvertListInterfaceToStringArray(orig interface{}) []string {
	interfaceList := orig.([]interface{})
	stringList := make([]string, len(interfaceList))
	for i, s := range interfaceList {
		stringList[i] = s.(string)
	}
	return stringList
}

func BoolAddr(x bool) *bool {
	return &x
}

func StringAddr(x string) *string {
	return &x
}
