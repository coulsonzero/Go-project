package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

func main() {
	dotSql, _ := LoadFromFile("./src/package/os/sql/query.sql")
	query, _ := dotSql.LookupQuery("create-user")
	fmt.Println(query) // INSERT INTO users (name, email) VALUES(?, ?)

}

type DotSql struct {
	queries map[string]string
}

type Scanner struct {
	line    string
	queries map[string]string
	current string
}

type stateFn func(*Scanner) stateFn

// var matchTag string = "name"

// func (d *DotSql) Get(match string) {
// 	matchTag = match
// }

// LookupQuery 根据注释读取sql文件中的sql语句
func (d DotSql) LookupQuery(name string) (query string, err error) {
	query, ok := d.queries[name]
	if !ok {
		err = fmt.Errorf("dotsql: '%s' could not be found", name)
	}

	return
}

// LoadFromFile 加载sql文件
func LoadFromFile(sqlFile string) (*DotSql, error) {
	f, err := os.Open(sqlFile)
	if err != nil {
		return nil, err
	}

	return load(f)
}

func load(r io.Reader) (*DotSql, error) {
	scanner := &Scanner{}
	queries := scanner.run(bufio.NewScanner(r))

	dotsql := &DotSql{
		queries: queries,
	}

	return dotsql, nil
}

func (s *Scanner) run(io *bufio.Scanner) map[string]string {
	s.queries = make(map[string]string)

	for state := initialState; io.Scan(); {
		s.line = io.Text()
		state = state(s)
	}

	return s.queries
}

func initialState(s *Scanner) stateFn {
	if tag := getTag(s.line); len(tag) > 0 {
		s.current = tag
		return queryState
	}
	return initialState
}

func getTag(line string) string {
	re := regexp.MustCompile("^\\s*--\\s*name:\\s*(\\S+)")
	matches := re.FindStringSubmatch(line)
	if matches == nil {
		return ""
	}
	return matches[1]
}

func queryState(s *Scanner) stateFn {
	if tag := getTag(s.line); len(tag) > 0 {
		s.current = tag
	} else {
		s.appendQueryLine()
	}
	return queryState
}

func (s *Scanner) appendQueryLine() {
	current := s.queries[s.current]
	line := strings.Trim(s.line, " \t")
	if len(line) == 0 {
		return
	}

	if len(current) > 0 {
		current = current + "\n"
	}

	current = current + line
	s.queries[s.current] = current
}
