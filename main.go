package main

import (
	"fries/plugin"
)

func main() {
	// 	工具开发目的：
	// 		1、输入url识别指纹
	// 		2、调取fofa资产的url识别指纹
	// 		3、调取ZoomEye资产的url识别指纹
	var f plugin.Finger
	plugin.Flag(&f)
}
