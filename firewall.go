package ecs

import (
	"fmt"
	"gitlab.landui.cn/gomod/global"
	"gitlab.landui.cn/gomod/logs"
)

func SetFirewall(userId, UserName, vpsId string) {
	header := map[string]string{
		"User-Agent":  "api-landui-lan",
		"X-RequestAU": fmt.Sprintf("%s|\t|%s", userId, UserName),
	}
	data := map[string]string{
		"vpsname":           vpsId,
		"product_type_name": "rds_mysql",
	}
	client := global.HttpClient
	resp, err := client.R().SetHeaders(header).SetFormData(data).Post(global.ApiUrl["set_firewall"])
	if err != nil {
		logs.New().SetAdditionalInfo("err", err.Error()).Error("设置防火墙发送请求错误", err)
	}
	logs.New().
		SetAdditionalInfo("header", header).
		SetAdditionalInfo("data", data).
		SetAdditionalInfo("url", global.ApiUrl["set_firewall"]).
		SetAdditionalInfo("resp", resp.String()).Info("开通主机设置防火墙记录")
}

// SetIPBlackWhiteListRule 设置ip黑白名单
func SetIPBlackWhiteListRule(userId, UserName, vpsId string, ipList []string) error {
	// TODO: 要重新实现，不一定是调用这个API

	header := map[string]string{
		"User-Agent":  "api-landui-lan",
		"X-RequestAU": fmt.Sprintf("%s|\t|%s", userId, UserName),
	}
	data := map[string]string{
		"vpsname":           vpsId,
		"product_type_name": "rds_mysql",
	}
	client := global.HttpClient
	resp, err := client.R().SetHeaders(header).SetFormData(data).Post(global.ApiUrl["set_ip_black_white_list_rule"])
	if err != nil {
		logs.New().SetAdditionalInfo("err", err.Error()).Error("设置黑白名单发送请求错误", err)
		return fmt.Errorf("设置黑白名单发送请求错误: %s", err.Error())
	}
	logs.New().
		SetAdditionalInfo("header", header).
		SetAdditionalInfo("data", data).
		SetAdditionalInfo("url", global.ApiUrl["set_ip_black_white_list_rule"]).
		SetAdditionalInfo("resp", resp.String()).Info("设置黑白名单记录")

	if resp.StatusCode() != 200 {
		return fmt.Errorf("设置黑白名单失败: %s", resp.String())
	}

	return nil
}

func DeleteIPBlackWhiteListRule(userId, UserName, vpsId string, ipList []string) error {
	// TODO: 要重新实现，不一定是调用这个API

	header := map[string]string{
		"User-Agent":  "api-landui-lan",
		"X-RequestAU": fmt.Sprintf("%s|\t|%s", userId, UserName),
	}
	data := map[string]string{
		"vpsname":           vpsId,
		"product_type_name": "rds_mysql",
	}
	client := global.HttpClient
	resp, err := client.R().SetHeaders(header).SetFormData(data).Post(global.ApiUrl["delete_ip_black_white_list_rule"])
	if err != nil {
		logs.New().SetAdditionalInfo("err", err.Error()).Error("删除黑白名单发送请求错误", err)
		return fmt.Errorf("删除黑白名单发送请求错误: %s", err.Error())
	}
	logs.New().
		SetAdditionalInfo("header", header).
		SetAdditionalInfo("data", data).
		SetAdditionalInfo("url", global.ApiUrl["delete_ip_black_white_list_rule"]).
		SetAdditionalInfo("resp", resp.String()).Info("删除黑白名单记录")

	if resp.StatusCode() != 200 {
		return fmt.Errorf("删除黑白名单失败: %s", resp.String())
	}

	return nil
}
