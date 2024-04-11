package ecs

import (
	"fmt"
	"gitlab.landui.cn/gomod/global"
	"go.uber.org/zap"
	"testing"
)

func TestGetStatus(t *testing.T) {
	initLog()
	newECS := ECS{
		APIUriPrefix:  "https://www.st.landui.cn",
		APISignSecret: "JMY13PagXnhl3rpiI1ht1hBBOaSF7dSOf8ktJ95zmOx19PWayRlyCtCm7UT0mghJ",
		Id:            100,
	}
	resp, code := newECS.GetStatus()
	fmt.Println(resp, code)
}
func initLog() {
	global.Logger = zap.NewExample() // 需要传入 zap.AddCaller() 才会显示打日志点的文件名和行数
}
