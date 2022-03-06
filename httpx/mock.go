package httpx

import (
	"bytes"
	"compress/gzip"
	"errors"
	"github.com/hetianyi/easygo/base"
	"github.com/hetianyi/easygo/convert"
	json "github.com/json-iterator/go"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"reflect"
	"strings"
	"time"
)

const (
	MethodGet                 = "GET"
	MethodPost                = "POST"
	MethodTrace               = "TRACE"
	MethodDelete              = "DELETE"
	MethodPut                 = "PUT"
	MethodOptions             = "OPTIONS"
	MethodHead                = "HEAD"
	MethodConnect             = "CONNECT"
	ContentTypeFormUrlEncoded = "application/x-www-form-urlencoded"
	ContentTypeJson           = "application/json"
	ContentTypeFormMultipart  = "multipart/form-data"
)

var (
	defaultHeaders       = make(map[string]string, 10)
	allowedResponseTypes = make(map[string]bool)
	defaultTimeout       = time.Second * 30
)

type FormItemType int

const (
	FormItemTypeField FormItemType = iota
	FormItemTypeFile
)

const (
	multipartBodyParamType         = "[]httpx.FormItem"
	urlencodedBodyOrQueryParamType = "map[string][]string"
)

type FormItem struct {
	Type       FormItemType
	Name       string
	Value      string
	FileName   string
	FileReader io.Reader
}

func init() {
	allowedResponseTypes["int"] = true
	allowedResponseTypes["int64"] = true
	allowedResponseTypes["float32"] = true
	allowedResponseTypes["float64"] = true
	allowedResponseTypes["bool"] = true
	allowedResponseTypes["string"] = true
	allowedResponseTypes["struct"] = true
	allowedResponseTypes["map"] = true
	allowedResponseTypes["nil"] = true
	allowedResponseTypes["io.Writer"] = true
}

// mock is a fake http request instance.
type mock struct {
	httpClient        *http.Client
	url               string
	method            string
	headers           map[string][]string
	contentType       string
	parameterMap      map[string][]string
	body              []byte
	multipartFormBody []FormItem
	request           http.Request
	response          http.Response
	responseContainer interface{}
	successCodes      []int
}

