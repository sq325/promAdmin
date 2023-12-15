package prom

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type PromSvc interface {
	Instances() []map[string]any
}

type promSvc struct {
	consulAddr string
}

func NewPromSvc(consulAddr string) PromSvc {
	return &promSvc{
		consulAddr: consulAddr,
	}
}

func (svc *promSvc) Instances() []map[string]any {
	if !strings.HasPrefix(svc.consulAddr, "http://") {
		svc.consulAddr = "http://" + svc.consulAddr
	}
	basepath := "/v1/catalog/"
	u, _ := url.Parse(svc.consulAddr)
	baseUrl := u.JoinPath(basepath)
	svcNames := svc.svcNames(baseUrl)

	var instances []map[string]any
	{
		for _, name := range svcNames {
			u := baseUrl.JoinPath("service", name)
			request, _ := http.NewRequest("GET", u.String(), nil)
			resp, err := http.DefaultClient.Do(request)
			if err != nil {
				fmt.Println("http.DefaultClient.Do is err:", err.Error())
			}
			defer resp.Body.Close()
			body, _ := io.ReadAll(resp.Body)
			var data []map[string]any
			err = json.Unmarshal(body, &data)
			if err != nil {
				fmt.Println("json.Unmarshal is err:", err.Error())
			}
			instances = append(instances, data...)
		}
	}

	return instances

}

func (svc *promSvc) svcNames(u *url.URL) []string {
	u = u.JoinPath("services")
	request, _ := http.NewRequest("GET", u.String(), nil)
	// 发送http request请求
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println("http.DefaultClient.Do is err:", err.Error())
	}
	defer response.Body.Close()
	// 读取返回body信息
	body, _ := io.ReadAll(response.Body)
	// 反序列化body为map[string]any
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("json.Unmarshal is err:", err.Error())
	}

	var svcNames []string
	for k := range data {
		if strings.HasPrefix(k, "prometheus:") {
			svcNames = append(svcNames, k)
		}
	}

	return svcNames
}
