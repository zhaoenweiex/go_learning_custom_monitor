package controllers

import (
	"github.com/astaxie/beego"
	"io/ioutil"
	"encoding/json"
	"net/http"
	"fmt"
	"net"
	"time"
)

type MonitorController struct {
	beego.Controller
}
type ServiceStatus struct {
	Name   string
	Status string
	Message string
}
type ServiceInfo struct {
	Name string
	URL  string
}
type JsonStruct struct {
}

func NewJsonStruct() *JsonStruct {
	return &JsonStruct{}
}

func (jst *JsonStruct) Load(filename string, v interface{}) {
	//ReadFile函数会读取文件的全部内容，并将结果以[]byte类型返回
	data, err := ioutil.ReadFile(filename)
		//读取的数据为json格式，需要进行解码
	err = json.Unmarshal(data, v)
	if err != nil {
		return
	}
}

func (this *MonitorController) Get() {
	//读取配置文件
	JsonParse := NewJsonStruct()
	var serviceInfoArray []ServiceInfo
	JsonParse.Load("C:\\Users\\zew\\go\\src\\goResearch\\conf\\config.json", &serviceInfoArray)
	//依次访问URL检查返回结果，确定状态
	var resultArray []ServiceStatus
	var c *http.Client = &http.Client{

		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				c, err := net.DialTimeout(netw, addr, time.Second*1)
				if err != nil {
					fmt.Println("dail timeout", err)
					return nil, err
				}
				return c, nil

			},
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second * 1,
		},
	}
	for _,serviceInfo := range serviceInfoArray {
		fmt.Println(serviceInfo)
		url:="http://"+serviceInfo.URL
		response, err := c.Get(url)
		if err!=nil {
			resultArray = append(resultArray,
				[]ServiceStatus{{serviceInfo.Name, "danger","连接错误"}}...)
		}else {
			if response.StatusCode==200{
				resultArray = append(resultArray,
					[]ServiceStatus{{serviceInfo.Name, "normal",""}}...)
			}else{
				body,_ := ioutil.ReadAll(response.Body)
				resultArray = append(resultArray,
					[]ServiceStatus{{serviceInfo.Name, "danger","错误码"+response.Status+"错误消息"+string(body)}}...)
			}
		}
	}
	this.Data["json"] = resultArray
	this.ServeJSON()
}
