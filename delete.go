package ecs

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/util/guid"
	"gitlab.landui.cn/gomod/logs"
	"time"
)

// Delete 删除云服务器
func (e *ECS) Delete() (*Response, error) {
	if e.Id == 0 {
		return nil, errors.New("主机不存在")
	}
	times := time.Now().Unix()
	randStr := guid.S()
	text := fmt.Sprintf("%d%s%s", times, randStr, e.SignSecret)
	newText := sorts(text)
	ciphertext := gmd5.MustEncryptString(newText)
	body := map[string]interface{}{
		"time_stamp": fmt.Sprintf("%d", times),
		"nonce_str":  randStr,
		"sign":       ciphertext,
		"id":         e.Id,
		"userid":     e.UserId,
	}
	logs.New().SetAdditionalInfo("body", body).Info("发送删除主机的时候请求的body")
	resp, err := post(body, e.DeleteAPIUrl)
	if err != nil {
		logs.New().Error("发送删除主机的时候请求失败了", err)
		return nil, err
	}
	logs.New().SetAdditionalInfo("resp", resp).Info("发送删除主机的时候请求的响应")
	res := new(Response)
	err = json.Unmarshal(resp.Body(), res)
	if err != nil {
		logs.New().Error("获取到php返回的内容json反序列化的时候失败了", err)
		return nil, errors.New("删除主机失败，解析失败")
	}
	return res, err

}
