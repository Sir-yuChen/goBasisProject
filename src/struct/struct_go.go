package main

import (
	"fmt"
	"reflect"
)

type resume struct {
	//结构体标签 `key：value`格式
	Name string `json:"name" doc:"我的名字"`
}

func findDoc(stru interface{}) map[string]string {
	//reflect反射
	t := reflect.TypeOf(stru).Elem()
	doc := make(map[string]string)

	for i := 0; i < t.NumField(); i++ {
		//结构体标签
		doc[t.Field(i).Tag.Get("json")] = t.Field(i).Tag.Get("doc")
	}
	return doc
}

func main() {
	var stru resume
	doc := findDoc(&stru)
	fmt.Printf("name字段为：%s\n", doc["name"])
}
