package initialize

import (
	"fmt"
	"github.com/mirrorsge/ginGen/initialize/config"
	"github.com/mirrorsge/ginGen/initialize/controller"
	"github.com/mirrorsge/ginGen/initialize/middleware"
	"github.com/mirrorsge/ginGen/initialize/model"
	"github.com/mirrorsge/ginGen/initialize/router"
	"github.com/mirrorsge/ginGen/initialize/runtime"
	"io/ioutil"
	"os"
	"os/exec"
)

var ProjectName string

func Init() {
	pwd, _ := os.Getwd()
	fileList, err := ioutil.ReadDir(pwd)
	if err != nil {
		panic("获取当前目录失败," + err.Error())
	}
	//todo 当前目录不为空，提示是否进行格式化(暂时不允许初始化非空目录)
	if len(fileList) > 1 || (len(fileList) == 1 && fileList[0].Name() != ".git") {
		panic("初始化当前目录失败,目录非空目录")
		//var in string
		//fmt.Println("current directory is not empty, are you sure to initialize?")
		//_, _ = fmt.Scanln(&in)
		//if in == "N" || in == "n" {
		//	fmt.Println("initialize abort")
		//	os.Exit(0)
		//}
		////清空目录
		//for _, item := range fileList {
		//	filename := "./" + item.Name()
		//	if item.IsDir() {
		//		if item.Name() != ".git" {
		//			err = os.RemoveAll(filename)
		//		}
		//	} else {
		//		err = os.Remove(filename)
		//	}
		//	if err != nil {
		//		panic("初始化当前目录失败," + err.Error())
		//	}
		//}
	}
	//创建项目目录结构
	//初始化git mod 配置
	fmt.Println("please enter your project name(it will also set in your go.mod file): ")
	_, _ = fmt.Scanln(&ProjectName)
	for ProjectName == "" {
		fmt.Println("项目名不能为空")
		_, _ = fmt.Scanln(&ProjectName)
	}
	command := exec.Command("go", "mod", "init", ProjectName)
	command.Stdout = os.Stdout
	err = command.Run()
	if err != nil {
		fmt.Println(err.Error())
	}
	//生成对应文件结构
	config.GenConfig()
	controller.GenController()
	middleware.GenMiddleware()
	model.GenModel()
	router.GenRouter()
	runtime.GenRuntime()
	return
}
