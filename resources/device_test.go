package resources


import (
    "testing"
)


func TestDeviceRP_Update(t *testing.T) {
    d := Device{}

    m := make(map[string]interface{})
    m["token"] = "xo"
    m["appVersion"] = "1.2"
    m["customerName"] = "marek"
    m["appId"] = "xos"

    d.Update(m)
    if (d.AppId != "a") { t.Fatalf("Fail name: %s", d.AppId) }
}

