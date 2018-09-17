package handlers

import (
	"go_wx/handlers/wx"
	"net/http"
)

func MyUrls() {
	http.HandleFunc("/", wx.Index)
}
