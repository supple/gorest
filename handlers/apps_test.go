package handlers

import (
    "github.com/gin-gonic/gin"
    "testing"
    "net/http/httptest"
    "net/http"
    "github.com/stretchr/testify/assert"
    "fmt"
    "github.com/supple/gorest/core"
    "github.com/supple/gorest/tests"
    "github.com/supple/gorest/resources"
)


func init() {
    // Init storage instances
    tests.GetTestStorage()
    resources.CreateCustomer(tests.TEST_CUSTOMER)
}

func TestAppApi_Get(t *testing.T) {
    //gin.SetMode(gin.TestMode)
    r := gin.New()
    var cc = &core.CustomerContext{ApiKey: "", CustomerName: tests.TEST_CUSTOMER}
    app, _ := resources.CreateAndroidApp(cc, "", "abc")

    var path = "/api/v1/apps/abc"
    r.GET("/api/v1/apps/:id", func(c *gin.Context) {
        //setup test and call endpoint handler here
        c.Set("cc", cc)
        appApi := AppApi{}
        appApi.Get(c)
    })

    // make request

    // undefined: httptest.NewRequest in go v1.6
    //req := httptest.NewRequest("GET", path, nil)
    req, _ := http.NewRequest("GET", path, nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    // check response
    resp := w.Body.String()
    fmt.Println(resp)
    assert.Equal(t, 200, w.Code)

    expected := fmt.Sprintf("{\"createdAt\":\"%s\",\"id\":\"abc\",\"name\":\"\",\"updatedAt\":\"%s\"}\n", app.CreatedAt, app.UpdatedAt)
    assert.Equal(t, expected, resp)
}

