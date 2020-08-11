package config

import "os"

func GenConfig() {
	err := os.Mkdir("./config", 0766)
	if err != nil {
		panic("创建config目录失败: " + err.Error())
	}
	genToml()
	genConf()
	return
}
