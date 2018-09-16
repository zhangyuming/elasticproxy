package runner

import (
	"elasticproxy/proxy"
	"strings"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	"elasticproxy/vlog"
)



type kibana struct{}

func init() {
	proxy.RegistryRequestModifyer(&kibana{})
}


func (k *kibana)RequestModify(req *proxy.LocalRequest){

	vlog.Debug("kibana request modify invoke begin")
	if strings.Contains(req.Path,"_msearch"){
		vlog.Debug("request body : ", req.Body)
		bs :=  strings.Split(req.Body,"\n")
		r := gjson.Get(bs[1],"sort")
		var err error
		if r.Exists() {
			bs[1],err = sjson.Set(bs[1], "sort.-1", map[string]interface{}{"offset": map[string]string{"order": "desc", "unmapped_type": "boolean"}})
			if err != nil {
				vlog.Error("convert request body failed ", err)
			}
			req.Body = strings.Join(bs,"\n")
			vlog.Debug("convert requestbody : ",req.Body )
		}
	}
	vlog.Debug("kibana request modify invoke end")
}
