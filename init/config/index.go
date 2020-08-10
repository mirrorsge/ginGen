package config

import "os"

func GenConfig() {
	err := os.Mkdir("./config", 0644)
	if err != nil {
		panic("创建config目录失败: " + err.Error())
	}
	file, err := os.Create("./config/config.toml")

	return
}
