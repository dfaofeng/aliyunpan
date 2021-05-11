package main

import (
	"aliyun/ut"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"os"
)
var Size int
var Authorization string
var Fileinfo FileInfo

// GetInfo 获取文件相关信息
func GetInfo(filePath string) *FileInfo {
	// 打开文件
	fileObj,fileErr :=os.Open(filePath)
	if  fileErr!= nil {
		return nil
	}
	// 获取文件状态
	fileStat,_ :=fileObj.Stat()
	// 获取文件大小
	size :=fileStat.Size()
	file_flo :=float64(size)
	chunk_size :=float64(Size)
	//获取块数量
	chunk :=int(math.Ceil(file_flo/chunk_size))
	log.Println("正在计算文件的sha1....")
	h :=sha1.New()
	_,ioError :=io.Copy(h,fileObj)
	if ioError != nil {
		log.Println("io出错")
		return nil
	}
	//获取到的文件少了1k......
	fileSha1 :=fmt.Sprintf("%x",h.Sum(nil))
	log.Printf("sha1计算成功:%s",fileSha1)
	// 赋值给结构体
	Fileinfo.FileId = "root"
	Fileinfo.FileName = fileStat.Name()
	Fileinfo.FileSize = size
	Fileinfo.FileSha1 = fileSha1
	Fileinfo.Part = chunk
	defer fileObj.Close()
	return &Fileinfo
}

// GetToken 获取access_token
func GetToken(refresh *Refresh) *UserInfo {
	var postdata UserInfo
	data,code := ut.PostNet(refresh,"https://auth.aliyundrive.com/v2/account/token")
	//判断返回的状态码
	if code !=200{
		log.Printf("获取access_token异常:%s",string(data))
		os.Exit(1)
	}
	err :=json.Unmarshal(data,&postdata)
	if err != nil {
		log.Printf("json序列化失败:%v",err)
		return nil
	}
	Authorization = postdata.AccessToken
	return &postdata
}

// GetUploadUrl 获取上传地址
func GetUploadUrl(refresh *Refresh) *CreateData {
	//构建需要post的信息
	t :=GetToken(refresh)
	var a Create
	part :=make([]map[string]int,0)
	for i := 0; i < Fileinfo.Part; i++ {
		v :=map[string]int{
			"part_number":i+1,
		}
		part = append(part,v)
	}
	//分p切片数组
	a.PartInfoList = part
	//设备id
	a.DriveID = t.DefaultDriveID
	//文件名称
	a.Name = Fileinfo.FileName
	//文件sha1
	a.ContentHash = Fileinfo.FileSha1
	//文件1k的sha1
	a.ParentFileID = Fileinfo.FileId
	//类型
	a.CheckNameMode = "auto_rename"
	//加密类型
	a.ContentHashName = "sha1"
	//文件类型
	a.Type = "file"
	//文件大小
	a.Size = int(Fileinfo.FileSize)
	dirfile,_ := ut.PostNet(a,"https://api.aliyundrive.com/v2/file/create",Authorization)
	//if code !=200{
	//	log.Printf("获取上传地址服务器地址异常:%s",string(dirfile))
	//	os.Exit(1)
	//}
	var v CreateData
	_=json.Unmarshal(dirfile, &v)
	return &v
}
