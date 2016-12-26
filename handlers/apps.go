package handlers

import (
    "github.com/gin-gonic/gin"
    "encoding/json"
    "github.com/supple/gorest/core"
    "github.com/supple/gorest/resources"
    "log"
)

type AppApi struct {
    ApiController
}


//
func (api *AppApi) Get(c *gin.Context) {
    var cc = core.GetCC(c)
    rp := resources.NewDeviceRP(cc)
    id := c.Param("id")
    //var content Content
    //if c.Bind(&content) == nil {
        ret, err := rp.FindOne(id)
        if (handleError(err, c)) {
            return
        }
        c.JSON(200, ret)
        //id, _ := c.Params.Get("id")
        //c.JSON(200, map[string]string{"id": id})
    //}

    //var id = c.URLParams["id"]
    //obj := app.Storage.Get(id)
    //jsonResponse(w, obj)
}

func (api *AppApi) Post(c *gin.Context) {
    var err error
    obj := &resources.Device{}

    decoder := json.NewDecoder(c.Request.Body)
    if err = decoder.Decode(obj); err != nil {
        log.Print(err.Error())
        c.JSON(422, err)
        return
    }
    dRp := resources.NewDeviceRP(core.GetCC(c))
    err = dRp.Create(obj)
    if (handleError(err, c)) {
        return
    }
    c.JSON(201, obj)
    //app.Storage.Set(obj.Id, &obj)
    //jsonResponse(w, obj)
}
