package main

import (
	"flag"
	"net/http"
	"net/http/httputil"
	"strings"
	"io/ioutil"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

var loglevel = "info"
var elastic_host = "elasticsearch-default-elasticsearch-client:9200"
var localport = "8899"

type offset struct {
	order         string
	unmapped_type string
}

func main() {

	flag.StringVar(&loglevel, "d", "info", "log level [info|debug|error]")
	flag.StringVar(&elastic_host, "elastic_host", "elasticsearch-default-elasticsearch-client:9200", " elastic address")
	flag.StringVar(&localport, "p", "8899", "service port")
	flag.Parse()
	vlog := GetLog(loglevel)

	vlog.Info("proxy port : ", localport, " elasticseach host : ",elastic_host)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		vlog.Info(r.Method, r.URL.Path)
		for _, h := range r.Header {
			vlog.Debug("header =============")
			vlog.Debug(h)
			vlog.Debug("header end =============")
		}
		director := func(req *http.Request) {
			req.URL.Scheme = "http"
			req.URL.Host = elastic_host
			req.Host = elastic_host
			req.Body = r.Body
			if strings.Contains(r.URL.Path,"_msearch"){
				bt, err := ioutil.ReadAll(req.Body)
				if err != nil {
					vlog.Error("read request body fail", err)
				}

				body := string(bt)
				vlog.Debug("request body : ", body)
				bs :=  strings.Split(body,"\n")
				r := gjson.Get(bs[1],"sort")
				if r.Exists() {
					bs[1], err = sjson.Set(bs[1], "sort.-1", map[string]interface{}{"offset": map[string]string{"order": "desc", "unmapped_type": "boolean"}})
					if err != nil {
						vlog.Error("convert request body failed ", err)
					}
					body = strings.Join(bs,"\n")
					vlog.Debug("convert requestbody : ",body )
				}

				req.Body = ioutil.NopCloser(strings.NewReader(body))
				req.ContentLength = int64(len(body))

			}
		}
		proxy := &httputil.ReverseProxy{Director: director}
		proxy.ServeHTTP(w, r)
	})

	vlog.Error(http.ListenAndServe(":"+localport, nil))

}
