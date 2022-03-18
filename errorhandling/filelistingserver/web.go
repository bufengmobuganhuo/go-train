package main

import (
	"net/http"
	"os"

	"mengyu.com/gotrain/errorhandling/filelistingserver/filelisting"
)

// 定义成一个类型，方便下面的方法写入参
type appHandler func(w http.ResponseWriter, r *http.Request) error

func errorWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := handler(w, r)
		// 错误处理
		if err != nil {
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

func main() {
	// http://localhost:8888/list/fib.txt
	http.HandleFunc("/list/", errorWrapper(filelisting.HandleFileList))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
