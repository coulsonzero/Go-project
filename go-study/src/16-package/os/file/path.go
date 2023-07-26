package main

import (
	"fmt"
	"path"
)

/**
 * path.Join() : 拼接路径
 * path.Dir()  : 返回除末尾元素外的路径名，即路径的目录
 * path.Base() : 返回路径末尾元素
 * path.IsAbs(): 判断是否为目录
 */

func main() {
	p := path.Join("dir1", "dir2", "filename") // dir1/dir2/filename
	fmt.Println("p:", p)

	fmt.Println(path.Join("dir1//", "filename"))       // dir1/filename
	fmt.Println(path.Join("dir1/../dir1", "filename")) // dir1/filename
	fmt.Println(path.Join("dir1/subDir2", "filename")) // dir1/subDir2/filename

	fmt.Println("Dir(p):", path.Dir(p))   // dir1/dir
	fmt.Println("Base(p):", path.Base(p)) // filename

	fmt.Println(path.IsAbs("dir/file"))  // false
	fmt.Println(path.IsAbs("/dir/file")) // true
	
}
