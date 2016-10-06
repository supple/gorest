package main

import (
    "testing"
)

type M struct {
    A string
    B int
    B2 int32
    B3 int64
    C float64
}

func TestUpdateModel(t *testing.T) {
    mp := make(map[string]interface{})
    mp["a"] = "a"
    mp["b"] = 123
    mp["b2"] = 1234
    mp["b3"] = 1234
    mp["c"] = 1.5
    md := &M{}

    var v32 int32 = 1234
    var v64 int64 = 1234
    updateModel(md, mp)
    if (md.A != "a") { t.Fatalf("Fail %s", md.A) }
    if (md.B != 123) { t.Fatalf("Fail %d", md.B) }
    if (md.B2 != v32) { t.Fatalf("Fail %d", md.B2) }
    if (md.B3 != v64) { t.Fatalf("Fail %d", md.B3) }
    if (md.C != 1.5) { t.Fatalf("Fail %f", md.C) }
}

func TestUpdateDevice(t *testing.T) {
    mp := make(map[string]interface{})
    mp["name"] = "a"
    mp["Version"] = 1234
    md := &Device{}

    var v32 int32 = 1234
    updateModel(md, mp)
    if (md.Name != "a") { t.Fatalf("Fail name: %s", md.Name) }
    if (md.Version != v32) { t.Fatalf("Fail version: %s", md.Version) }
}
