package handlers

import (
	"go_wx/handlers/wx"
	"net/http"
)

func MyUrls() {
	http.HandleFunc("/", wx.Index)
	http.HandleFunc("/go/page/v1",wx.Go_to_page)
}
