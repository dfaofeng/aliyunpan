package aliyun

import (
	"bytes"
	"github.com/go-resty/resty/v2"
	"log"
	"net/http"
	"os"
)
//post通用网络请求
//body为结构体,PostUrl为地址
func PostNet(body interface{},PostUrl string,ck...string)[]byte  {
	k :=""
	client :=resty.New()
	if ck !=nil{
		for _, s := range ck {
			k = s
		}
	}
	resp,err :=client.R().
		SetHeader("authorization",k).
		SetBody(body).
		Post(PostUrl)
	if err != nil {
		log.Printf("错误:%v\n",err)
	}
	return resp.Body()
}
//put通用网络请求
//body为请求数据的[]byte,PutUrl为地址,返回为[]byte
func PutNet(body []byte,PutUrl string) int {
	client, _ :=http.NewRequest("PUT",PutUrl,bytes.NewReader(body))
	respoest,e :=http.DefaultClient.Do(client)
	st :=respoest.StatusCode
	defer respoest.Body.Close()
	if e != nil {
		os.Exit(0)
		return 0
	}
	return st
}
