package aliyun

import (
	"log"
	"os"
)
//上传主函数
//filepath为文件路径,UploadUrl为上传地址
func Upload(filepath string)  {
	//获取文件大小,sha1,part数量
	if GetFileInfo() == false{
		os.Exit(0)
	}
	var a Refresh
	a.TokenType = "refresh_token"
	a.RefreshToken = RefreshToken.RefreshToken
	info :=GetInfo(filepath)
	c :=GetUploadUrl(&a)
	//判断是否可以秒传(服务器已有文件)
	if c.RapidUpload == true{
		log.Println("文件可以秒传...")
		Save(c)
		return
	}else {
		file,fileerr :=os.Open(filepath)
		defer file.Close()
		if fileerr != nil {
			log.Printf("打开文件失败:%v",fileerr)
			return
		}
		//声明一个5M切片
		var tmp = make([]byte,Size)
		//循环切片
		for i := 0; i <info.Part ; i++ {
			n,err :=file.Read(tmp)
			log.Printf("切片数量:%d\n",info.Part)
			if err != nil {
				log.Printf("文件切片错误:%v\n",err)
				return
			}
			log.Printf("这是第%d块\n",i+1)
			//遍历获取到的上传地址,传入part上传函数
			UploadPart(tmp[:n],c.PartInfoList[i]["upload_url"])
		}
		//保存文件
		Save(c)
	}
}
//分片上传函数
// n为次数,tmp为数据,PutUrl为上传地址
func UploadPart(tmp[]byte,PutUrl string )  {
		//循环put发送数据
	stutus :=PutNet(tmp,PutUrl)
	//判断put请求后的状态码
	if stutus !=0{
		log.Println("put发送数据成功....")
	}else {
		log.Println("put发送数据失败....")
	}
}
//保存函数
func Save(a *CreateData)  {
	Da :=&SvData{
		DriveID:  a.DriveID,
		FileID:   a.FileID,
		UploadID: a.UploadID,
	}
	data :=PostNet(Da,"https://api.aliyundrive.com/v2/file/complete",Authorization)
	log.Println(string(data))
}
