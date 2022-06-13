package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	u, _ := url.Parse(s)

	fmt.Println(u.Scheme)          // postgres
	fmt.Println(u.User)            // user:pass
	fmt.Println(u.User.Username()) // user
	p, _ := u.User.Password()      // pass
	fmt.Println(p)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host, port) // host.com 5432

	fmt.Println(u.Path)     // /path
	fmt.Println(u.RawQuery) // k=v
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)          // map[k:[v]]
	fmt.Println(u.Fragment) // f

}
