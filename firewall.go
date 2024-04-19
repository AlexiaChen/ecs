package ecs

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"gitlab.landui.cn/gomod/logs"
)

type AclData struct {
	Weight     string `json:"weight"`    // 不可重复，当为空时php接口会自动生成
	Direction  string `json:"direction"` // Inbound  Outbound
	Action     string `json:"action"`    // Deny  Allow
	Protocol   string `json:"protocol"`  // ANY  TCP ICMP UDP
	RemoteIP   string `json:"local_ip"`  // 10.10.1.1 10.10.0.0/24
	RemotePort string `json:"local_port"`
	LocalIP    string `json:"remote_ip"`
	LocalPort  string `json:"remote_port"` // 3306 6379 1-65535
	Desc       string `json:"desc"`
}

type aclResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// SetFirewall 设置云服务器云防火墙
func (e *ECS) SetFirewall(productTypeName string) error {
	header := map[string]string{
		"User-Agent":  "api-landui-lan",
		"X-RequestAU": fmt.Sprintf("%d|\t|%s", e.UserId, e.UserName),
	}
	data := map[string]string{
		"vpsname":           fmt.Sprintf("%d", e.Id),
		"product_type_name": productTypeName,
	}
	client := resty.New()
	resp, err := client.R().SetHeaders(header).SetFormData(data).Post(e.APIUriPrefix + SetFirewallAPI)
	if err != nil {
		logs.New().SetAdditionalInfo("err", err.Error()).Error("设置防火墙发送请求错误", err)
		return err
	}
	logs.New().
		SetAdditionalInfo("header", header).
		SetAdditionalInfo("data", data).
		SetAdditionalInfo("url", e.APIUriPrefix+SetFirewallAPI).
		SetAdditionalInfo("resp", resp.String()).Info("开通主机设置防火墙记录")
	return nil
}

// AddFirewallACL 设置ip黑白名单
func (e *ECS) AddFirewallACL(acl *AclData) error {
	header := map[string]string{
		"User-Agent":  "api-landui-lan",
		"X-RequestAU": fmt.Sprintf("%d|\t|%s", e.UserId, e.UserName),
	}
	data := map[string]interface{}{
		"username":   e.UserName,
		"vpsID":      e.Id,
		"aclGroupID": "0",
		"aclData":    acl,
	}
	var result aclResponse
	client := resty.New()
	resp, err := client.R().SetHeaders(header).SetBody(data).SetResult(&result).Post(e.APIUriPrefix + AddFirewallACLAPI)
	if err != nil {
		logs.New().SetAdditionalInfo("err", err.Error()).Error("设置黑白名单发送请求错误", err)
		return fmt.Errorf("设置黑白名单发送请求错误: %s", err.Error())
	}
	logs.New().
		SetAdditionalInfo("header", header).
		SetAdditionalInfo("data", data).
		SetAdditionalInfo("url", e.APIUriPrefix+AddFirewallACLAPI).
		SetAdditionalInfo("resp", resp.String()).Info("设置黑白名单记录")

	if result.Code != 200 {
		return fmt.Errorf("设置黑白名单失败: %s", resp.String())
	}
	return nil
}

// DelFirewallACL 设置ip黑白名单
func (e *ECS) DelFirewallACL(acl *AclData) error {
	header := map[string]string{
		"User-Agent":  "api-landui-lan",
		"X-RequestAU": fmt.Sprintf("%d|\t|%s", e.UserId, e.UserName),
	}
	data := map[string]interface{}{
		"username":   e.UserName,
		"vpsID":      e.Id,
		"aclGroupID": "0",
		"aclData":    acl,
	}
	var result aclResponse
	client := resty.New()
	resp, err := client.R().SetHeaders(header).SetBody(data).SetResult(&result).Post(e.APIUriPrefix + DelFirewallACLAPI)
	if err != nil {
		logs.New().SetAdditionalInfo("err", err.Error()).Error("设置黑白名单发送请求错误", err)
		return fmt.Errorf("删除黑白名单发送请求错误: %s", err.Error())
	}
	logs.New().
		SetAdditionalInfo("header", header).
		SetAdditionalInfo("data", data).
		SetAdditionalInfo("url", e.APIUriPrefix+DelFirewallACLAPI).
		SetAdditionalInfo("resp", resp.String()).Info("设置黑白名单记录")

	if result.Code != 200 {
		return fmt.Errorf("删除黑白名单失败: %s", resp.String())
	}
	return nil
}
