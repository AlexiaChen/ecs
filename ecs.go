package ecs

import (
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"gitlab.landui.cn/gomod/logs"
	"sort"
)

const (
	CreateHostAPI                 = "/api/interimCreate"
	DeleteHostAPI                 = "/api/delVps"
	GetHostStatusAPI              = "/api/getCloudStatus"
	SetFirewallAPI                = "/lan/cloud/defaultAclRuleSet"
	SetFirewallIPBlackWhiteAPI    = ""
	DeleteFirewallIPBlackWhiteAPI = ""
)

type ECS struct {
	UserId   uint
	UserName string

	Id         uint
	RegionId   string //区域ID
	NodeId     uint32 //节点ID
	SystemId   string //系统版本id
	SystemType string //系统类型 Centos：7 Ubuntu：8

	Cpu       uint32 //cpu核心数
	Memory    uint32 //内存大小
	Bandwidth uint32 //带宽 后续换成iops
	HardDisks uint32 //系统盘大小
	Disks     uint32 //数据盘大小
	Defense   string //防御值
	Months    uint32 //有效期月数

	DiskType    string //磁盘类型
	IsIntranet  uint8  //是否为内网IP 1内网
	ProductType uint8  //产品类型  1 rds，2 redis

	APIUriPrefix  string
	APISignSecret string
}

type Response struct {
	Message string       `json:"message"`
	Code    int          `json:"code"`
	Info    ResponseInfo `json:"info"`
	Status  string       `json:"status"`
}
type ResponseInfo struct {
	Id       string `json:"id"`
	User     string `json:"user"`
	Pass     string `json:"pass"`
	Ip       string `json:"ip"`
	Password string `json:"VPSpassword"`
	Status   string `json:"status"`
}

func getHardDisks(disk, hardDisk uint32) ([]string, string) {
	diskMap := []string{
		fmt.Sprint(disk),
		//fmt.Sprint(hardDisk),
	}
	return diskMap, fmt.Sprint(hardDisk + disk)
}

func sorts(text string) string {
	var array []string
	for _, v := range text {
		array = append(array, string(v))
	}
	sort.Strings(array)
	newText := ""
	for _, v := range array {
		newText += v
	}
	return newText
}

func unauthorizedPost(url string) (*resty.Response, error) {
	client := resty.New()
	//resp, err := client.R().SetBody(body).Post(url)
	resp, err := client.R().Get(url)
	if err != nil {
		log := logs.New()
		log.SetAdditionalInfo("resp", resp).Error("创建主机的时候出现错误", err)
		return nil, err
	}
	return resp, err
}

func post(body map[string]interface{}, url string) (*resty.Response, error) {
	client := resty.New()
	resp, err := client.R().SetBody(body).Post(url)
	if err != nil {
		logs.New().Error("发送到php的请求失败", err)
		return nil, errors.New("创建云主机失败")
	}
	return resp, err
}
