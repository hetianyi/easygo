package httpx

import (
	"fmt"
	"github.com/hetianyi/easygo/file"
	"github.com/hetianyi/easygo/logger"
	jsoniter "github.com/json-iterator/go"
	"reflect"
	"testing"
)

func TestMock_RequestBody(t *testing.T) {
	reader, _ := file.GetFile("C:\\Users\\Jason\\Downloads\\JetBrains 2022 最新版本全家桶激活.zip")
	_ = []FormItem{
		{
			Type:  FormItemTypeField,
			Name:  "userName",
			Value: "张三",
		},
		{
			Type:       FormItemTypeFile,
			Name:       "userName",
			FileName:   "张三",
			FileReader: reader,
		},
	}
	p := map[string][]string{
		"": []string{"", ""},
	}
	Mock().RequestBody(p)
}

type ResultContainer struct {
	Name string
}

func Test1(t *testing.T) {
	var resp = &ResultContainer{}
	logger.Info(checkResponseType(resp))
	typ := reflect.TypeOf(resp)
	logger.Info(typ)
	logger.Info(typ.Kind() == reflect.Ptr)
	logger.Info(typ.Elem())
	for {
		if typ.Kind() == reflect.Ptr {
			typ = typ.Elem()
			continue
		}
		break
	}
}

func TestMock_Get(t *testing.T) {
	/*var ret string
	res, status, _ := Mock().Get("https://pti-ms-test.nuanyangjk.com/static/js/app.642a80de.js").
		Success(&ret, 200).
			ErrorCallback(func(status int, response []byte) {
				logger.Fatal("请求错误，状态：", status, ": ", string(response))
			}).
		Do()
	logger.Info(status)
	logger.Info("请求返回：\n", res)*/

	logger.SetColorable(false)

	type BilibiliResponse struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		TTL     int    `json:"ttl"`
		Data    struct {
			GotoType  int    `json:"goto_type"`
			GotoValue string `json:"goto_value"`
			Id        int64  `json:"id"`
			Name      string `json:"name"`
			Seid      string `json:"seid"`
			ShowName  string `json:"show_name"`
			Type      int    `json:"type"`
			Url       string `json:"url"`
		} `json:"data"`
	}
	res, status, err := Mock().Get("https://api.bilibili.com/x/web-interface/search/default").
		Success(&BilibiliResponse{}, 200).
		Do()
	if err != nil {
		logger.Fatal(err)
	}
	logger.Info(status)
	marshal, _ := jsoniter.MarshalToString(res)
	fmt.Println(marshal)
	createFile, _ := file.CreateFile("D:/demo.txt")
	createFile.WriteString(marshal)
	createFile.Close()
	logger.Info("请求返回：\n", marshal)
	logger.Info("asdasdasdasdasdasdasdasdasdasdasdasdasdasdasdasdas{dasdasdasdasdasdasdasdasdasd}asdasdas\"dasdasdasdasdasdasdasdasdasdasdasdasdasdasdasdasdasdasdasdasdasdasdasdasdasdasdasdasdasdasdasdasdasdasdasdasdasdasdasdasdasdasdasdasdasdasdasdasdasd\"")
}

