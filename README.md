# ginGen
####gin project generator
####gin项目脚手架
####功能
1. genGen I 创建初始化gin项目
2. genGen U 根据router文件中的路由，生成对应的controller文件/代码

####已实现
1. swagger支持
2. gorm支持


####前置
1. go modules开启
2. $GOPATH/bin 已加入环境变量(PATH)

####使用方法
1. go get github.com/mirrorsge/ginGen
2. 在空目录中，执行ginGen I ，会初始化整个项目。注：会删除目录中原有的所有文件，不包括'.git'目录，可以先初始化git，再执行此命令。



