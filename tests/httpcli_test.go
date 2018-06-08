package tests

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"github.com/kmilinho/twcli/pkg/httpcli"
	"encoding/json"
	"io/ioutil"
)

func TestGet(t *testing.T) {

	headerKey := "x-tests-header"
	headerVal := "abc"
	headers := map[string]string{
		headerKey: headerVal,
	}

	queryKey := "tests-param"
	queryVal := "tests-value"
	queryParams := map[string]string{
		queryKey: queryVal,
	}

	checkRequest := func(w http.ResponseWriter, r *http.Request) {
		receivedHeader := r.Header.Get(headerKey)
		if receivedHeader != headerVal {
			t.Errorf("Missing expected header: key: %v value: %v", headerKey, headerVal)
		}

		receivedParam := r.URL.Query().Get(queryKey)
		if receivedParam != queryVal {
			t.Errorf("Missing expected query param: key: %v value: %v", queryKey, queryVal)
		}
	}

	testServer := httptest.NewServer(http.HandlerFunc(checkRequest))

	httpcli.Get(testServer.URL, headers, queryParams)
	defer testServer.Close()
}

type SomeTestStruct struct {
	PropA string `json:"prop_a"`
	PropB int `json:"prop_b"`
}

func TestPost(t *testing.T) {

	headerKey := "x-tests-header"
	headerVal := "abc"
	headers := map[string]string{
		headerKey: headerVal,
	}

	queryKey := "tests-param"
	queryVal := "tests-value"
	queryParams := map[string]string{
		queryKey: queryVal,
	}

	checkRequest := func(w http.ResponseWriter, r *http.Request) {
		receivedHeader := r.Header.Get(headerKey)
		if receivedHeader != headerVal {
			t.Errorf("Missing expected header: key: %v value: %v", headerKey, headerVal)
		}

		receivedParam := r.URL.Query().Get(queryKey)
		if receivedParam != queryVal {
			t.Errorf("Missing expected query param: key: %v value: %v", queryKey, queryVal)
		}

		defer r.Body.Close()

		bytesBodyReq, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Error("error reading request body")
		}
		var reqBody SomeTestStruct
		json.Unmarshal(bytesBodyReq, &reqBody)
		if reqBody.PropA != "mambo" || reqBody.PropB != 5 {
			t.Error("error invalid request object")
		}
	}

	testServer := httptest.NewServer(http.HandlerFunc(checkRequest))

	httpcli.Post(testServer.URL, headers, queryParams, SomeTestStruct{"mambo", 5})
	defer testServer.Close()
}