/*
type Response struct {
	Detail string `json:"detail"`
}

func TestMockGet1(t *testing.T) {
	result, _, err := httpx.Mock().URL("https://hub.docker.com/v2/user/").Get().Success(&Response{}, 200, 401).ErrorCallback(func(status int, response []byte) {
		log.Fatal("status ", status, ", response: ", string(response))
	}).Do()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println(result.(*Response).Detail)
	}
}

func TestMockGet2(t *testing.T) {
	result, _, err := httpx.Mock().URL("https://hub.docker.com/v2/user/").Get().Success("", 200, 401).ErrorCallback(func(status int, response []byte) {
		log.Println("status ", status, ", response: ", string(response))
	}).Do()
	if err != nil {
		log.Println(err)
	} else {
		log.Println(result)
	}
}

type Edition struct {
	Label string `json:"label"`
	Name  string `json:"name"`
}

type Response1 struct {
	Editions []Edition
}

func TestMockGet3(t *testing.T) {
	result, _, err := httpx.Mock().URL("https://hub.docker.com/api/content/v1/platforms").Get().Success(&Response1{}).ErrorCallback(func(status int, response []byte) {
		log.Println("status ", status, ", response: ", string(response))
	}).Do()
	if err != nil {
		log.Println(err)
	} else {
		bs, _ := json.MarshalIndent(result, "", " ")
		log.Println(string(bs))
	}
}

func TestMockGet4(t *testing.T) {
	result, _, err := httpx.Mock().URL("https://hub.docker.com/v2/repositories/library/redix/tags/").
		Parameter("page_size", "25").
		Parameter("page", "1").
		Get().Success(new(map[string]interface{})).ErrorCallback(func(status int, response []byte) {
		log.Println("status ", status, ", response: ", string(response))
	}).Do()
	if err != nil {
		log.Println(err)
	} else {
		bs, _ := json.MarshalIndent(result, "", " ")
		log.Println(string(bs))
	}
}

func TestMockGet5(t *testing.T) {
	result, _, err := httpx.Mock().URL("https://hub.docker.com/api/content/v1/products/images/redix").
		Get().Success("").ErrorCallback(func(status int, response []byte) {
		log.Println("status ", status, ", response: ", string(response))
	}).Do()
	if err != nil {
		log.Println(err)
	} else {
		log.Println(result)
	}
}

func TestMockPost1(t *testing.T) {
	log.Println("start ", time.Now().UTC())
	result, _, err := httpx.Mock().URL("http://cdn.yifuls.com/api/cdn/detail").
		Post().
		Body(map[string][]string{"id": {"10004"}}).
		Success(new(map[string]interface{})).ErrorCallback(func(status int, response []byte) {
		log.Println("status ", status, ", response: ", string(response))
	}).Do()
	if err != nil {
		log.Println(err)
	} else {
		bs, _ := json.MarshalIndent(result, "", " ")
		log.Println(string(bs))
	}
	log.Println("end ", time.Now().UTC())
}

func TestMockPost2(t *testing.T) {
	log.Println("start ", time.Now().UTC())
	result, _, err := httpx.Mock().URL("http://cdn.yifuls.com/api/cdn/detail").
		Post().
		Body(map[string][]string{"id": {"10004"}}).
		Success(nil).ErrorCallback(func(status int, response []byte) {
		log.Println("status ", status, ", response: ", string(response))
	}).Do()
	if err != nil {
		log.Println(err)
	} else {
		bs, _ := json.MarshalIndent(result, "", " ")
		log.Println(string(bs))
	}
	log.Println("end ", time.Now().UTC())
}

func TestMockPost3(t *testing.T) {
	out, err := file.CreateFile("D:\\tmp\\轻松一刻语音版-所谓塑料同事-下班缘分就尽了.mp3")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	log.Println("start ", time.Now().UTC())
	result, _, err := httpx.Mock().URL("http://mobilepics.ws.126.net/UIg9y7iEZIrCxoikpo3HeNeikjLMqkV7%3D%3DFTRPUTL6.mp3").
		Get().
		Success(out).ErrorCallback(func(status int, response []byte) {
		log.Println("status ", status, ", response: ", string(response))
	}).Do()
	if err != nil {
		log.Println(err)
	} else {
		bs, _ := json.MarshalIndent(result, "", " ")
		log.Println(string(bs))
	}
	log.Println("end ", time.Now().UTC())
}

func TestMultipart(t *testing.T) {

	httpClient := &http.Client{
		Timeout: time.Second * 20,
	}

	r, w := io.Pipe()
	m := multipart.NewWriter(w)
	go func() {
		defer w.Close()
		defer m.Close()

		o, _ := m.CreateFormField("Name")
		o.Write([]byte("zhangsan"))
		o, _ = m.CreateFormField("Name")
		o.Write([]byte("lisi"))

		fi, _ := file.GetFile("E:\\godfs-storage\\123.zip")
		defer fi.Close()
		o, _ = m.CreateFormFile("secrets", "123.zip")
		io.Copy(o, fi)
	}()

	req, err := http.NewRequest("POST", "http://localhost:8001/upload", r)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Content-Type", m.FormDataContentType())

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(bs))
}

func TestMockUpload(t *testing.T) {
	httpx.SetTTL(0)
	log.Println("start ", time.Now().UTC())
	result, _, err := httpx.Mock().URL("http://localhost:8001/upload").
		Success(nil).
		ErrorCallback(func(status int, response []byte) {
			log.Println("status ", status, ", response: ", string(response))
		}).
		Multipart(func(writer *multipart.Writer) {
			o, _ := writer.CreateFormField("Name")
			o.Write([]byte("zhangsan"))
			o, _ = writer.CreateFormField("Name")
			o.Write([]byte("lisi"))

			fi, _ := file.GetFile("F:\\Software\\fastdfs_client_v1.24.jar")
			defer fi.Close()
			o, _ = writer.CreateFormFile("secrets", filepath.Base("F:\\Software\\fastdfs_client_v1.24.jar"))
			io.Copy(o, fi)
		}).
		Do()
	if err != nil {
		log.Println(err)
	} else {
		bs, _ := json.MarshalIndent(result, "", " ")
		log.Println(string(bs))
	}
	log.Println("end ", time.Now().UTC())
}*/
