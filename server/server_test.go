package server

import (
	"fmt"
	a "github.com/stretchr/testify/assert"
	"github.com/supple/gorest/core"
	"github.com/supple/gorest/model"
	"github.com/supple/gorest/resources"
	"github.com/supple/gorest/tests"
	"net/http"
	"net/http/httptest"
	"testing"
)

func init() {
	// Init storage instances
	tests.GetTestStorage()
	resources.CreateCustomer(tests.TEST_CUSTOMER)
}

func createTestApiKey() *model.ApiKey {
	var cc = &core.CustomerContext{ApiKey: "", CustomerName: tests.TEST_CUSTOMER}
	apiKey, err := resources.CreateApiKey(cc)
	if err != nil {
		fmt.Println(err.Error())
	}

	return apiKey
}

// router_test.go
func TestDeviceHandler_Update(t *testing.T) {
	testRouter := SetupRouter()
	apiKey := createTestApiKey()

	req, err := http.NewRequest("GET", "/api/v1/devices/1", nil)
	req.Header.Add("API-KEY", apiKey.ApiKey)
	if err != nil {
		fmt.Println(err.Error())
	}

	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	a.Equal(t, "{\"error\":\"Object not found\"}\n", resp.Body.String())
}
