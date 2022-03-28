package filelisting

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const Prefix = "/list/"

type userError string

// 实现error接口
func (e userError) Error() string {
	return e.Message()
}

// 实现userError接口
func (e userError) Message() string {
	return string(e)
}

func HandleFileList(w http.ResponseWriter, r *http.Request) error {
	if strings.Index(r.URL.Path, Prefix) != 0 {
		return userError("path must start with " + Prefix)
	}
	// 取文件的地址
	path := r.URL.Path[len("/list/"):]
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	w.Write(all)
	return nil
}
