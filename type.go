package aliyun

import (
	"os"
	"time"
)
// 文件相关的结构体
type FileInfo struct {
	FileSize int64
	FileSha1,FileName,FileId string
	Part int
	FileObj *os.File
}
//用户信息结构体
type UserInfo struct {
	AccessToken        string    `json:"access_token"`
	Avatar             string    `json:"avatar"`
	DefaultDriveID     string    `json:"default_drive_id"`
	DefaultSboxDriveID string    `json:"default_sbox_drive_id"`
	DeviceID           string    `json:"device_id"`
	ExpireTime         time.Time `json:"expire_time"`
	ExpiresIn          int       `json:"expires_in"`
	IsFirstLogin       bool      `json:"is_first_login"`
	NeedLink           bool      `json:"need_link"`
	NeedRpVerify       bool      `json:"need_rp_verify"`
	NickName           string    `json:"nick_name"`
	PinSetup           bool      `json:"pin_setup"`
	RefreshToken       string    `json:"refresh_token"`
	Role               string    `json:"role"`
	State              string    `json:"state"`
	Status             string    `json:"status"`
	TokenType          string    `json:"token_type"`
	UserData           struct {
		BackUpConfig struct {
			PhoBack struct {
				FolderID      string `json:"folder_id"`
				PhotoFolderID string `json:"photo_folder_id"`
				SubFolder     struct {
				} `json:"sub_folder"`
				VideoFolderID string `json:"video_folder_id"`
			} `json:"手机备份"`
		} `json:"back_up_config"`
		DingDingRobotURL  string `json:"DingDingRobotUrl"`
		DingDingRobotURL2 string `json:"ding_ding_robot_url"`
		EncourageDesc     string `json:"EncourageDesc"`
		EncourageDesc2    string `json:"encourage_desc"`
		FeedBackSwitch    bool   `json:"feed_back_switch"`
		FeedBackSwitch2   bool   `json:"FeedBackSwitch"`
		FollowingDesc     string `json:"following_desc"`
		FollowingDesc2    string `json:"FollowingDesc"`
	} `json:"user_data"`
	UserID   string `json:"user_id"`
	UserName string `json:"user_name"`
}
//refresh_token结构体
type Refresh struct {
	RefreshToken string `json:"refresh_token"`
	TokenType string `json:"grant_type"`
}
//创建文件结构体
type Create struct {
	CheckNameMode   string `json:"check_name_mode"`
	ContentHash     string `json:"content_hash"`
	ContentHashName string `json:"content_hash_name"`
	DriveID         string `json:"drive_id"`
	Name            string `json:"name"`
	ParentFileID    string `json:"parent_file_id"`
	PartInfoList    []map[string]int `json:"part_info_list"`
	Size int    `json:"size"`
	Type string `json:"type"`
}
//创建文件返回结构体
type CreateData struct {
	DomainID     string `json:"domain_id"`
	DriveID      string `json:"drive_id"`
	EncryptMode  string `json:"encrypt_mode"`
	FileID       string `json:"file_id"`
	FileName     string `json:"file_name"`
	Location     string `json:"location"`
	ParentFileID string `json:"parent_file_id"`
	PartInfoList []map[string]string`json:"part_info_list"`
	RapidUpload bool   `json:"rapid_upload"`
	Type        string `json:"type"`
	UploadID    string `json:"upload_id"`
}
//yaml解析结构体
type Info struct {
	RefreshToken string `yaml:"refresh_token"`
}
//保存文件file结构体
type SvData struct {
	DriveID  string `json:"drive_id"`
	FileID   string `json:"file_id"`
	UploadID string `json:"upload_id"`
}
