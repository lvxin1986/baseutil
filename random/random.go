package random

import (
	"crypto/rand"
	"fmt"
	"github.com/lvxin1986/baseutil/runtime/stack"
	"github.com/spf13/cast"
	"math/big"
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
//create by bjlvxin at 10:17 for project baseutil

func GetRandomInt(number int) int64 {
	if result, err := rand.Int(rand.Reader, big.NewInt(cast.ToInt64(number))); err != nil {
		fmt.Println(err)
		stack.PrintRuntimeFullStack()
		return 0
	} else {
		return result.Int64()
	}

}
