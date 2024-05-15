package ecs

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/util/guid"
)

// Delete 删除云服务器
func (e *ECS) Delete() (*Response, error) {
	if e.Id == 0 {
		return nil, errors.New("主机不存在")
	}
	times := time.Now().Unix()
	randStr := guid.S()
	text := fmt.Sprintf("%d%s%s", times, randStr, e.APISignSecret)
	newText := sorts(text)
	ciphertext := gmd5.MustEncryptString(newText)
	body := map[string]interface{}{
		"time_stamp": fmt.Sprintf("%d", times),
		"nonce_str":  randStr,
		"sign":       ciphertext,
		"id":         e.Id,
		"userid":     e.UserId,
	}
	resp, err := post(body, e.APIUriPrefix+DeleteHostAPI)
	if err != nil {
		return nil, fmt.Errorf("发送删除主机的时候请求失败了: %s resp: %s request body: %+v", err.Error(), resp.String(), body)
	}
	res := new(Response)
	err = json.Unmarshal(resp.Body(), res)
	if err != nil {
		return nil, fmt.Errorf("获取到php返回的内容json反序列化的时候失败了: %s resp: %s request body: %+v", err.Error(), resp.String(), body)
	}
	return res, err

}
