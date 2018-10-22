package main

import (
	"fmt"
	"runtime"

	"github.com/BurntSushi/toml"
	"github.com/astaxie/beego"
	"github.com/bitly/go-simplejson"
	"github.com/singcl/go-dep-glide/config"
	"github.com/skip2/go-qrcode"
	"log"
	"image/color"
)

func main() {
	configContent()
	generateQrcode()

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
	var con config.TomlConfig
	filePath := "conf.toml"
	if _, err := toml.DecodeFile(filePath, &con); err != nil {
		panic(err)
	}
	fmt.Printf("conf.toml的内容为：%v", con)
}

// 生成二维码
func generateQrcode() {
	// 快捷方式生成二维码
	//qrcode.WriteFile("https://baidu.com", qrcode.Medium, 256, "baidu.png")

	// 自定义二维码
	qr, err := qrcode.New("https://baidu.com", qrcode.Medium)
	if err != nil {
		log.Fatal(err)
	} else {
		qr.BackgroundColor = color.RGBA{50, 205, 50, 255 }
		qr.ForegroundColor = color.White
		qr.WriteFile(256, "./baidu.png")
	}
}
