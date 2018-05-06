package httpcli

import (
	"net/http"
	"sync"
	"time"
	"io/ioutil"
)

var httpCliInstance *HTTPCli
var once sync.Once

//TODO: Make it configurable
const requestTimeout = 30

// HTTPCli - HTTP Client
type HTTPCli struct {
	*http.Client
}

// Get - Perform HTTP Get Request.
func Get(url string, headers map[string]string, queryParams map[string]string) ([]byte, error) {

	httCli := getHTTPCli()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	addQueryParams(req, queryParams)

	resp, err := httCli.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func addQueryParams(req *http.Request, queryParams map[string]string) {
	query := req.URL.Query()
	for k, v := range queryParams {
		query.Add(k, v)
	}
	req.URL.RawQuery = query.Encode()
}

func getHTTPCli() *HTTPCli {
	once.Do(func() {
		httpClient := &http.Client{
			Timeout: time.Second * time.Duration(requestTimeout),
		}
		httpCliInstance = &HTTPCli{httpClient}
	})
	return httpCliInstance
}