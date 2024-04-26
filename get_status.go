package ecs

import (
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/util/guid"
	"gitlab.landui.cn/gomod/logs"
	"time"
)

const (
	VMStatusUnknown  = -1
	VMStatusNotRun   = 0
	VMStatusRunning  = 1
	VMStatusPowerOFF = -2
	VMStatusNotFound = -3
)

// GetStatus 调用PHP服务API获取云主机状态
func (e *ECS) GetStatus() (*Response, int8) {
	response := new(Response)
	times := time.Now().Unix()
	randStr := guid.S()
	text := fmt.Sprintf("%d%s%s", times, randStr, e.APISignSecret)
	newText := sorts(text)
	ciphertext := gmd5.MustEncryptString(newText)
	url := fmt.Sprintf(
		"%s?%s=%s&%s=%s&%s=%s&%s=%d",
		e.APIUriPrefix+GetHostStatusAPI,
		"time_stamp",
		fmt.Sprintf("%d", times),
		"nonce_str",
		randStr,
		"sign",
		ciphertext,
		"id",
		e.Id,
	)
	logs.New().SetAdditionalInfo("url", url).Info("获取主机状态")
	resp, err := unauthorizedPost(url)
	if err != nil {
		log := logs.New()
		log.SetAdditionalInfo("resp", resp).Error("获取主机状态失败", err)
		return response, VMStatusUnknown
	}
	err = json.Unmarshal(resp.Body(), response)
	if err != nil {
		log := logs.New()
		log.SetAdditionalInfo("resp", resp.String()).Error("获取主机状态参数解析失败", err)
		return response, VMStatusUnknown
	}

	if response.Info.Status == "" {
		return response, VMStatusUnknown
	}
	switch response.Info.Status {
	case "正常":
	case "运行中":
		return response, VMStatusRunning
	case "关机":
		return response, VMStatusPowerOFF
	case "不存在":
		return response, VMStatusNotFound
	default:
		return response, VMStatusUnknown
	}
	return response, VMStatusNotRun
}
