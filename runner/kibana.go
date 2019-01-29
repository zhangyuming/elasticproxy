package runner

import (
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	"elasticproxy/proxy"
	"elasticproxy/vlog"
	"strings"
)

type kibana struct{}

func (k *kibana)RequestModify(req *proxy.LocalRequest){

	vlog.Debug("kibana request modify invoke begin")
	if strings.Contains(req.Path,"_msearch"){
		vlog.Debug("request body : ", req.Body)
		bs :=  strings.Split(req.Body,"\n")
		for index, doc := range bs{
			r := gjson.Get(doc,"sort.#.@timestamp.order")
			if(r.Exists()){
				direction := extraStrFromArr(r)
				sort := gjson.Get(doc,"sort");
				var err error
				switch direction {
				case "desc":{
					bs[index],err = sjson.Set(doc,"sort", updateSortStrategy(sort,map[string]interface{}{"offset": map[string]string{"order": "desc", "unmapped_type": "boolean"}}))
				}
				case "asc":{
					bs[index],err = sjson.Set(doc,"sort", updateSortStrategy(sort,map[string]interface{}{"offset": map[string]string{"order": "asc", "unmapped_type": "boolean"}}))
				}
				default:
					continue

				}
				if err != nil{
					vlog.Error("update kibana query str failed")
					return
				}
			}
		}
		req.Body = strings.Join(bs,"\n")
		vlog.Debug("convert requestbody : ",req.Body )

	}
	vlog.Debug("kibana request modify invoke end")

}

// 从gjson result中提取字符串， 如果result是array 则获取第一个的值
func extraStrFromArr(result gjson.Result) string  {
	if(result.IsArray()){
		return extraStrFromArr(result.Array()[0])
	}else {
		return result.Str
	}
}

// 把offset 排序规则插入到第二条排序规则位置
func updateSortStrategy(result gjson.Result, v interface{})(interface{})  {
	var arr = "[]"
	if(result.IsArray()){
		for i,a := range result.Array(){

			var err error
			arr,err = sjson.Set(arr,"-1",a.Value());
			if err != nil{
				vlog.Error("set sort failed", err)
			}
			if(i == 0){
				var err error
				arr,err = sjson.Set(arr,"-1",v)
				if err != nil{
					vlog.Error("set sort failed", err)
				}
			}
		}
	}else{
		vlog.Debug("sort field is not array skip")
		return result.Value()
	}
	return  gjson.Parse(arr).Value();
}
