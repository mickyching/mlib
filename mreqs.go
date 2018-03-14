package mlib

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

// HttpServe run a http server
// addr format 0.0.0.0:8080
// performance test tool and setting
// **socket: too many open files**
//   ulimit -a | grep open      // shows 1024
//   ulimit -n 65535            // set a large number
//   lsof -n|awk '{print $2}'|sort|uniq -c|sort -nr|head // list num pid
// **connect: cannot asign requested address
//   cat /proc/net/sockstat     // shows socket usage status
// **apr_socket_recv: Connection reset by peer**
//   server has too many connections possible syn flooding
//   set server /etc/sysctl.conf: net.ipv4.tcp_syncookies = 0
// **profile using pprof and graphviz**
//   go tool pprof ./binary URL/debug/pprof/profile // CPU-profile, MEM-heap
//   usful cmd: top10/web [func]/list [func]
// **useful mux/handler**
//   NewServeMux() // create mux replace DefaultServeMux
//   FileServer, NotFoundHandler, RedirectHandler
func HttpServe(addr string, route func(w http.ResponseWriter, r *http.Request)) error {
	http.HandleFunc("/", route)
	return http.ListenAndServe(addr, nil)
}

// HttpPost post request to url
func HttpPost(url string, reqs, resp interface{}, timeout int) error {
	body, err := json.Marshal(reqs)
	if err != nil {
		return err
	}
	client := http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}
	raw, err := client.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	defer raw.Body.Close()

	body, err = ioutil.ReadAll(raw.Body)
	if err != nil {
		return err
	}

	if len(body) != 0 && resp != nil {
		err = json.Unmarshal(body, resp)
		if err != nil {
			return err
		}
	}

	return nil
}

// HttpDelete run delete method to url
func HttpDelete(url string, resp interface{}, timeout int) error {
	client := &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}
	request, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	raw, err := client.Do(request)
	if err != nil {
		return err
	}
	defer raw.Body.Close()

	body, err := ioutil.ReadAll(raw.Body)
	if err != nil {
		return err
	}

	if len(body) != 0 && resp != nil {
		err = json.Unmarshal(body, resp)
		if err != nil {
			return err
		}
	}

	return nil
}
