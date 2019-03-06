package reflect

import (
	"fmt"
	"reflect"
)

// Copyright 2019 The ChuBao Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.
//create by bjlvxin at 16:06 for project baseutil

func StructToString(prefix string, u interface{}) (str string) {

	t := reflect.TypeOf(u)
	v := reflect.ValueOf(u)
	returnStr := ""
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).CanInterface() { //判断是否为可导出字段
			//判断是否是嵌套结构
			if v.Field(i).Type().Kind() == reflect.Struct {
				returnStr = returnStr + StructToString(prefix+"."+t.Field(i).Name+"", v.Field(i).Interface())
				continue
			} else if v.Field(i).Type().Kind() == reflect.Ptr {
				returnStr = returnStr + PtrToString(prefix+"."+t.Field(i).Name+"", v.Field(i).Interface())
				continue
			} else {
				returnStr = returnStr + fmt.Sprintf("%s.%s = %v \n", prefix, t.Field(i).Name, v.Field(i).Interface())
			}
		}
	}
	return returnStr
}

func PtrToString(prefix string, i interface{}) (str string) {
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)
	returnStr := ""
	for i := 0; i < v.Elem().NumField(); i++ {
		if v.Elem().Field(i).CanInterface() { //判断是否为可导出字段
			//判断是否是嵌套结构
			if v.Elem().Field(i).Type().Kind() == reflect.Struct {
				returnStr = returnStr + StructToString(prefix+"."+t.Elem().Field(i).Name+"", v.Elem().Field(i).Interface())
				continue
			} else if v.Elem().Field(i).Type().Kind() == reflect.Ptr {
				returnStr = returnStr + PtrToString(prefix+"."+t.Elem().Field(i).Name+"", v.Elem().Field(i).Interface())
				continue
			} else {
				returnStr = returnStr + fmt.Sprintf("%s.%s = %v \n", prefix, t.Elem().Field(i).Name, v.Elem().Field(i).Interface())
			}
		}
	}
	return returnStr
}

func ToString(prefix string, u interface{}) (str string) {
	t := reflect.TypeOf(u)
	returnStr := ""
	if t.Kind() == reflect.Ptr {
		returnStr = PtrToString(prefix, u)
	} else if t.Kind() == reflect.Struct {
		returnStr = StructToString(prefix, u)
	}
	return returnStr
}
