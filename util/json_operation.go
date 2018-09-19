package util

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"go_wx/models"
	"html/template"
	"io/ioutil"
	"net/http"
)

//将类型转化为字符串json
func Get_json_string(m interface{}) string {
	res, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	res_str := string(res)
	return res_str
}

//将字符串json转化为类型
func Json_to_object(json_str string, i interface{}) {
	json.Unmarshal([]byte(json_str), i)
}

//在web中返回json字符串
func Return_json(w http.ResponseWriter, i interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json_str := Get_json_string(i)
	w.Write([]byte(json_str))
}

//返回跨域的json
func Return_jsonp(w http.ResponseWriter, i interface{}) {
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json
	json_str := Get_json_string(i)
	w.Write([]byte(json_str))
}

//渲染模板的封装
func Render_template(w http.ResponseWriter, html_path string, data interface{}) {
	tmpl, _ := template.ParseFiles(html_path)
	tmpl.Execute(w, data)
	return
}

// 读取xml的方法
func Read_xml_file(r *http.Request) models.WX_event {
	data, err := ioutil.ReadAll(r.Body)
	var one_model models.WX_event
	CheckErr(err)
	fmt.Println(string(data))
	xml.Unmarshal(data, &one_model)
	return one_model
}

// 返回xml的方法
func Response_xml(tousername string, from_username string, create_time int, content string) string {
	xml_str := `
		<xml> 
		<ToUserName><![CDATA[%s]]></ToUserName> 
		<FromUserName><![CDATA[%s]]></FromUserName> 
		<CreateTime>%d</CreateTime> 
		<MsgType><![CDATA[text]]></MsgType> 
		<Content><![CDATA[%s]]></Content>
		</xml>
`
	new_xml_str := fmt.Sprintf(xml_str, tousername, from_username, create_time, content)
	return new_xml_str
}
