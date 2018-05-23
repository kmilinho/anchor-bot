package httpcli

import (
	"net/http"
	"sync"
	"time"
	"io/ioutil"
	"fmt"
)

//TODO: How to avoid package scoped vars
var httpCliInstance *http.Client
var once sync.Once

//TODO: Make it configurable
const requestTimeout = 30

// Get - Perform HTTP Get Request.
func Get(url string, headers map[string]string, queryParams map[string]string) ([]byte, error) {

	httCli := getHTTPCli()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("creating the request: %v", err)
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	addQueryParams(req, queryParams)

	resp, err := httCli.Do(req)
	if err != nil {
		return nil, fmt.Errorf("sending the request: %v", err)
	}
	defer resp.Body.Close()

	//TODO: Handle HTTP response codes
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading the response: %v", err)
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

func getHTTPCli() *http.Client {
	var httpCliInstance *http.Client
	once.Do(func() {
		httpCliInstance = &http.Client{
			Timeout: time.Second * time.Duration(requestTimeout),
		}
	})
	return httpCliInstance
}

