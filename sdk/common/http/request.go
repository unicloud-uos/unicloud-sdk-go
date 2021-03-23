package common

import (
	"bytes"
	"encoding/json"
	"io"
	//"log"
	"math/rand"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const (
	POST = "POST"
	GET  = "GET"
	PUT  = "PUT"

	HTTP  = "http"
	HTTPS = "https"

	RootDomain = "api.unicloud.com"
	Path       = ""
)

type Request interface {
	GetAction() string
	GetBodyReader() io.Reader
	GetBodyData() io.Reader
	GetScheme() string
	GetRootDomain() string
	GetServiceDomain(string) string
	GetDomain() string
	GetHttpMethod() string
	GetParams() map[string]string
	GetService() string
	GetVersion() string
	GetRequestUrl() string

	SetScheme(string)
	SetRootDomain(string)
	SetDomain(string)
	SetHttpMethod(string)
	SetBodyData(v interface{})
}

type BaseRequest struct {
	httpMethod string
	scheme     string
	rootDomain string
	domain     string
	path       string
	params     map[string]string
	formParams map[string]string

	service string
	version string
	action  string
	data io.Reader // post body
	accessKeyId string
}

// 获取请求url,未编码的
func (r *BaseRequest) GetRequestUrl() string {
	if r.httpMethod == GET {
		return r.GetScheme() + "://" + r.domain + r.path + "?" + GetUrlQueries(r.params)
	} else if r.httpMethod == POST {
		return r.GetScheme() + "://" + r.domain + r.path + "?" + GetUrlQueries(r.params)
	} else if r.httpMethod == PUT{
		return r.GetScheme() + "://" + r.domain + r.path + "?" + GetUrlQueries(r.params)
	} else {
		return ""
	}
}

func (r *BaseRequest) GetAk() string  {
	return r.accessKeyId
}

func (r *BaseRequest) GetBodyData()  io.Reader{
	return r.data
}

func (r *BaseRequest) SetBodyData(v interface{}) {
	bytesData, _ := json.Marshal(v)
	r.data = bytes.NewReader(bytesData)
}

func (r *BaseRequest) GetAction() string {
	return r.action
}

func (r *BaseRequest) GetHttpMethod() string {
	return r.httpMethod
}

func (r *BaseRequest) GetParams() map[string]string {
	return r.params
}

func (r *BaseRequest) GetPath() string {
	return r.path
}

func (r *BaseRequest) GetDomain() string {
	return r.domain
}

func (r *BaseRequest) GetScheme() string {
	return r.scheme
}

func (r *BaseRequest) GetRootDomain() string {
	return r.rootDomain
}

func (r *BaseRequest) GetServiceDomain(service string) (domain string) {
	// request 中更改domain
	rootDomain := r.rootDomain
	if rootDomain == "" {
		rootDomain = RootDomain
	}
	domain = RootDomain + "/" + service
	return
}

func (r *BaseRequest) SetDomain(domain string) {
	r.domain = domain
}

func (r *BaseRequest) SetScheme(scheme string) {
	scheme = strings.ToLower(scheme)
	switch scheme {
	case HTTP:
		r.scheme = HTTP
	default:
		r.scheme = HTTPS
	}
}

func (r *BaseRequest) SetRootDomain(rootDomain string) {
	r.rootDomain = rootDomain
}

func (r *BaseRequest) SetHttpMethod(method string) {
	switch strings.ToUpper(method) {
	case POST:
		{
			r.httpMethod = POST
		}
	case GET:
		{
			r.httpMethod = GET
		}
	case PUT:
		{
			r.httpMethod = PUT
		}
	default:
		{
			r.httpMethod = GET
		}
	}
}

func (r *BaseRequest) GetService() string {
	return r.service
}

// url encode 之后的 ，暂时用不上
func (r *BaseRequest) GetUrl() string {
	if r.httpMethod == GET {
		return r.GetScheme() + "://" + r.domain + r.path + "?" + GetUrlQueriesEncoded(r.params)
	} else if r.httpMethod == POST {
		return r.GetScheme() + "://" + r.domain + r.path
	} else {
		return ""
	}
}

func (r *BaseRequest) GetVersion() string {
	return r.version
}

func GetUrlQueries(params map[string]string) string {
	values := ""
	for key, value := range params {
		if value != "" {
			values += "&" + key + "=" +value
		}
	}
	return values[1:]
}

func GetUrlQueriesEncoded(params map[string]string) string {
	values := url.Values{}
	for key, value := range params {
		if value != "" {
			values.Add(key, value)
		}
	}
	return values.Encode()
}

func (r *BaseRequest) GetBodyReader() io.Reader {
	if r.httpMethod == POST {
		s := GetUrlQueriesEncoded(r.params)
		return strings.NewReader(s)
	} else {
		return strings.NewReader("")
	}
}

func (r *BaseRequest) Init() *BaseRequest {
	r.domain = ""
	r.path = Path
	r.params = make(map[string]string)
	r.formParams = make(map[string]string)
	return r
}

func (r *BaseRequest) WithApiInfo(service, version, action string) *BaseRequest {
	r.service = service
	r.version = version
	r.action = action
	return r
}

/**
公共参数
 */
func CompleteCommonParams(request Request, region string) {
	params := request.GetParams()
	params["RegionId"] = region
	if request.GetVersion() != "" {
		params["Version"] = request.GetVersion()
	}
	params["Action"] = request.GetAction()
	//params["Timestamp"] = strconv.FormatInt(time.Now().Unix(), 10)
	params["Timestamp"] = time.Now().Format("2006-01-02T15:04:05Z")
	params["SignatureNonce"] = strconv.Itoa(rand.Int())
	params["SignatureMethod"] = "HMAC-SHA1"
	params["Format"] = "json"
	params["SignatureVersion"] = "2.0"
}

// 把对象的值反射到param
func ConstructParams(req Request) (err error) {
	value := reflect.ValueOf(req).Elem()
	err = flatStructure(value, req, "")
	//log.Printf("[DEBUG] params=%s", req.GetParams())
	return
}

func flatStructure(value reflect.Value, request Request, prefix string) (err error) {
	//log.Printf("[DEBUG] reflect value: %v", value.Type())
	valueType := value.Type()
	for i := 0; i < valueType.NumField(); i++ {
		tag := valueType.Field(i).Tag
		nameTag, hasNameTag := tag.Lookup("name")
		// 属性无tag
		if !hasNameTag {
			continue
		}
		field := value.Field(i)
		kind := field.Kind()
		// 类型是指针
		if kind == reflect.Ptr && field.IsNil() {
			continue
		}
		if kind == reflect.Ptr {
			field = field.Elem()
			kind = field.Kind()
		}
		key := prefix + nameTag
		if kind == reflect.String {
			s := field.String()
			if s != "" {
				request.GetParams()[key] = s
			}
		} else if kind == reflect.Bool {
			request.GetParams()[key] = strconv.FormatBool(field.Bool())
		} else if kind == reflect.Int || kind == reflect.Int64 {
			request.GetParams()[key] = strconv.FormatInt(field.Int(), 10)
		} else if kind == reflect.Uint || kind == reflect.Uint64 {
			request.GetParams()[key] = strconv.FormatUint(field.Uint(), 10)
		} else if kind == reflect.Float64 {
			request.GetParams()[key] = strconv.FormatFloat(field.Float(), 'f', -1, 64)
		} else if kind == reflect.Slice {
			list := value.Field(i)
			for j := 0; j < list.Len(); j++ {
				vj := list.Index(j)
				key := prefix + nameTag + "." + strconv.Itoa(j)
				kind = vj.Kind()
				if kind == reflect.Ptr && vj.IsNil() {
					continue
				}
				if kind == reflect.Ptr {
					vj = vj.Elem()
					kind = vj.Kind()
				}
				if kind == reflect.String {
					request.GetParams()[key] = vj.String()
				} else if kind == reflect.Bool {
					request.GetParams()[key] = strconv.FormatBool(vj.Bool())
				} else if kind == reflect.Int || kind == reflect.Int64 {
					request.GetParams()[key] = strconv.FormatInt(vj.Int(), 10)
				} else if kind == reflect.Uint || kind == reflect.Uint64 {
					request.GetParams()[key] = strconv.FormatUint(vj.Uint(), 10)
				} else if kind == reflect.Float64 {
					request.GetParams()[key] = strconv.FormatFloat(vj.Float(), 'f', -1, 64)
				} else {
					if err = flatStructure(vj, request, key+"."); err != nil {
						return
					}
				}
			}
		} else {
			if err = flatStructure(reflect.ValueOf(field.Interface()), request, prefix+nameTag+"."); err != nil {
				return
			}
		}
	}
	return
}
