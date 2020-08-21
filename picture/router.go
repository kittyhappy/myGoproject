package main

import (
	"log"
	"net/http"
)

//Run 路由
func Run() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	http.HandleFunc("/", IndexView)
	http.HandleFunc("/upload", UploadView)
	http.HandleFunc("/list", ListView)
	http.HandleFunc("/detail", DetailView)

	http.HandleFunc("/api/upload", APIUpLoad)
	http.HandleFunc("/api/list", APIList)
	http.HandleFunc("/api/drop", APIDrop)
	log.Println("run 8080...")
	http.ListenAndServe(":8080", nil)
}
