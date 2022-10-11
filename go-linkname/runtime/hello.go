package runtime

import (
	"fmt"

	_ "unsafe" // for go:linkname
)

//go:linkname link go-linkname/hello.Link
func link() {
	fmt.Println("hello world...")
}
