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

// 微信错误信息返回的类
type Error_message struct {
	Err_code int
	err_msg string
}

// 定义微信动作类
type WX_event struct {
	To_user_name string   `xml:"ToUserName"`
	From_user_name string	`xml:"FromUserName"`
	Create_time int			`xml:"CreateTime"`
	Msg_type string			`xml:"MsgType"`
	Event string			`xml:"Event"`
	Event_key string		`xml:"EventKey"`
	Content string			`xml:"Content"`
}


// 定义聊天机器人的配置文件类
type Robot_conf struct {
	R1 string
	R2 string
}

type One_emotion struct {
	A int
	D int
	EmotionId int
	P int
}

type Emotions struct {
	RobotEmotion One_emotion
	UserEmotion One_emotion
}

type Intent struct {
	ActionName string
	Code int
	IntentName string
}

type One_value struct {
	Text string
}

type One_result struct {
	Group_type int
	ResultType int
	Values One_value
}

// 接口最后返回的格式
type Result struct {
	Emotion Emotions
	Intent Intent
	Results []One_result
}