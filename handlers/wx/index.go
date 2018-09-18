package wx

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"go_wx/util"
	"net/http"
	"sort"
)

func Index(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024 * 1024 * 3)
	if r.Method == "GET" {
		signature:=util.Get_argument(r,"signature")
		echostr:=util.Get_argument(r,"echostr")
		timestamp:=util.Get_argument(r,"timestamp")
		nonce:=util.Get_argument(r,"nonce")
		is_equal:=Check_wx_signature(signature,timestamp,nonce)
		if is_equal{
			w.Write([]byte(echostr))
		}else{
			fmt.Println("验证出错！")
		}

		}
}


// 验证微信公众号的signature
func Check_wx_signature(signature string,timestamp string,nonce string) bool{
	wx_conf:=util.Get_wx_conf()
	//获取token
	token:=wx_conf.App_token
	fmt.Println("token是：",token)
	var s string
	//s=timestamp+nonce+token
	s_list:=[]string{signature,timestamp,nonce}
	fmt.Println(s_list)
	sort.Strings(s_list)
	fmt.Println(s_list)
	s=s_list[0]+s_list[1]+s_list[2]
	fmt.Println(s)
	//运行go中的hash算法
	h:=sha1.New()
	_,err:=h.Write([]byte(s))
	util.CheckErr(err)
	res:=h.Sum(nil)
	// 得到hash算法计算得到的字符串
	res_str:=hex.EncodeToString(res)
	fmt.Println(res_str)
	fmt.Println(signature)
	// 对比signature和hash计算的结果进行对比
	is_equal:=res_str==signature
	return is_equal
	}