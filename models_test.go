package main

import (
    "testing"
)

type M struct {
    A string
    B int
    B2 int32
    C float64
}

func TestUpdateModel(t *testing.T) {
    mp := make(map[string]interface{})
    mp["a"] = "a"
    mp["b"] = 123
    mp["b2"] = 1234
    mp["c"] = 1.5
    md := &M{}

    var v32 int32 = 1234
    updateModel(md, mp)
    if (md.A != "a") { t.Fatalf("Fail") }
    if (md.B != 123) { t.Fatalf("Fail") }
    if (md.B2 != v32) { t.Fatalf("Fail %d", md.B2) }
    if (md.C != 1.5) { t.Fatalf("Fail") }
}

func TestUpdateDevice(t *testing.T) {
    mp := make(map[string]interface{})
    mp["name"] = "a"
    mp["Version"] = 1234
    md := &Device{}

    var v32 int32 = 1234
    updateModel(md, mp)
    if (md.Name != "a") { t.Fatalf("Fail Name") }
    if (md.Version != v32) { t.Fatalf("Fail Version") }
}