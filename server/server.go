package server

import (
	"github.com/gin-gonic/gin"
	"github.com/supple/gorest/core"
	"github.com/supple/gorest/handlers"
	"github.com/supple/gorest/resources"
)

func AuthMiddleware(c *gin.Context) {
	//log.Print("[x] Request\n")
	// zbCrVUXQSLseDVruJIBwYgke-cRaddKsc
	ac := resources.AccessTo{Resource: "test", Action: "test"}
	cc, err := resources.Auth(c.Request.Header.Get("API-KEY"), ac)

	if err == nil {
		core.Log("Invalid credentials: ", cc.CustomerName, cc.ApiKey)
		c.Set("cc", cc)
		c.Next()
		return
	}

	// handle error
	switch err.(type) {
	case *core.ApiError:
		ae := err.(*core.ApiError)
		c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		c.AbortWithError(ae.Status, ae)
		return
	case *core.ValidationError:
		ae := err.(*core.ValidationError)
		c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		c.AbortWithError(410, ae)
		return
	case error:
		c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		c.AbortWithError(500, err)
		return
	}
}

func ErrorHandler(c *gin.Context) {
	c.Next()
	if len(c.Errors) > 0 {
		c.JSON(-1, c.Errors) // -1 == not override the current error code
	}
}

func SetupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	// set middleware
	//r.Use(gzip.Gzip(gzip.DefaultCompression))
	//r.Use(CORSMiddleware())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(ErrorHandler)
	r.Use(AuthMiddleware)

	// api handlers
	ca := handlers.CustomerApi{}
	d := handlers.DeviceApi{}
	apps := handlers.AppApi{}

	v1 := r.Group("api/v1")
	{
		v1.POST("/events", handlers.HandleEvents)

		//
		v1.GET("/customers", ca.Get)
		v1.POST("/customers", ca.Post)

		//
		v1.GET("/devices/:id", d.Get)
		v1.POST("/devices", d.Post)
		v1.PATCH("/devices", d.Patch)

		//
		v1.POST("/apps", apps.Create)
		v1.GET("/apps/:id", apps.Get)
		v1.GET("/apps", apps.List)
	}
	// customers
	// api-keys
	// apps
	// devices

	return r
}
