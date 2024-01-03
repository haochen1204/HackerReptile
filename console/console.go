package console

import (
	"HackerReptile/reptile"
	"github.com/spf13/viper"
	"github.com/thep0y/go-logger/log"
)

func Console() {
	url := viper.GetString("url")
	url_list := viper.GetString("list")
	if url != "" && url_list == "" {
		reptile.Test(url)
	} else if url == "" && url_list != "" {

	} else {
		log.Error("请输入-u或-l至少一个参数！")
	}
}
