package models


// 微信配置文件类
type Wx_Gzh struct {
	Appid string
	App_secret string
	App_token string
}


// 微信获取到的access_token类
type Acess_token struct {
	Access_token string
	Expires_in int
}