package core

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "encoding/json"
)

func TestPartialStruct(t *testing.T) {
    type ParentTs struct {
        ParentFieldA string `json:"parentFieldA,omitempty"`
        ParentFieldB int     `json:"testParentField"`
    }
    type StringMap map[string]string
    type StringArray []string
    type Ts struct {
        ParentTs
        FieldA string `json:"fieldA"`
        FieldC StringArray `json:"fieldC"`
        FieldD StringMap    `json:"fieldD"`
    }

    sub := ParentTs{ParentFieldA: "aaa", ParentFieldB: 45}
    m := &Ts{ParentTs: sub, FieldA: "testA", FieldC: []string{"testCa", "testCb"}, FieldD: map[string]string{"testDa": "valueA", "testDb": "valueB"}}

    result := PartialStruct(*m, "testParentField", "fieldD", "fieldC")

    resultJson, _ := json.Marshal(result)
    var resultMap = make(map[string]interface{})
    json.Unmarshal(resultJson, &resultMap)

    expected := "{\"testParentField\":45,\"fieldC\":[\"testCa\",\"testCb\"],\"fieldD\":{\"testDa\":\"valueA\",\"testDb\":\"valueB\"}}"
    var expectedMap = make(map[string]interface{})
    json.Unmarshal([]byte(expected), &expectedMap)

    expectedFail := "{\"fieldC\":[\"testCa\",\"testCb\"],\"fieldD\":{\"testDa\":\"valueA\",\"testDb\":\"valueB\"},\"testParentField\":46}"
    var expectedFailMap = make(map[string]interface{})
    err := json.Unmarshal([]byte(expectedFail), &expectedFailMap)

    if (err != nil) { Log(err.Error()) }

    assert.Equal(t, expectedMap, resultMap)
    assert.NotEqual(t, expectedFailMap, resultMap)

}
