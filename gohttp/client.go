package gohttp

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

const (
	POST   = "POST"
	PUT    = "PUT"
	GET    = "GET"
	DELETE = "DELETE"
)

type HttpResponseEntity struct {
	Code  string      `json:"code"`
	Data  interface{} `json:"data"`
	Msg   string      `json:"message"`
	Extra interface{} `json:"extra"`
}

type GoHttpRequest struct {
	URL           string
	Host          string
	Body          []byte
	Header        map[string]string
	Data          map[string]interface{}
	HttpClient    *http.Client
	CheckRedirect func(r *http.Request, v []*http.Request) error
	//FormData         url.Values
	//QueryData        url.Values
	//RawStringData    string
	//RawBytesData     []byte
	//FilePath         string
	//FileParam        string
	//Client           *http.Client
	//Transport        *http.Transport
	//Cookies          []*http.Cookie
	//Errors           []error
	//BasicAuth        struct{ Username, Password string }
	//Debug            bool
	//CurlCommand      bool
	//logger           *log.Logger
	//retry            *RetryConfig
	//bindResponseBody interface{}
}

func New(host, url string) *GoHttpRequest {
	gr := &GoHttpRequest{
		Host:       host,
		URL:        url,
		HttpClient: getDefaultHttpClient(),
		Data:       make(map[string]interface{}),
		Header:     make(map[string]string),
		//FormData:         url.Values{},
		//QueryData:        url.Values{},
		//Client:           nil,
		//Transport:        &http.Transport{},
		//Cookies:          make([]*http.Cookie, 0),
		//Errors:           nil,
		//BasicAuth:        struct{ Username, Password string }{},
		//Debug:            false,
		//CurlCommand:      false,
		//logger:           log.New(os.Stderr, "[goreq]", log.LstdFlags),
		//retry:            &RetryConfig{RetryCount: 0, RetryTimeout: 0, RetryOnHTTPStatus: nil},
		//bindResponseBody: nil,
	}
	return gr
}

func getDefaultHttpClient() *http.Client {
	client := &http.Client{
		Timeout: 20 * time.Second,
	}
	return client
}

func (ghr GoHttpRequest) getURL() string {
	return ghr.Host + ghr.URL
}

func (ghr GoHttpRequest) HttpGet() (*http.Response, error) {
	req, err := http.NewRequest(GET, ghr.getURL(), nil)
	if err != nil {
		return nil, err
	}
	//req.Header.Add("Content-Type", `application/json`)
	req.Header.Set("Content-Type", "application/json")
	return ghr.HttpClient.Do(req)
}

func (ghr GoHttpRequest) HttpPost() (*http.Response, error) {
	req, err := http.NewRequest(POST, ghr.getURL(), bytes.NewBuffer(ghr.Body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return ghr.HttpClient.Do(req)
}

func GetHttpResponseEntity(response *http.Response) (HttpResponseEntity, error) {
	var res HttpResponseEntity
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(response.Body)
	if err != nil {
		return res, nil
	}
	if err := json.Unmarshal(buf.Bytes(), &res); err != nil {
		return res, nil
	}
	return res, nil
}
