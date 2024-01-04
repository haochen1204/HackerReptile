package console

import (
	"HackerReptile/reptile"
	"HackerReptile/templateType"
	"github.com/spf13/viper"
	"github.com/thep0y/go-logger/log"
)

func Console() {
	yamlfile := viper.GetString("yaml")
	if yamlfile == "" {
		log.Error("请输入爬虫的yaml文件")
	} else {
		tmp := ReadYamlReptile(yamlfile)
		// 循环遍历每一个Headless节点
		for _, headlessVal := range tmp.Headless {
			// 循环遍历每一个Step
			reptile.CreatActionQuery(headlessVal.Steps)
		}
	}
}

func ReadYamlReptile(yamlfile string) templateType.Template {
	v := viper.New()
	v.SetConfigFile(yamlfile) // 设置文件路径
	v.SetConfigType("yaml")   // 设置文件类型为yaml

	if err := v.ReadInConfig(); err != nil { // 读取并解析文件
		log.Error("Error reading config file, %s", err)
	}

	var tmp templateType.Template
	err := v.Unmarshal(&tmp) // 将获取到的数据解码到相应的结构体中
	if err != nil {
		log.Error("Unable to decode into struct, %v", err)
	}

	return tmp
}
