package ecs

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"gitlab.landui.cn/gomod/logs"
)

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

// SetIPBlackWhiteListRule 设置ip黑白名单
func (e *ECS) SetIPBlackWhiteListRule(productTypeName string, ipList []string) error {
	header := map[string]string{
		"User-Agent":  "api-landui-lan",
		"X-RequestAU": fmt.Sprintf("%d|\t|%s", e.UserId, e.UserName),
	}
	ipListStr, _ := json.Marshal(ipList)
	data := map[string]string{
		"vpsname":           fmt.Sprintf("%d", e.Id),
		"product_type_name": productTypeName,
		"ip_list":           string(ipListStr),
	}
	client := resty.New()
	resp, err := client.R().SetHeaders(header).SetFormData(data).Post(e.APIUriPrefix + SetFirewallIPBlackWhiteAPI)
	if err != nil {
		logs.New().SetAdditionalInfo("err", err.Error()).Error("设置黑白名单发送请求错误", err)
		return fmt.Errorf("设置黑白名单发送请求错误: %s", err.Error())
	}
	logs.New().
		SetAdditionalInfo("header", header).
		SetAdditionalInfo("data", data).
		SetAdditionalInfo("url", e.APIUriPrefix+SetFirewallIPBlackWhiteAPI).
		SetAdditionalInfo("resp", resp.String()).Info("设置黑白名单记录")

	if resp.StatusCode() != 200 {
		return fmt.Errorf("设置黑白名单失败: %s", resp.String())
	}

	return nil
}

// DeleteIPBlackWhiteListRule 删除云防火墙黑白名单
func (e *ECS) DeleteIPBlackWhiteListRule(productTypeName string, ipList []string) error {

	header := map[string]string{
		"User-Agent":  "api-landui-lan",
		"X-RequestAU": fmt.Sprintf("%d|\t|%s", e.UserId, e.UserName),
	}
	ipListStr, _ := json.Marshal(ipList)
	data := map[string]string{
		"vpsname":           fmt.Sprintf("%d", e.Id),
		"product_type_name": productTypeName,
		"ip_list":           string(ipListStr),
	}
	client := resty.New()
	resp, err := client.R().SetHeaders(header).SetFormData(data).Post(e.APIUriPrefix + DeleteFirewallIPBlackWhiteAPI)
	if err != nil {
		logs.New().SetAdditionalInfo("err", err.Error()).Error("删除黑白名单发送请求错误", err)
		return fmt.Errorf("删除黑白名单发送请求错误: %s", err.Error())
	}
	logs.New().
		SetAdditionalInfo("header", header).
		SetAdditionalInfo("data", data).
		SetAdditionalInfo("url", e.APIUriPrefix+DeleteFirewallIPBlackWhiteAPI).
		SetAdditionalInfo("resp", resp.String()).Info("删除黑白名单记录")

	if resp.StatusCode() != 200 {
		return fmt.Errorf("删除黑白名单失败: %s", resp.String())
	}

	return nil
}
