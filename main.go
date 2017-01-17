package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func main() {
	count := flag.Int("c", 100, "Count")
	interval := flag.Int("i", 1, "Interval in seconds")
	flag.Parse()
	target := flag.Arg(0)

	// parse url
	var host string
	var port string

	if !strings.HasPrefix(target, "http") {
		target = "http://" + target
	}

	u, err := url.Parse(target)
	if err != nil {
		log.Fatal(err)
	}

	if strings.Contains(u.Host, ":") {
		host, port, _ = net.SplitHostPort(u.Host)
	} else {
		host = u.Host
		if strings.Contains(u.Scheme, "https") {
			port = "443"
		} else {
			port = "80"
		}
	}

	// ip lookup
	ip := getIP(host)

	// http get
	for i := 0; ; i++ {
		response, elapsed := httpGet(u.Scheme + "://" + u.Host + u.Path)
		fmt.Printf("connected to %s:%s, seq=%d, time=%d ms, response=%s\n", ip, port, i, elapsed, response)
		if i >= *count {
			break
		} else {
			time.Sleep(time.Duration(*interval) * time.Second)
		}
	}
}

func httpGet(s string) (string, int64) {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	start := time.Now()
	resp, err := client.Get(s)
	if err != nil {
		return "\033[31mNo response\033[0m", 0
	}
	elapsed := int64(time.Since(start) / time.Millisecond)

	statusCode := strconv.Itoa(resp.StatusCode)
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		return "\033[32m" + statusCode + "\033[0m", elapsed
	} else if resp.StatusCode >= 300 && resp.StatusCode <= 399 {
		return "\033[33m" + statusCode + "\033[0m", elapsed
	} else {
		return "\033[31m" + statusCode + "\033[0m", elapsed
	}
}

func getIP(s string) string {
	addr, err := net.LookupHost(s)
	if err != nil {
		log.Fatal(err)
	}
	return addr[0]
}
