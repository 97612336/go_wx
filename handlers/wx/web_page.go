package wx

import (
	"go_wx/util"
	"net/http"
)

func Go_to_page(w http.ResponseWriter, r *http.Request) {
	home_path:=util.Get_home_path()
	html_path:=home_path+"/templates/index.html"

	util.Render_template(w,html_path,nil)
}
