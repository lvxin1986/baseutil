package reflect

import (
	"fmt"
	"testing"
)

/*
Copyright 2019 Tiglabs

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreedto in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

created by lvxin for project baseutil at 19-2-22 下午3:42
*/

type User struct {
	Id      int
	Name    string
	Address *Address
}

type Address struct {
	Add string
	Res int
}

func TestToString(t *testing.T) {
	//u := User{Id: 1001, Name: "aaa", Address: &Address{Add: "ccccccccc", Res: 12}}
	fmt.Print(ToString("u", "test"))

}
