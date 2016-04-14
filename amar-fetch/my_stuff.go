package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/daniel-fanjul-alcuten/amar-dashboard-backend"
)

var stockRegexp *regexp.Regexp = regexp.MustCompile("<tr><td><a class='link' href='" +
	"([^']+)'>([^<]+)</a><td>([\\d,]*)<td>([\\d,]*)<td>([\\d,]*)<td>([\\d,]*)<td>([\\d,]*)")

type MyStuffFetcher struct {
	Pid, Uid string
}

func (ms MyStuffFetcher) Fetch() (s string, err error) {
	var req *http.Request
	if req, err = http.NewRequest("GET", "http://amar.bornofsnails.net/man/my_stuff", nil); err != nil {
		return
	}
	req.AddCookie(&http.Cookie{Name: "pid", Value: ms.Pid})
	req.AddCookie(&http.Cookie{Name: "uid", Value: ms.Uid})
	c := &http.Client{}
	var resp *http.Response
	if resp, err = c.Do(req); err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		err = fmt.Errorf("Get http://amar.bornofsnails.net/man/my_stuff: %v", resp.Status)
		return
	}
	b := &bytes.Buffer{}
	if _, err = io.Copy(b, resp.Body); err != nil {
		return
	}
	s = b.String()
	return
}

func (ms MyStuffFetcher) Parse(t time.Time, s string) (page amar.MyStuffPage, err error) {
	page.Time = t
	page.Stuff = make(map[string]amar.MyStuff)
	for _, m := range stockRegexp.FindAllStringSubmatch(s, -1) {
		item := amar.MyStuff{Link: m[1], Name: m[2]}
		if item.Total, err = atoi(m[3]); err != nil {
			return
		}
		if item.Inventory, err = atoi(m[4]); err != nil {
			return
		}
		if item.House, err = atoi(m[5]); err != nil {
			return
		}
		if item.Shared, err = atoi(m[6]); err != nil {
			return
		}
		if item.Guild, err = atoi(m[7]); err != nil {
			return
		}
		page.Stuff[item.Name] = item
	}
	return
}

func atoi(s string) (int, error) {
	s = strings.Replace(s, ",", "", 1)
	if s == "" {
		return 0, nil
	}
	return strconv.Atoi(s)
}