// Mock returns an initialized mock.
func Mock() *mock {
	return &mock{
		httpClient: &http.Client{
			Timeout: defaultTimeout,
		},
		headers: map[string][]string{
			"User-Agent": {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36"},
		},
		contentType:       ContentTypeJson,
		successCodes:      []int{http.StatusOK},
		responseContainer: "",
		parameterMap:      make(map[string][]string),
	}
}

// Header adds a http header to the mock.
func (m *mock) Header(name, value string) *mock {
	if !checkHeaderExist(m.headers[name], value) {
		m.headers[name] = append(m.headers[name], value)
	}
	return m
}

// Headers adds many http headers to the mock.
func (m *mock) Headers(headers map[string]string) *mock {
	for k, v := range headers {
		if !checkHeaderExist(m.headers[k], v) {
			m.headers[k] = append(m.headers[k], v)
		}
	}
	return m
}

func checkHeaderExist(vals []string, value string) bool {
	for _, v := range vals {
		if v == value {
			return true
		}
	}
	return false
}

// ContentType sets ContentType of the mock.
func (m *mock) ContentType(contentType string) *mock {
	if strings.HasPrefix(contentType, ContentTypeFormUrlEncoded) ||
		strings.HasPrefix(contentType, ContentTypeJson) ||
		strings.HasPrefix(contentType, ContentTypeFormMultipart) {
		m.contentType = contentType
	} else {
		panic(errors.New("not supported ContentType: '" + contentType +
			"', ContentType is currently only support " + "'" + ContentTypeFormUrlEncoded +
			"', '" + ContentTypeFormMultipart + "' and '" + ContentTypeJson + "'"))
	}
	return m
}

// Parameters add parameters on the request url.
func (m *mock) Parameters(params map[string]string) *mock {
	for k, v := range params {
		m.parameterMap[k] = append(m.parameterMap[k], v)
	}
	return m
}

// Parameter add a parameter on the request url.
func (m *mock) Parameter(name, value string) *mock {
	m.parameterMap[name] = append(m.parameterMap[name], value)
	return m
}

// RequestBody 设置请求Body.
//
//  如果 ContentType 是 'application/x-www-form-urlencoded'，那么请求Body的类型必须是 map[string][]string,
//
//  如果 ContentType 是 'multipart/form-data'，那么请求Body的类型必须是 TODO,
//
//  如果 ContentType 是 'application/json'，那么请求Body的类型可以是任意自定义struct.
func (m *mock) RequestBody(body interface{}) *mock {
	if body == nil {
		m.body = nil
		return m
	}
	bodyType := reflect.TypeOf(body).String()
	if m.contentType == ContentTypeJson {
		jv, err := json.Marshal(body)
		if err != nil {
			panic(err)
		}
		m.body = jv
	} else if m.contentType == ContentTypeFormMultipart {
		// processes it in func Do()
		if bodyType != multipartBodyParamType {
			panic(errors.New("request body type must be '" + multipartBodyParamType + "' if ContentType is '" + ContentTypeFormMultipart + "'"))
		}
		m.multipartFormBody = body.([]FormItem)
	} else {
		if bodyType != urlencodedBodyOrQueryParamType {
			panic(errors.New("body type must be '" + urlencodedBodyOrQueryParamType + "' if ContentType is '" + ContentTypeFormUrlEncoded + "'"))
		}
		m.body = encodeParameters(body.(map[string][]string))
	}
	return m
}

// Success defines the response type and tells what status codes should be recognized as success request,
//
// response type must be one of:
//
// int int64 float32 float64 bool string map nil or pointer of a struct.
func (m *mock) Success(response interface{}, successCodes ...int) *mock {
	m.responseContainer = response
	m.successCodes = successCodes
	if !allowedResponseTypes[checkResponseType(response)] {
		panic("response type not allowed")
	}
	return m
}

// Do is the end of the mock chain,
// which will send the request and return the result.
func (m *mock) Do() (interface{}, int, error) {
	paramsStr := string(encodeParameters(m.parameterMap))

	isMultipart := false
	var mw *multipart.Writer
	var pipeReader *io.PipeReader
	var pipeWriter *io.PipeWriter

	// 如果是FormMultipart，则将提供请求Body
	if m.contentType == ContentTypeFormMultipart && len(m.multipartFormBody) > 0 {
		isMultipart = true
		pipeReader, pipeWriter = io.Pipe()
		mw = multipart.NewWriter(pipeWriter)

		go func() {
			defer pipeWriter.Close()
			defer mw.Close()
			for _, item := range m.multipartFormBody {
				if item.Type == FormItemTypeFile {
					o, err := mw.CreateFormFile(item.Name, item.FileName)
					if err != nil {
						panic(err)
					}
					_, err = io.Copy(o, item.FileReader)
					if err != nil {
						panic(err)
					}
				} else {
					o, err := mw.CreateFormField(item.Name)
					if err != nil {
						panic(err)
					}
					_, err = o.Write([]byte(item.Value))
					if err != nil {
						panic(err)
					}
				}
			}
		}()
	}

	// 确定请求的BodyReader
	var bodyReader io.Reader
	if isMultipart {
		bodyReader = pipeReader
	} else {
		bodyReader = bytes.NewReader(m.body)
	}

	// 构造请求
	req, err := http.NewRequest(m.method, base.TValue(paramsStr == "", m.url, m.url+"?"+paramsStr).(string), bodyReader)
	if err != nil {
		return m.responseContainer, 0, err
	}

	// 设置请求headers
	req.Header = m.headers
	if isMultipart {
		req.Header.Set("Content-Type", mw.FormDataContentType())
	}

	resp, err := m.httpClient.Do(req)
	if err != nil {
		return m.responseContainer, 0, err
	}

	// 判断请求是否成功
	if !m.isSuccess(resp.StatusCode) {
		_, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return m.responseContainer, resp.StatusCode, err
		}
		return m.responseContainer, resp.StatusCode, nil
	}

	// 处理gzp结果
	decodeUseGzip := isGzipContentType(resp.Header)

	var r io.Reader
	if decodeUseGzip {
		reader, err := gzip.NewReader(resp.Body)
		if err != nil {
			return m.responseContainer, resp.StatusCode, err
		}
		r = reader
	} else {
		r = resp.Body
	}
	ret, err := convertResponse(checkResponseType(m.responseContainer), r, m.responseContainer)
	return ret, resp.StatusCode, err
}

