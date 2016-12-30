package core

import (
    "reflect"
    "strings"
)

func fieldSet(fields ...string) map[string]bool {
    set := make(map[string]bool, len(fields))
    for _, s := range fields {
        set[s] = true
    }
    return set
}

func PartialStruct(s interface{}, fields ...string) map[string]interface{} {
    fs := fieldSet(fields...)
    rt := reflect.TypeOf(s)
    rv :=  reflect.ValueOf(s)
    out := make(map[string]interface{})
    for i := 0; i < rt.NumField(); i++ {
        field := rt.Field(i)
        if (field.Type.Kind() == reflect.Struct) {
            sub := PartialStruct(rv.Field(i).Interface(), fields...)
            for k, v := range sub {
                out[k] = v
            }
        }

        jsonKey := strings.Split(field.Tag.Get("json"), ",")[0]
        if fs[jsonKey] {
            out[jsonKey] = rv.Field(i).Interface()
        }
    }

    return out
}
