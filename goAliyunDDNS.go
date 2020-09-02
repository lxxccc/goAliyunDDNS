package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/denverdino/aliyungo/dns"
	"lxxccc.top/Library/ToolkitsGo/config/json"
)

// ConfigInfo 定义域名相关配置信息
type ConfigInfo struct {
	AccessKeyID     string
	AccessKeySecret string
	DomainName      string
	RR              string
}

func main() {
	// 读取配置文件信息
	config := new(ConfigInfo)
	json.ReadJSONConfigToStruct("config.json", &config)
	// 获取公网IP信息
	resp, err := http.Get("https://myexternalip.com/raw")
	if err != nil {
		log.Println("发生错误", err)
		return
	}
	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)
	publicIP := string(content)
	log.Println("当前公网IP：" + publicIP)

	// 连接阿里云服务器，获取DNS信息
	client := dns.NewClient(config.AccessKeyID, config.AccessKeySecret)
	client.SetDebug(false)
	domainInfo := new(dns.DescribeDomainRecordsArgs)
	domainInfo.DomainName = config.DomainName
	oldRecord, err := client.DescribeDomainRecords(domainInfo)
	if err != nil {
		log.Println("链接错误", err)
		return
	}

	var exsitRecordID string
	for _, record := range oldRecord.DomainRecords.Record {
		if record.DomainName == config.DomainName && record.RR == config.RR {
			if record.Value == publicIP {
				fmt.Println("当前配置解析地址与公网IP相同，不需要修改。")
				return
			}
			exsitRecordID = record.RecordId
		}
	}

	if 0 < len(exsitRecordID) {
		// 有配置记录，则匹配配置文件，进行更新操作
		updateRecord := new(dns.UpdateDomainRecordArgs)
		updateRecord.RecordId = exsitRecordID
		updateRecord.RR = config.RR
		updateRecord.Value = publicIP
		updateRecord.Type = dns.ARecord
		rsp := new(dns.UpdateDomainRecordResponse)
		rsp, err := client.UpdateDomainRecord(updateRecord)
		if nil != err {
			fmt.Println("修改解析失败", err)
		} else {
			fmt.Println("修改解析成功", rsp)
		}
	} else {
		// 没有找到配置记录，那么就新增一个
		newRecord := new(dns.AddDomainRecordArgs)
		newRecord.DomainName = config.DomainName
		newRecord.RR = config.RR
		newRecord.Value = publicIP
		newRecord.Type = dns.ARecord
		rsp := new(dns.AddDomainRecordResponse)
		rsp, err = client.AddDomainRecord(newRecord)
		if nil != err {
			fmt.Println("添加DNS解析失败", err)
		} else {
			fmt.Println("添加DNS解析成功", rsp)
		}
	}
}