func isGzipContentType(header http.Header) bool {
	for k, v := range header {
		if strings.ToLower(k) == "content-encoding" {
			if len(v) > 0 {
				if strings.ToLower(v[0]) == "gzip" {
					return true
				}
			}
		}
	}
	return false
}

// Get sets http method to GET.
func (m *mock) Get(url string) *mock {
	m.method = MethodGet
	m.url = url
	return m
}

// Post sets http method to Post.
func (m *mock) Post(url string) *mock {
	m.method = MethodPost
	m.url = url
	return m
}

// Options sets http method to Options.
func (m *mock) Options(url string) *mock {
	m.method = MethodOptions
	m.url = url
	return m
}

// Head sets http method to Head.
func (m *mock) Head(url string) *mock {
	m.method = MethodHead
	m.url = url
	return m
}

// Put sets http method to Put.
func (m *mock) Put(url string) *mock {
	m.method = MethodPut
	m.url = url
	return m
}

// Delete sets http method to Delete.
func (m *mock) Delete(url string) *mock {
	m.method = MethodDelete
	m.url = url
	return m
}

// Connect sets http method to Connect.
func (m *mock) Connect(url string) *mock {
	m.method = MethodConnect
	m.url = url
	return m
}

// Trace sets http method to Trace.
func (m *mock) Trace(url string) *mock {
	m.method = MethodTrace
	m.url = url
	return m
}

// SetTTL sets http client request timeout value.
// Default timeout value is 20s.
// A Timeout of zero means no timeout.
func (m *mock) SetTTL(timeout time.Duration) {
	m.httpClient.Timeout = timeout
}

// encodeParameters encodes parameters to the pattern of 'a=xx&b=xx'.
func encodeParameters(params map[string][]string) []byte {
	if len(params) == 0 {
		return []byte{}
	}
	var buffer bytes.Buffer
	for k, vl := range params {
		if len(vl) == 0 {
			buffer.WriteString(k)
			buffer.WriteString("=")
			buffer.WriteString("&")
			continue
		}
		for i, v := range vl {
			buffer.WriteString(k)
			buffer.WriteString("=")
			buffer.WriteString(v)
			if i != len(vl)-1 {
				buffer.WriteString("&")
			}
		}
	}
	return buffer.Bytes()
}

// isSuccess determines whether the request is success.
func (m *mock) isSuccess(code int) bool {
	for _, v := range m.successCodes {
		if v == code {
			return true
		}
	}
	return code == http.StatusOK
}

// checkResponseType returns the type of response data container.
func checkResponseType(resp interface{}) string {
	if resp == nil {
		return "nil"
	}
	if _, c := resp.(io.Writer); c {
		return "io.Writer"
	}
	typ := reflect.TypeOf(resp)
	for {
		if typ.Kind() == reflect.Ptr {
			typ = typ.Elem()
			continue
		}
		break
	}
	return typ.Kind().String()
}

// convertResponse converts response to the type of response.
func convertResponse(typeName string, response io.Reader, responseContainer interface{}) (interface{}, error) {
	switch typeName {
	case "nil":
		_, err := ioutil.ReadAll(response)
		return nil, err
	case "io.Writer":
		io.Copy(responseContainer.(io.Writer), response)
		return response, nil
	case "int":
		bs, err := ioutil.ReadAll(response)
		if err != nil {
			return 0, err
		}
		return convert.StrToInt(string(bs))
	case "int64":
		bs, err := ioutil.ReadAll(response)
		if err != nil {
			return 0, err
		}
		return convert.StrToInt64(string(bs))
	case "float32":
		bs, err := ioutil.ReadAll(response)
		if err != nil {
			return 0, err
		}
		return convert.StrToFloat32(string(bs))
	case "float64":
		bs, err := ioutil.ReadAll(response)
		if err != nil {
			return 0, err
		}
		return convert.StrToFloat64(string(bs))
	case "bool":
		bs, err := ioutil.ReadAll(response)
		if err != nil {
			return 0, err
		}
		return convert.StrToBool(string(bs))
	case "string":
		bs, err := ioutil.ReadAll(response)
		if err != nil {
			return 0, err
		}
		return string(bs), nil
	case "map", "struct":
		bs, err := ioutil.ReadAll(response)
		if err != nil {
			return 0, err
		}
		err = json.Unmarshal(bs, responseContainer)
		return responseContainer, err
	}
	return nil, errors.New("cannot convert response")
}
