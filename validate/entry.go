package validate

import (
	"fmt"
	"reflect"
	"strings"
)

type Entry struct {
}

//https://github.com/astaxie/beego
func (v *Entry) Validate(obj interface{}) (bool, error) {
	objT := reflect.TypeOf(obj)
	objV := reflect.ValueOf(obj)
	switch {
	case isStruct(objT):
	case isStructPtr(objT):
		objT = objT.Elem()
		objV = objV.Elem()
	default:
		return false, fmt.Errorf("%v must be a struct or a struct pointer", obj)
	}
	for i := 0; i < objT.NumField(); i++ {
		fmt.Println(objT.Field(i).Type.Kind() == reflect.Struct)
		tag := getValidTag(objT.Field(i))
		fmt.Println(tag)
	}
	return true, nil
}

func isStruct(t reflect.Type) bool {
	return t.Kind() == reflect.Struct
}

func isStructPtr(t reflect.Type) bool {
	return t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Struct
}

func getValidTag(f reflect.StructField) string {
	tags := f.Tag.Get(Tag)
	if len(tags) == 0 {
		return ""
	}
	tagArray := strings.Split(tags, ";")
	for _, tag := range tagArray {
		if len(tag) == 0 {
			continue
		}
		test(tag, f.Name)
	}
	return tags
}

func test(tag, fieldName string) {
	fmt.Println(fmt.Sprintf("%s -> %s", fieldName, tag))
}
