package main

import (
	"fmt"
	"runtime"

	"github.com/BurntSushi/toml"
	"github.com/astaxie/beego"
	simplejson "github.com/bitly/go-simplejson"
	"github.com/singcl/go-dep-slide/config"
)

func main() {
	configContent()
}

// 启动beego服务器
func runBeego() {
	maxCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(maxCPU)
	strJson := `{"announcer": {"nickname": "非议讲史", "kind": "user", "created_at": 1494904539000, "updated_at": 1494983507000, "track_id": 38088960}}`
	mapJson, _ := simplejson.NewJson([]byte(strJson))
	println(mapJson)
	beego.Run()
}

// 测试config包读取的内容
func configContent() {
	var config config.TomlConfig
	filePath := "conf.toml"
	if _, err := toml.DecodeFile(filePath, &config); err != nil {
		panic(err)
	}
	fmt.Printf("conf.toml的内容为：%v", config)
}
