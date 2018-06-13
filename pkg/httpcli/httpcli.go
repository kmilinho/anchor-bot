package httpcli

import (
	"net/http"
	"time"
	"io/ioutil"
	"fmt"
	"encoding/json"
	"bytes"
)

//TODO: How to avoid package scoped vars
var httpCliInstance *http.Client

//TODO: Make it configurable
const requestTimeout = 30

// Get - Perform HTTP Get Request.
func Get(url string, headers map[string]string, queryParams map[string]string) ([]byte, error) {
	return httReq(url, "GET", headers, queryParams, nil)
}

func Post(url string, headers map[string]string, queryParams map[string]string, reqBody interface{}) ([]byte, error) {
	bodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("marshalling the request: %v", err)
	}
	return httReq(url, "POST", headers, queryParams, bodyBytes)
}

func httReq(url string, method string, headers map[string]string, queryParams map[string]string, reqBody []byte) ([]byte, error) {

	req, err := http.NewRequest(method, url, bytes.NewReader(reqBody))
	if err != nil {
		return nil, fmt.Errorf("creating the request: %v", err)
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}
	addQueryParams(req, queryParams)

	httCli := getHTTPCli()
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
	if httpCliInstance == nil {
		httpCliInstance = &http.Client{
			Timeout: time.Second * time.Duration(requestTimeout),
		}
	}
	return httpCliInstance
}
