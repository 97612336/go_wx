package wx

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"go_wx/models"
	"go_wx/util"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
)

// 微信公众号的接口
func Index(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024 * 1024 * 3)
	if r.Method == "GET" {
		signature := util.Get_argument(r, "signature")
		echostr := util.Get_argument(r, "echostr")
		timestamp := util.Get_argument(r, "timestamp")
		nonce := util.Get_argument(r, "nonce")
		is_equal := Check_wx_signature(signature, timestamp, nonce)
		if is_equal {
			w.Write([]byte(echostr))
		} else {
			fmt.Println("验证出错！")
		}
	} else if r.Method == "POST" {
		// 读取请求的body
		one_envent := util.Read_xml_file(r)
		// 获取请求的类型
		msg_type := one_envent.Msg_type
		// 获取事件类型
		event := one_envent.Event
		// 获取事件key
		event_key := one_envent.Event_key
		current_time := util.Get_current_time_num()
		fmt.Println("+++++++++++++++++++++++++")
		if msg_type == "event" {
			// 如果是点击事件，就执行点击事件的操作
			if event == "CLICK" {
				if event_key == "one" {
					content := "这是在测试今日话语！"
					xml_str := util.Response_xml(one_envent.From_user_name, one_envent.To_user_name, current_time, content)
					w.Write([]byte(xml_str))
					return
				} else if event_key == "two_two" {
					content := "这是在测试点赞！"
					xml_str := util.Response_xml(one_envent.From_user_name, one_envent.To_user_name, current_time, content)
					w.Write([]byte(xml_str))
					return
				}
			}
		} else if msg_type == "text" {
			// 得到用户发送的消息
			user_words := one_envent.Content
			content := util.Talk_with_robot(user_words, "123123")
			xml_str := util.Response_xml(one_envent.From_user_name, one_envent.To_user_name, current_time, content)
			fmt.Println(xml_str)
			w.Write([]byte(xml_str))
			return
		}
	}
}

// 验证微信公众号的signature
func Check_wx_signature(signature string, timestamp string, nonce string) bool {
	wx_conf := util.Get_wx_conf()
	//获取token
	token := wx_conf.App_token
	var s string
	// 把要进行对比的参数进行排序
	s_list := []string{token, timestamp, nonce}
	sort.Strings(s_list)
	s = s_list[0] + s_list[1] + s_list[2]
	//运行go中的hash算法
	h := sha1.New()
	_, err := h.Write([]byte(s))
	util.CheckErr(err)
	res := h.Sum(nil)
	// 得到hash算法计算得到的字符串
	res_str := hex.EncodeToString(res)
	// 对比signature和hash计算的结果进行对比
	is_equal := res_str == signature
	return is_equal
}

// 获取access_token的方法
func Get_access_token() string {
	old_wx_toekn, one_err := util.Get_redis("wx_token")
	if one_err == nil {
		return old_wx_toekn
	}
	// redis中如果没有，就执行获取操作
	wx_conf := util.Get_wx_conf()
	appid := wx_conf.Appid
	appsecret := wx_conf.App_secret
	url := "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=" + appid + "&secret=" + appsecret
	res, err := http.Get(url)
	util.CheckErr(err)
	// 读取得到的请求
	res_info, res_err := ioutil.ReadAll(res.Body)
	util.CheckErr(res_err)
	res_str := string(res_info)
	var one_access_token models.Acess_token
	util.Json_to_object(res_str, &one_access_token)
	wx_token := one_access_token.Access_token
	util.Set_redis("wx_token", wx_token, "6000")
	return wx_token
}

//自定义菜单的方法
func Wx_menu() {
	access_token := Get_access_token()
	url := "https://api.weixin.qq.com/cgi-bin/menu/create?access_token=" + access_token
	json_str := `
	 {
     "button":[
     {    
          "type":"click",
          "name":"今日话语",
          "key":"one"
      },
      {
           "name":"菜单",
           "sub_button":[
           {    
               "type":"view",
               "name":"搜索",
               "url":"http://www.soso.com/"
            },
            {
               "type":"click",
               "name":"点赞",
               "key":"two_two"
            }]
       }]
 }
`

	resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(json_str))
	util.CheckErr(err)
	defer resp.Body.Close()
	res_body, body_err := ioutil.ReadAll(resp.Body)
	util.CheckErr(body_err)
	var err_meg models.Error_message
	util.Json_to_object(string(res_body), &err_meg)
	if err_meg.Err_code != 0 {
		fmt.Println("自定义菜单出错")
	}
}
