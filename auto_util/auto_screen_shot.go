package auto_util

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"time"
)

var (
	BeforeScreenShoot = 2
	XStart            = 200
	XEnd              = 1860
	YStart            = 180
	YEnd              = 980
)

func AutoScreenShoot(interval, maxCount int, filePath string) error {

	time.Sleep(time.Duration(BeforeScreenShoot) * time.Second)

	// 检查最大截图数量是否为正数
	if maxCount <= 0 {
		return fmt.Errorf("最大截图数量必须为正数")
	}

	// 循环截图
	for i := 0; i < maxCount; i++ {
		// 截图
		bitmap := robotgo.CaptureScreen(XStart, YStart, XEnd-XStart, YEnd-YStart)

		// 生成保存路径和文件名
		timestamp := time.Now().Format("20060102150405")
		fileName := fmt.Sprintf("%s/screenshot_%s.jpg", filePath, timestamp)

		robotgo.SaveBitmap(bitmap, fileName)

		// 释放截图资源
		robotgo.FreeBitmap(bitmap)

		// 检查是否达到最大截图数量
		if i+1 == maxCount {
			break
		}

		// 等待指定的时间间隔
		time.Sleep(time.Duration(interval) * time.Second)
		// 模拟按下方向键
		robotgo.KeyTap("down")
	}
	return nil
}
