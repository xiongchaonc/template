package template

import (
	"sync"
	"html/template"
	"fmt"
	"os"
)

var (
	temp  *template.Template
	temps sync.Map
)

func Temp(name string) *template.Template {
	if ret, ok := temps.Load(name); ok {
		return ret.(*template.Template)
	}

	temp = newTmp(name)
	temps.Store(name, temp)
	return temp
}

func newTmp(name string) *template.Template {
	//需要判断文件是否存在
	ok := PathExists(name)
	if !ok {
		panic(name + " template is not exists")
	}

	myTemplate, err := template.ParseFiles(name)
	if err != nil {
		fmt.Println("parse file err:", err)
		panic(name + " template init error")
	}

	return myTemplate
}


// 判断文件夹是否存在
func PathExists(path string) (bool) {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
