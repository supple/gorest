package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/supple/gorest/core"
	"github.com/supple/gorest/resources"
	"github.com/supple/gorest/tests"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIntegration_Apps_Get(t *testing.T) {
	tests.GetTestStorage()
	resources.CreateCustomer(tests.TEST_CUSTOMER)

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

	expected := fmt.Sprintf("{\"createdAt\":\"%s\",\"id\":\"abc\",\"name\":\"\",\"updatedAt\":\"%s\"}", app.CreatedAt, app.UpdatedAt)
	assert.Equal(t, expected, resp)
}
