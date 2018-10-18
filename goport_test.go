package goport

import (
	"fmt"
	"testing"
)

// 测试Command初始化
func TestGetFreePort(t *testing.T) {
	var port = GetFreePort()
	if port < 0 {
		t.Errorf("本地5000-50000的范围内没有可用的端口!")
	}

	if !CheckFree(port) {
		t.Errorf("获取到的端口不可用!")
	}

	fmt.Printf("成功获得可用端口:%d", port)
	fmt.Println()
}
