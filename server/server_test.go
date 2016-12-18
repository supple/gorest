package server

import (
    "testing"
    a "github.com/stretchr/testify/assert"
    "fmt"
    "net/http/httptest"
    "net/http"
    "github.com/supple/gorest/storage"
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
    storage.SetInstance("crm", storage.NewMongoDB("192.168.1.106:27017", "crm"))
}

// router_test.go
func TestDeviceRP_Update(t *testing.T) {
    testRouter := SetupRouter()

    req, err := http.NewRequest("GET", "/devices/1", nil)
    if err != nil {
        fmt.Println(err)
    }

    resp := httptest.NewRecorder()
    testRouter.ServeHTTP(resp, req)
    a.Equal(t, 200, resp.Code)
}
