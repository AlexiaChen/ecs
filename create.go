package ecs

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/util/guid"
)

// Create 创建云服务器
func (e *ECS) Create() (*Response, error) {
	times := time.Now().Unix()
	randStr := guid.S()
	text := fmt.Sprintf("%d%s%s", times, randStr, e.APISignSecret)
	newText := sorts(text)
	ciphertext := gmd5.MustEncryptString(newText)
	diskMap, hardDiskSize := getHardDisks(e.Disks, e.HardDisks)
	body := map[string]interface{}{
		"time_stamp":   fmt.Sprintf("%d", times),
		"nonce_str":    randStr,
		"sign":         ciphertext,
		"serverid":     e.RegionId,
		"nodeid":       fmt.Sprint(e.NodeId), //129,130,131 研发集群
		"months":       e.Months,
		"system_type":  e.SystemType,
		"systemid":     e.SystemId,
		"disktype":     e.DiskType,
		"cpu":          e.Cpu,
		"memory":       fmt.Sprint(e.Memory),
		"bandwidth":    fmt.Sprint(e.Bandwidth / 1000),
		"harddisks":    hardDiskSize,
		"disks":        diskMap,
		"defense":      e.Defense,
		"userid":       e.UserId,
		"is_intranet":  fmt.Sprint(e.IsIntranet),  //为1代表是内网ip
		"product_type": fmt.Sprint(e.ProductType), //rds的标识 redis 2 云防火墙
	}
	resp, err := post(body, e.APIUriPrefix+CreateHostAPI)
	if err != nil {
		return nil, fmt.Errorf("创建主机失败: %s resp: %s request body: %+v", err.Error(), resp.String(), body)
	}
	res := new(Response)
	err = json.Unmarshal(resp.Body(), res)
	if err != nil {
		return nil, fmt.Errorf("获取到php返回的内容json反序列化的时候失败了: %s resp: %s request body: %+v", err.Error(), resp.String(), body)
	}
	return res, nil
}
