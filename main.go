package main

import (
	"runtime"

	"github.com/astaxie/beego"
	simplejson "github.com/bitly/go-simplejson"
)

func main() {
	maxCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(maxCPU)
	strJson := `{"announcer": {"nickname": "非议讲史", "kind": "user", "created_at": 1494904539000, "updated_at": 1494983507000, "track_id": 38088960}}`
	mapJson, _ := simplejson.NewJson([]byte(strJson))
	println(mapJson)
	beego.Run()
}
