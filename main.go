package main

import (
	"flag"
	"net/http"
	"net/http/httputil"
	"strings"
	"io/ioutil"
	"elasticproxy/proxy"
	"strconv"
	"fmt"
	_ "elasticproxy/runner"
	"elasticproxy/vlog"
)

var loglevel = "info"
var elastic_host = "elasticsearch:9200"
var localport = "8899"

func main() {

	flag.StringVar(&loglevel, "d", "info", "log level [info|debug|error]")
	flag.StringVar(&elastic_host, "elastic_host", "elasticsearch:9200", " elastic address")
	flag.StringVar(&localport, "p", "8899", "service port")
	flag.Parse()

	vlog.SetLogLevel(loglevel)

	vlog.Info("proxy port : ", localport, " elasticseach host : ",elastic_host)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		vlog.Info(r.Method, r.URL.Path)
		/*for _, h := range r.Header {
			vlog.Debug("header =============")
			vlog.Debug(h)
			vlog.Debug("header end =============")
		}*/
		director := func(req *http.Request) {

			// wrapper localrequest
			localRequest := proxy.LocalRequest{}
			localRequest.Scheme = r.URL.Scheme
			localRequest.Host = strings.Split(elastic_host,":")[0]
			localRequest.Path = r.URL.Path
			p,err := strconv.Atoi(strings.Split(elastic_host,":")[1])
			if err != nil{
				vlog.Error("convert elastic port failt ,", elastic_host, err)
			}
			localRequest.Port = p
			bt,err := ioutil.ReadAll(r.Body)
			if err != nil{
				vlog.Error("get request body faild ",err)
			}
			localRequest.Body = string(bt)

			headers := http.Header{}
			for k,v := range r.Header{
				for _,v2 := range v{
					headers.Add(k,v2)
				}
			}
			localRequest.Header = headers

			reqmds := proxy.GetRquestModifyers()
			for _,md := range  reqmds{
				md.RequestModify(&localRequest)
			}

			if localRequest.Scheme == ""{
				req.URL.Scheme = "http"
			}else{
				req.URL.Scheme = localRequest.Scheme
			}

			req.URL.Host = fmt.Sprint(localRequest.Host,":", localRequest.Port)
			req.Header = localRequest.Header
			req.Body = ioutil.NopCloser(strings.NewReader(localRequest.Body))
			req.ContentLength = int64(len(localRequest.Body))
			req.Host = r.Host


		}
		proxy := &httputil.ReverseProxy{Director: director}
		proxy.ServeHTTP(w, r)
	})

	vlog.Error(http.ListenAndServe(":"+localport, nil))

}
