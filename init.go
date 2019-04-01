package template

import (
	"fmt"
	"time"
	"html/template"
)

func init() {
	Watch()
}

// Watch 定时处理模板文件，可以改成监听形式
func Watch() {
	go func() {
		for {
			time.Sleep(60 * time.Second)
			temps.Range(LoadTemp)
		}
	}()
}

//定时刷新模板内容
func LoadTemp(name, value interface{}) bool {
	value1, ok := LoadTmp(name.(string))
	if ok {
		value = value1
	}
	temps.Store(name, value)
	return true
}


//重新加载模板，如果不存在或者报错则返回之前的
func LoadTmp(name string) (*template.Template, bool) {
	//需要判断文件是否存在
	ok := PathExists(name)
	if !ok {
		return nil, false
	}

	myTemplate, err := template.ParseFiles(name)
	if err != nil {
		fmt.Println("parse file err:", err)
		return nil, false
	}

	return myTemplate, true
}
