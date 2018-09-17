package wx

import (
	"go_wx/util"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	data["code"] = 200
	util.Return_json(w, data)
}
