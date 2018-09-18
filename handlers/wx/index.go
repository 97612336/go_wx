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
	old_wx_toekn, one_err := util.Get_redis("wx_token2")
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
