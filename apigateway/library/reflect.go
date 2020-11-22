package library

import (
	"reflect"
	"strings"

	"gorm.io/gorm"
)

//Find contains
func Find(a []string, x string) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return -1
}

//IsIn checks model fields
func IsIn(val string, x interface{}) bool {
	gormFields := Fields(gorm.Model{})
	modelFields := Fields(x)
	modelFields = append(modelFields, gormFields...)
	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return false
	}
	return Find(modelFields, vals[0]) != -1
}

//Fields gets fields of Struct
func Fields(x interface{}) []string {
	v := reflect.ValueOf(x)
	typeOfS := v.Type()
	var fields []string
	for i := 0; i < v.NumField(); i++ {
		fields = append(fields, typeOfS.Field(i).Name)
	}
	return fields
}
