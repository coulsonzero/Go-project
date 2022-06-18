// static-files.go
package main

import (
	"errors"
	"net/http"
	"os"
	"path/filepath"
)

/**
创建assets多及目录，写入styles.css文件，再curl读取css文件内容
*/

func main() {
	dir_name, file_name := "assets/css", "styles.css"

	createDir(dir_name)
	// defer os.RemoveAll(strings.Split(dir_name, "/")[0])
	writeCssFile(dir_name, file_name)

	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", nil)

}

func createDir(dir_name string) error {
	err := os.MkdirAll(dir_name, os.ModePerm)
	if err != nil {
		return errors.New("创建目录失败！")
	}
	return nil
}

func writeCssFile(dir_name string, file_name string) {
	s := `
* {
	margin: 0;
	padding: 0;
	box-sizing: border-box;
}

body {
    background-color: black;
}
`
	path := filepath.Join(dir_name, file_name)
	err := os.WriteFile(path, []byte(s), os.ModePerm)
	if err != nil {
		panic(err)
	}
}

/*
$ curl -s http://localhost:8080/static/css/styles.css

* {
        margin: 0;
        padding: 0;
        box-sizing: border-box;
}

body {
    background-color: black;
}

*/
