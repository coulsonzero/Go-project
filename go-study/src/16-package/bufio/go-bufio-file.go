package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// 逐行读取文件内容

func main() {
	zeroSql, _ := Load("./src/package/bufio/query.sql")
	// zeroSql.SetTag("type")

	// zeroSql.PrintResult()
	// println(zeroSql.matchTag)
	// query, _ := zeroSql.LookupQuery("create-users-table")
	// query, _ := zeroSql.LookupQuery("create-user")
	// query, _ := zeroSql.LookupQuery("find-users-by-email")
	// query, _ := zeroSql.LookupQuery("drop-users-table")

	// query, err := zeroSql.LookupQueryAny("drop-users-table")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(query)

	// println(zeroSql.Get("type").QueryAny("find-one-user-by-email"))
	// println(zeroSql.Get("type").QueryAny("find-users-by-email"))
	// println(zeroSql.Get("tag").QueryAny("drop-users-table"))
	query, _ := zeroSql.Get("name").QueryAny("create-users-table")
	fmt.Println(query)
}

type ZeroSql struct {
	queries  map[string]string
	matchTag string
}

type QuerySql struct {
	queries map[string]string
}

const (
	defaultMatch string = "name"
	sep          string = " >> "
)

func getByTag(line string) string {
	re := regexp.MustCompile("^\\s*--\\s*\\W?(\\w+)\\W?\\s*(\\S+)")
	// re := regexp.MustCompile("^\\s*--\\s*name:\\s*(\\S+)")

	matches := re.FindStringSubmatch(line)
	if matches == nil {
		return ""
	}
	return matches[1] + sep + matches[2]
}

func Load(sqlFile string) (*ZeroSql, error) {
	file, err := os.Open(sqlFile)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	var title []string
	var context []string
	current := ""
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "--") {
			if len(current) != 0 {
				current = strings.TrimSpace(current)
				context = append(context, current)
				current = ""
			}
			// title = append(title, strings.Split(line, ": ")[1])
			// title = append(title, line)
			title = append(title, getByTag(line))
		} else {
			if len(current) > 0 {
				current = current + "\n"
			}
			current = current + line
		}
	}
	current = strings.TrimSpace(current)
	context = append(context, current)
	current = ""

	m := make(map[string]string)
	for i := 0; i < len(title); i++ {
		m[title[i]] = context[i]
	}

	res := &ZeroSql{
		queries:  m,
		matchTag: defaultMatch,
	}

	return res, nil
}

func (d *ZeroSql) LookupQuery(name string) (query string, err error) {
	key := d.matchTag + sep + name
	query, ok := d.queries[key]
	if !ok {
		err = fmt.Errorf("sql: '%s' could not be found", name)
	}

	return
}

func (d *ZeroSql) LookupQueryAny(name string) (query string, err error) {
	for k, v := range d.queries {
		key := strings.Split(k, sep)[1]
		if key == name {
			return v, nil
		}
	}
	return "", fmt.Errorf("sql: '%s' could not be found", name)
}

func (d *ZeroSql) SetTag(match string) {
	d.matchTag = match
	m := make(map[string]string)
	for k, v := range d.queries {
		if strings.Split(k, sep)[0] == match {
			m[k] = v
		}
	}

	d.queries = m
}

func (d ZeroSql) Get(match string) *QuerySql {
	d.matchTag = match
	m := make(map[string]string)
	for k, v := range d.queries {
		if strings.Split(k, sep)[0] == match {
			m[k] = v
		}
	}

	d.queries = m
	return &QuerySql{queries: m}
}

func (d *QuerySql) QueryAny(name string) (query string, err error) {
	for k, v := range d.queries {
		key := strings.Split(k, sep)[1]
		if key == name {
			return v, nil
		}
	}
	return "", fmt.Errorf("sql: '%s' could not be found", name)
}

func (d *ZeroSql) PrintResult() {
	for k, v := range d.queries {
		fmt.Printf("key: %s\nvalue: %s\n", k, v)
	}
}
