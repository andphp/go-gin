package utils

import (
	"fmt"
	"reflect"

	"github.com/fatih/structs"
)

func Struct2Slice(f interface{}) []string {
	v := reflect.ValueOf(f)
	ss := make([]string, v.NumField())
	for i := range ss {
		ss[i] = fmt.Sprintf("%v", v.Field(i))
	}
	return ss
}

func Struct2ColumnsValues(model interface{}, tag string) (columns []string, values []interface{}) {
	s := structs.New(model)
	for _, f := range s.Fields() {
		if f.IsEmbedded() {
			sf := s.Field(f.Name())
			columnsC, valuesC := recursionOfModel2ColumnsValues(sf)
			columns = append(columns, columnsC...)
			values = append(values, valuesC...)
		} else {
			column := f.Tag(tag)
			if column == "" {
				column = SnakeString(f.Name())
			}
			columns = append(columns, column)
			values = append(values, f.Value())
		}
	}
	//fmt.Println("columns :", columns)
	//fmt.Println("values :", values)
	return
}

func recursionOfModel2ColumnsValues(s *structs.Field) (columns []string, values []interface{}) {
	for _, f := range s.Fields() {
		if f.IsEmbedded() {
			sf := s.Field(f.Name())
			columnsC, valuesC := recursionOfModel2ColumnsValues(sf)
			columns = append(columns, columnsC...)
			values = append(values, valuesC...)
		} else {
			column := f.Tag("form")
			if column == "" {
				column = SnakeString(f.Name())
			}
			columns = append(columns, column)
			values = append(values, f.Value())
		}
	}
	return
}

// ToMap 结构体转为Map[string]interface{}
func Struct2Map(in interface{}, tagName string) (map[string]interface{}, error) {
	out := make(map[string]interface{})

	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct { // 非结构体返回错误提示
		return nil, fmt.Errorf("ToMap only accepts struct or struct pointer; got %T", v)
	}

	t := v.Type()
	// 遍历结构体字段
	// 指定tagName值为map中key;字段值为map中value
	for i := 0; i < v.NumField(); i++ {
		fi := t.Field(i)
		if tagValue := fi.Tag.Get(tagName); tagValue != "" {
			out[tagValue] = v.Field(i).Interface()
		}
	}
	return out, nil
}
