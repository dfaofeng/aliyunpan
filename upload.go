package main

import (
	"aliyun/ut"
	"bufio"
	"io"
	"log"
	"os"
)

// Upload 上传主函数
//filepath为文件路径,UploadUrl为上传地址
func Upload() {
	//获取文件大小,sha1,part数量
	st,path :=GetFileInfo()
	if st == false{
		os.Exit(0)
	}
	a :=Refresh{
		RefreshToken: RefreshToken.RefreshToken,
		TokenType:   "refresh_token" ,
	}
	info :=GetInfo(path)
	c :=GetUploadUrl(&a)
	//判断是否可以秒传(服务器已有文件)
	if c.RapidUpload == true{
		log.Println("文件可以秒传...")
		Save(c)
		return
	}else {
		file,fileerr :=os.Open(path)
		if fileerr != nil {
			log.Printf("打开文件失败:%v",fileerr)
			return
		}
		log.Printf("开始上传文件:%s",file.Name())
		slip(file,info,c)
		//保存文件
		Save(c)
		defer file.Close()
	}
}

// UploadPart 分片上传函数
// n为次数,tmp为数据,PutUrl为上传地址
func UploadPart(tmp[]byte,PutUrl string)  {
		//循环put发送数据
	stutus := ut.PutNet(tmp,PutUrl)
	//判断put请求后的状态码
	if stutus !=0{
		log.Println("put发送数据成功....")
	}else {
		log.Println("put发送数据失败....")
	}
}

// Save 保存函数
func Save(a *CreateData)  {
	Da :=&SvData{
		DriveID:  a.DriveID,
		FileID:   a.FileID,
		UploadID: a.UploadID,
	}
	data,code := ut.PostNet(Da,"https://api.aliyundrive.com/v2/file/complete",Authorization)
	if code !=200{
		log.Printf("文件保存异常:%s",string(data))
		os.Exit(1)
	}
	log.Println(string(data))
}
func slip(file *os.File,info *FileInfo,c *CreateData)  {
	//声明一个5M切片
	var tmp = make([]byte,Size)
	//循环切片
	log.Println("开始文件切片")
	r :=bufio.NewReader(file)
	log.Printf("切片数量:%d\n",info.Part)
	for i := 0; i <info.Part ; i++ {
		n,err :=r.Read(tmp)
		if err == io.EOF {
			log.Printf("文件切片结束:%v\n",err)
			break
		}else {
			log.Printf("这是第%d块:总共%d块\n",i+1,info.Part)
			//遍历获取到的上传地址,传入part上传函数
			UploadPart(tmp[:n],c.PartInfoList[i]["upload_url"])
		}
	}
}
