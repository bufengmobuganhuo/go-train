package main

import (
	"log"
	"net/http"
	"os"

	"mengyu.com/gotrain/lang/errorhandling/filelistingserver/filelisting"
)

// 定义成一个类型，方便下面的方法写入参
type appHandler func(w http.ResponseWriter, r *http.Request) error

func errorWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				log.Printf("Panic:%s", r)
			}
		}()
		// 处理请求
		err := handler(w, r)
		// 错误处理
		if err != nil {
			log.Printf("Error occured handling request, err=%s", err.Error())
			// 判断自定义异常
			if userError, ok := err.(userError); ok {
				http.Error(w, userError.Message(), http.StatusBadRequest)
				return
			}
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(w, http.StatusText(code), code)
		}
	}
}

// 定义一个用户可见的error
type userError interface {
	error
	Message() string
}

func main() {
	// http://localhost:8888/*
	http.HandleFunc("/", errorWrapper(filelisting.HandleFileList))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
