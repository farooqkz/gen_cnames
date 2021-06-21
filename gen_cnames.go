package main

import (
	"flag"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/cbroglie/mustache"
)

var (
	maxDomains int = 16
	users      map[string][]string
	validate   = regexp.MustCompile(`[a-zA-Z0-9\-\.]+\.[a-zA-Z]+`)
)

func init() {
	flag.IntVar(&maxDomains, "max-domains", 16, "Max domains")
	flag.Parse()
}

func main() {
	dirs, _ := ioutil.ReadDir("/home/")
	for _, user := range dirs {
		cname := "/home/" + user.Name() + "/.CNAME"
		var domains []string
		if b, err := os.ReadFile(cname); err == nil {
			for _, v := range strings.Split(string(b), `\n`) {
				if validate.MatchString(v) {
					domains = append(domains, v)
				}
			}
			domains = domains[:(maxDomains - 1)]
			users[user.Name()] = domains
		}
	}

	http, _ := mustache.RenderFile("nginx.mustache", users)
	_ = ioutil.WriteFile("/tmp/result5.http", []byte(http), 0644)

	https, _ := mustache.RenderFile("nginx.ssl.mustache", users)
	_ = ioutil.WriteFile("/tmp/result5.https", []byte(https), 0644)
}
