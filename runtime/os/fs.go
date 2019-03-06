package os

import (
	"os/exec"
	"os"
	"path/filepath"
	"strings"
	"errors"
	"github.com/lvxin1986/baseutil/runtime/stack"
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

func GetCurrentPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	i := strings.LastIndex(path, "/")
	if i < 0 {
		i = strings.LastIndex(path, "\\")
	}
	if i < 0 {
		return "", errors.New(`error: Can't find "/" or "\".`)
	}
	return string(path[0 : i+1]), nil
}

func PathExist(p string) bool {
	_, err := os.Stat(p)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

func GetCurrentSourceCodePath() (fileName string, err error){
	_, fileName, _, ok := stack.CallerName(2)
	if !ok {
		err = errors.New("Can not get the current source code path!")
	}
	return fileName,err
}
