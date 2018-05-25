package tests

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"github.com/kmilinho/twcli/pkg/httpcli"
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
}
