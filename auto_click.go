package main

import (
	"auto_click/auto_util"
	"fmt"
)

func main() {
	interval := 1                    // 截图时间间隔（秒）
	maxCount := 50                   // 最大截图数量
	filePath := "screen_shoot_store" // 保存截图的文件路径

	err := auto_util.AutoScreenShoot(interval, maxCount, filePath)
	if err != nil {
		fmt.Println("截图出错:", err)
	} else {
		fmt.Println("截图完成")
	}
}
