package docs

import (
	"encoding/json"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/swagger"
)

var rootinfo string = `{"apiVersion":"1.0.0","swaggerVersion":"1.2","apis":[{"path":"/admin","description":""},{"path":"/common","description":""}],"info":{"title":"羽球之家网站API","description":"羽球之家网站API。","contact":"doomsplayer@gmail.com"}}`
var subapi string = `{"/admin":{"apiVersion":"1.0.0","swaggerVersion":"1.2","basePath":"","resourcePath":"/admin","produces":["application/json","application/xml","text/plain","text/html"],"apis":[{"path":"/carousel","description":"","operations":[{"httpMethod":"PUT","nickname":"add carousel","type":"","summary":"添加首页轮播图","parameters":[{"paramType":"form","name":"title","description":"标题","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"form","name":"subtitle","description":"小标题","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"form","name":"url","description":"Url","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"form","name":"picid","description":"图片的Id","dataType":"int","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"结果json","responseModel":""},{"code":404,"message":"Not found errmsg","responseModel":""},{"code":401,"message":"Need Permmision errmsg","responseModel":""},{"code":500,"message":"Server Error errmsg","responseModel":""}]}]}]},"/common":{"apiVersion":"1.0.0","swaggerVersion":"1.2","basePath":"","resourcePath":"/common","produces":["application/json","application/xml","text/plain","text/html"],"apis":[{"path":"/mainBar","description":"","operations":[{"httpMethod":"GET","nickname":"mainBar","type":"","summary":"获得首页目录","responseMessages":[{"code":200,"message":"目录的json","responseModel":""},{"code":404,"message":"Not found","responseModel":""}]}]},{"path":"/upload","description":"","operations":[{"httpMethod":"PUT","nickname":"Upload File","type":"","summary":"上传文件","parameters":[{"paramType":"form","name":"file","description":"上传的文件","dataType":"file","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"文件Id的json","responseModel":""},{"code":404,"message":"Not found errmsg","responseModel":""},{"code":401,"message":"Need Permmision errmsg","responseModel":""},{"code":500,"message":"Server Error errmsg","responseModel":""}]}]},{"path":"/upload","description":"","operations":[{"httpMethod":"DELETE","nickname":"Delete File","type":"","summary":"删除文件","parameters":[{"paramType":"query","name":"id","description":"文件id","dataType":"int","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"返回json","responseModel":""},{"code":404,"message":"Not found errmsg","responseModel":""},{"code":401,"message":"Need Permmision errmsg","responseModel":""},{"code":500,"message":"Server Error errmsg","responseModel":""}]}]},{"path":"/upload","description":"","operations":[{"httpMethod":"GET","nickname":"Get File","type":"","summary":"获得文件","parameters":[{"paramType":"query","name":"id","description":"文件id","dataType":"int","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"返回json","responseModel":""},{"code":404,"message":"Not found errmsg","responseModel":""},{"code":401,"message":"Need Permmision errmsg","responseModel":""},{"code":500,"message":"Server Error errmsg","responseModel":""}]}]},{"path":"/carousel","description":"","operations":[{"httpMethod":"GET","nickname":"list carousel","type":"","summary":"拉取首页轮播图","parameters":[{"paramType":"query","name":"num","description":"拉取几张","dataType":"int","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"轮播的json","responseModel":""},{"code":404,"message":"Not found errmsg","responseModel":""},{"code":401,"message":"Need Permmision errmsg","responseModel":""},{"code":500,"message":"Server Error errmsg","responseModel":""}]}]}]}}`
var rootapi swagger.ResourceListing

var apilist map[string]*swagger.ApiDeclaration

func init() {
	basepath := "/api"
	err := json.Unmarshal([]byte(rootinfo), &rootapi)
	if err != nil {
		beego.Error(err)
	}
	err = json.Unmarshal([]byte(subapi), &apilist)
	if err != nil {
		beego.Error(err)
	}
	beego.GlobalDocApi["Root"] = rootapi
	for k, v := range apilist {
		for i, a := range v.Apis {
			a.Path = urlReplace(k + a.Path)
			v.Apis[i] = a
		}
		v.BasePath = basepath
		beego.GlobalDocApi[strings.Trim(k, "/")] = v
	}
}


func urlReplace(src string) string {
	pt := strings.Split(src, "/")
	for i, p := range pt {
		if len(p) > 0 {
			if p[0] == ':' {
				pt[i] = "{" + p[1:] + "}"
			} else if p[0] == '?' && p[1] == ':' {
				pt[i] = "{" + p[2:] + "}"
			}
		}
	}
	return strings.Join(pt, "/")
}
