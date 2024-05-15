package ecs

import (
	"fmt"

	"github.com/go-resty/resty/v2"
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
		return fmt.Errorf("设置防火墙发送请求错误: %s resp: %s data: %+v header: %+v", err.Error(), resp.String(), data, header)
	}
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
		return fmt.Errorf("设置黑白名单发送请求错误: %s resp: %s data: %+v header: %+v", err.Error(), resp.String(), data, header)
	}

	if result.Code != 200 {
		return fmt.Errorf("设置黑白名单失败: %s status code %d", resp.String(), result.Code)
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
		return fmt.Errorf("删除黑白名单发送请求错误: %s resp: %s data: %+v header: %+v", err.Error(), resp.String(), data, header)
	}

	if result.Code != 200 {
		return fmt.Errorf("删除黑白名单失败: %s status code: %d", resp.String(), result.Code)
	}
	return nil
}
