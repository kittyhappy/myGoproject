package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

//控制器

//页面

//IndexView 主页面
func IndexView(w http.ResponseWriter, r *http.Request) {
	html := loadHTML("./views/index.html")
	w.Write(html)
}

//UploadView 上传页面
func UploadView(w http.ResponseWriter, r *http.Request) {
	html := loadHTML("./views/upload.html")
	w.Write(html)
}

//APIUpLoad 图片上传
func APIUpLoad(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	f, h, err := r.FormFile("file")
	if err != nil {
		io.WriteString(w, "上传错误")
		return
	}
	t := h.Header.Get("Content-Type")
	if !strings.Contains(t, "image") {
		io.WriteString(w, "文件类型错误")
		return
	}
	os.Mkdir("./static", 0666)
	now := time.Now()
	name := now.Format("2006-01-02150405") + path.Ext(h.Filename) //获取后缀名
	out, err := os.Create("./static/" + name)
	if err != nil {
		io.WriteString(w, "文件创建错误")
		return
	}
	io.Copy(out, f)
	out.Close()
	f.Close()
	mod := Info{
		Name: h.Filename,
		Path: "/static/" + name,
		Note: r.Form.Get("note"),
		Unix: now.Unix(),
	}
	log.Println(InfoAdd(&mod))
	http.Redirect(w, r, "/list", 302)
}

//ListView 列表页面
func ListView(w http.ResponseWriter, r *http.Request) {
	html := loadHTML("./views/list.html")
	w.Write(html)
}

//APIList 相册列表接口
func APIList(w http.ResponseWriter, r *http.Request) {
	mods,_ := InfoList()
	buf,_ := json.Marshal(mods)
	w.Write(buf)
}

//DetailView 详细页面
func DetailView(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	idStr := r.Form.Get("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	mod, _ := InfoGet(id)
	html := loadHTML("./views/detail.html")
	date := time.Unix(mod.Unix, 0).Format("2006-01-02 15:04:05")
	html = bytes.Replace(html, []byte("@src"), []byte(mod.Path), 1)
	html = bytes.Replace(html, []byte("@note"), []byte(mod.Note), 1)
	html = bytes.Replace(html, []byte("@unix"), []byte(date), 1)
	w.Write(html)
}

//APIDrop 删除
func APIDrop(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	idStr := r.Form.Get("id")
	id,_ :=strconv.ParseInt(idStr,10,64)
	err :=InfoDrop(id)
	if err !=nil {
		io.WriteString(w,"删除失败")
		return	
	}
	io.WriteString(w,"删除成功")
	return
}

//loadHTML 加载html文件
func loadHTML(name string) []byte {
	buf, err := ioutil.ReadFile(name)
	if err != nil {
		return []byte("")
	}
	return buf
}
