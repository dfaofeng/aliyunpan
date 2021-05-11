package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"time"
)
var RefreshToken Info
func GetFileInfo() (bool,string) {
	_,fileErr :=os.Lstat("config.yaml")
	if os.IsNotExist(fileErr) == true {
		mkdir()
		log.Println("config.yaml不存在,已生成默认文件,请填写..")
		time.Sleep(5*time.Second)
		return false,""
	}
	fileObj,err :=ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Printf("配置文件config打开失败,请检查:%v",err)
		return false,""
	}
	yamlerr := yaml.Unmarshal(fileObj,&RefreshToken)
	if yamlerr != nil {
		log.Printf("yaml配置文件解析失败,请检查:%v",err)
		return false,""
	}
	if RefreshToken.RefreshToken =="" {
		log.Println("token为空,请检查...")
		return  false,""
	}
	//获取文件块大小
	if RefreshToken.Chunk == 0 {
		Size = 50*1024*1024
	}else {
		Size = RefreshToken.Chunk*1024*1024
	}
	return true,RefreshToken.Path
}
//创建配置文件
func mkdir()  {
	data,err :=yaml.Marshal(&RefreshToken)
	if err != nil {
		log.Printf("写入错误:%v",err)
		return
	}
	ioutil.WriteFile("config.yaml",data,0644)

}
