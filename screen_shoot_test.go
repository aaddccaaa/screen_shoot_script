package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"github.com/vcaesar/tt"
	"os"
	"testing"
	"time"
)

func TestBitmap(t *testing.T) {
	filePath, _ := os.Getwd() // 保存截图的文件路径
	filePath = fmt.Sprintf("%s/screen_shoot_store", filePath)

	bit := robotgo.CaptureScreen()
	tt.NotNil(t, bit)
	timestamp := time.Now().Format("20060102150405")
	fileName := fmt.Sprintf("%s/screenshot_%s.jpg", filePath, timestamp)
	e := robotgo.SaveBitmap(bit, fileName)
	tt.Empty(t, e)
}

func TestBitmapV1(t *testing.T) {
	bit := robotgo.CaptureScreen()
	tt.NotNil(t, bit)
	e := robotgo.SaveBitmap(bit, "screen_shoot_store/test1.jpg")
	tt.Empty(t, e)
}
