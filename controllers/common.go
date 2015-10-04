package controllers

import (
	"io/ioutil"
	"net/http"
)

func Page_not_found(rw http.ResponseWriter, r *http.Request) {
	f, _ := ioutil.ReadFile("./static/html/error.html")
	rw.Write(f)
}

func StaticPageRender(fname string, rw http.ResponseWriter) {
	f, _ := ioutil.ReadFile(fname)
	rw.Write(f)
}
