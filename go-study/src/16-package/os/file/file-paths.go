package main

/**
 * filepath 已废弃 -> 请使用 path 代替
 */

import (
	"fmt"
	"path/filepath"
)

func main() {

	p := filepath.Join("dir1", "dir2", "filename") // dir1/dir2/filename
	fmt.Println("p:", p)

	fmt.Println(filepath.Join("dir1//", "filename"))       // dir1/filename
	fmt.Println(filepath.Join("dir1/../dir1", "filename")) // dir1/filename
	fmt.Println(filepath.Join("dir1/subDir2", "filename")) // dir1/subDir2/filename

	fmt.Println("Dir(p):", filepath.Dir(p))   // dir1/dir
	fmt.Println("Base(p):", filepath.Base(p)) // filename

	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))

	// sqlFile, err := filepath.Abs(path)
}
