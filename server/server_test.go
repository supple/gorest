package server

import (
    "testing"
    a "github.com/stretchr/testify/assert"
    "fmt"
    "net/http/httptest"
    "net/http"
    "github.com/supple/gorest/tests"
)
//
//import (
//    "testing"
//    "net/http"
//    "net/http/httptest"
//)
//
//func TestHandleIndexReturnsWithStatusOK(t *testing.T) {
//    request, _ := http.NewRequest("GET", "/", nil)
//    response := httptest.NewRecorder()
//
//
//    ts := httptest.NewServer(GetMainEngine())
//
//    IndexHandler(response, request)
//
//    if response.Code != http.StatusOK {
//        t.Fatalf("Non-expected status code%v:\n\tbody: %v", "200", response.Code)
//    }
//}


// package main

// server.go
// This is where you create a gin.Default() and add routes to it

func init() {
    // Init storage instances
    tests.GetStorage()
    tests.CreateTestCustomer()
}

// router_test.go
func TestDeviceHandler_Update(t *testing.T) {
    testRouter := SetupRouter()
    apiKey := tests.CreateTestApiKey()

    req, err := http.NewRequest("GET", "/api/v1/devices/1", nil)
    req.Header.Add("API-KEY", apiKey.ApiKey)
    if err != nil {
        fmt.Println(err.Error())
    }

    resp := httptest.NewRecorder()
    testRouter.ServeHTTP(resp, req)
    a.Equal(t, "{\"error\":\"Object not found\"}\n", resp.Body.String())
}
