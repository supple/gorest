package handlers

import (
    "github.com/gin-gonic/gin"
    "encoding/json"
    "github.com/supple/gorest/core"
    "github.com/supple/gorest/resources"
    "log"
)

type Content struct {
    Id string `form:"id" binding:"required"`
}

type ApiController struct {

}

type DeviceApi struct {
    ApiController
}


func handleError(err error, c *gin.Context) bool {
    if (err == nil) {
        return false
    }

    switch err.(type) {
    case *core.ApiError:
        ae := err.(*core.ApiError)
        c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
        c.AbortWithError(ae.Status, ae)

    case error:
        c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
        c.AbortWithError(500, err)
    }

    return true
}

//
func (api *DeviceApi) Get(c *gin.Context) {
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

func (api *DeviceApi) Post(c *gin.Context) {
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

type CustomerApi struct {
    ApiController
}

//
func (api *CustomerApi) Get(c *gin.Context) {

    //var cc *core.CustomerContext = &core.CustomerContext{CustomerName: customerName}
    //cRp := r.NewCustomerRP(cc)

    id, _ := c.Params.Get("id")
    c.JSON(200, map[string]string{"id": id})

    //var id = c.URLParams["id"]
    //obj := app.Storage.Get(id)
    //jsonResponse(w, obj)
}
//
//func CampaignList(c web.C, w http.ResponseWriter, r *http.Request) {
//    names := app.Storage.GetByCriteria()
//    jsonResponse(w, names)
//}
//
func (api *CustomerApi) Post(c *gin.Context) {
    obj := resources.Customer{}
    decoder := json.NewDecoder(c.Request.Body)
    if err := decoder.Decode(&obj); err != nil {
        log.Print(err.Error())
        c.JSON(422, err)
        //http.Error(w, http.StatusText(422), 422)
        return
    }
    c.JSON(201, obj)
    //app.Storage.Set(obj.Id, &obj)
    //jsonResponse(w, obj)
}
//
//// https://github.com/quintans/goSQL#update
//func CampaignUpdate(c web.C, w http.ResponseWriter, r *http.Request) {
//    decoder := json.NewDecoder(r.Body)
//    defer r.Body.Close()
//    tmp := make(map[string]interface{})
//    if err := decoder.Decode(&tmp); err != nil {
//        http.Error(w, http.StatusText(422), 422)
//        return
//    }
//
//    var id = c.URLParams["id"]
//    obj := app.Storage.Get(id).(*resources.Device)
//    resources.UpdateModel(obj, tmp)
//
//    app.Storage.Set(id, obj)
//    jsonResponse(w, obj)
//}
