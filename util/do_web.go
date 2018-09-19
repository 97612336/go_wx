package util

import (
	"feidu/util"
	"fmt"
	"go_wx/models"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

//获取表单提交的值
func Get_argument(r *http.Request, key string, wantDefault ...string) string {
	argument := r.FormValue(key)
	if argument == "" {
		if wantDefault == nil {
			return ""
		}
		return wantDefault[0]
	}
	return argument
}

// 模拟POST请求
func My_post(url string, json_str string) string {
	resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(json_str))
	CheckErr(err)
	defer resp.Body.Close()
	res_body, body_err := ioutil.ReadAll(resp.Body)
	CheckErr(body_err)
	res_str := string(res_body)
	return res_str
}

// 获取当前时间戳
func Get_current_time_num() int {
	current_time := int(time.Now().Unix())
	return current_time
}

// 自动聊天
func Talk_with_robot(one_word string, from_user string) string {
	one_robot := Get_robot()
	send_json_str := `
	{
	"reqType":0,
    "perception": {
        "inputText": {
            "text": "%s"
        }
    },
    "userInfo": {
        "apiKey": "%s",
        "userId": "%s"
    }
}
`
	// 得到规范化的json
	format_send_json := fmt.Sprintf(send_json_str, one_word, one_robot.R1, from_user)
	url := "http://openapi.tuling123.com/openapi/api/v2"
	res_str := My_post(url, format_send_json)
	var one_result models.Result
	util.Json_to_object(res_str, &one_result)
	rep_word := one_result.Results[0].Values.Text
	return rep_word
}
