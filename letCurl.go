/**
 * go版本的curl请求库
 * 目前仅支持简单的请求 GET POST 满足一般的要求
 * @author xiaowen <952065841@qq.com>
 * @blog http://www.az1314.cn
 */
package letCurl

import (
	"net/http"
	"errors"
	"strings"
	"io/ioutil"
)

type letCurl struct {
	url      string
	method   string
	request  *request
}
type request struct {
	Header   map[string]string
	Form     map[string]string
	Raw      *raw
}
type raw struct {
	RawType  string
	RawValue string
}
type response struct {
	Headers map[string]string
	Body    string
}
var RawType = [...]string{"text", "json", "xml", "html"}
var err = make([]error, 0, 10)
func (this *letCurl)SetUrl(url string)*letCurl  {
	this.url = url
	return this
}
func (this *letCurl)SetHeader(key string, value string)*letCurl  {
	this.request.Header[key] = value
	return this
}
func (this *letCurl)SetForm(key string, value string)*letCurl  {
	this.request.Form[key] = value
	return this
}
func (this *letCurl)Get()(string, error) {
	this.method = "GET"
	return this.Start()
}
func (this *letCurl)Post()(string, error) {
	this.method = "POST"
	//默认
	if this.request.Raw == nil {
		this.request.Header["Content-Type"] = "application/x-www-form-urlencoded"
	}
	return this.Start()
}
func (this *letCurl)Start()(string, error) {
	if len(err) != 0 {
		return "", err[0]
	}
	if this.url == "" {
		return "", errors.New("curl param url is null")
	}
	if this.method == "" {
		return "", errors.New("curl param method is null")
	}
	var form string
	if len(this.request.Form) != 0 && this.request.Raw == nil{
		for k, v := range this.request.Form {
			form += k+"="+v+"&"
		}
		form = strings.TrimRight(form, "&")
	}
	req, err := http.NewRequest(this.method, this.url, strings.NewReader(form))
	if err != nil {
		return "", err
	}
	for k, v := range this.request.Header {
		req.Header.Set(k, v)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body), err
}
func (this *letCurl)SetRaw(t string, value string)*letCurl  {
	for _, v := range RawType{
		if v == t {
			raw := &raw{t,value}
			this.request.Raw = raw
			this.request.Header["Content-Type"] = "application/"+t
			return this
		}
	}
	err = append(err, errors.New("RowType is not supported"))
	return this
}
func NewCurl() *letCurl {
	obj := &letCurl{}
	obj.request = &request{}
	obj.request.Header = make(map[string]string)
	obj.request.Form = make(map[string]string)
	return obj
}