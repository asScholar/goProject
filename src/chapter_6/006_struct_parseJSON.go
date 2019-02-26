/*
手机拥有屏幕、电池、指纹识别等信息，将这些信息填充为JSON格式的数据。
如果需要选择性地分离JSON中的数据则较为麻烦。Go语言中的匿名结构体可以方便地完成这个操作。
*/
package main

import (
	"encoding/json"
	"fmt"
)

/*模拟手机屏幕*/
type Screen struct {
	Size       float32
	ResX, ResY int
}

/*模拟电池*/
type Battery struct {
	Capacity int
}

/*生成JSON数据*/
func genJSONData() []byte {
	/*完整的数据结构*/
	raw := &struct {
		Screen
		Battery
		HasTouchID bool //序列化时添加的字段：是否有指纹识别
	}{
		Screen: Screen{
			Size: 5.5,
			ResX: 1920,
			ResY: 1080,
		},

		Battery: Battery{
			2910,
		},

		HasTouchID: true,
	}

	jsonData, _ := json.Marshal(raw)

	return jsonData
}

func main() {
	jsonData := genJSONData()
	fmt.Println(string(jsonData))

	/*只需要屏幕和指纹信息*/
	screenAndTouch := struct {
		Screen
		HasTouchID bool
	}{}
	/*反序列化*/
	json.Unmarshal(jsonData, &screenAndTouch)
	fmt.Println(screenAndTouch)

}
