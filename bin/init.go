package bin

import (
	"fmt"
	"io/ioutil"
	"os"
)

func Init() {
	pwd, _ := os.Getwd()
	fileList, err := ioutil.ReadDir(pwd)
	if err != nil {
		panic("获取当前目录失败," + err.Error())
	}
	//当前目录不为空，提示是否进行格式化
	if len(fileList) > 1 || (len(fileList) == 1 && fileList[0].Name() != ".git") {
		var in string
		fmt.Println("current directory is not empty, are you sure to initialize?")
		_, _ = fmt.Scanln(&in)
		if in == "N" || in == "n" {
			fmt.Println("initialize abort")
			os.Exit(0)
		}
		//清空目录
		for _, item := range fileList {
			filename := "./" + item.Name()
			if item.IsDir() {
				if item.Name() != ".git" {
					err = os.RemoveAll(filename)
				}
			} else {
				err = os.Remove(filename)
			}
			if err != nil {
				panic("初始化当前目录失败," + err.Error())
			}
		}
	}
	//创建项目目录结构

	return
}
