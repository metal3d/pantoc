package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"gopkg.in/yaml.v2"
)

func appendToBuff(buff []string, filename string) []string {

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	return append(buff, string(content))
}

func setTitle(buff []string, title string, level int) []string {
	return append(buff, fmt.Sprintf("%s %s\n", strings.Repeat("#", level), title))
}

func printLevel(v interface{}, level int) []string {

	buff := make([]string, 0)

	switch v := v.(type) {
	case string:
		buff = appendToBuff(buff, v)
	case map[string]interface{}:
		for title, content := range v {
			buff = setTitle(buff, title, level)

			if val, ok := content.(string); ok {
				buff = appendToBuff(buff, val)
			} else {
				buff = append(buff, printLevel(content, level)...)
			}
		}
	case []map[string]interface{}:
		for _, val := range v {
			buff = append(buff, printLevel(val, level+1)...)
		}
	case []interface{}:
		for _, val := range v {
			buff = append(buff, printLevel(val, level+1)...)
		}
	case map[interface{}]interface{}:
		for key, val := range v {
			buff = setTitle(buff, key.(string), level)
			buff = append(buff, printLevel(val, level)...)
		}
	}
	return buff
}

func main() {
	var toc string
	flag.StringVar(&toc, "tocfile", "toc.yaml", "YAML file that describes table of contents")
	flag.Parse()

	content, _ := ioutil.ReadFile(toc)
	v := make([]map[string]interface{}, 0)
	yaml.Unmarshal(content, &v)

	doc := printLevel(v, 0)
	fmt.Println(strings.Join(doc, "\n"))
}
