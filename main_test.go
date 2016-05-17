package main

import (
    "encoding/json"
    "io/ioutil"
    "net/http"
    "net/http/httptest"
    "testing"
    "reflect"
    _"fmt"
    _ "strconv"
    _"hash/crc32"
)

func TestRespondSuccess(t *testing.T) {
    names := map[string]interface{}{}
    names["test"] = "ok2"
    s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        jsonResponse(w, names)
    }))
    defer s.Close()

    resp, err := http.Get(s.URL)
    if err != nil {
        t.Fatalf("Error on test request: %s", err)
    }
    body, err := ioutil.ReadAll(resp.Body)
    defer resp.Body.Close()
    if err != nil {
        t.Fatalf("Error reading response body: %s", err)
    }

    if resp.StatusCode != 200 {
        t.Fatalf("Return code %d expected in success response but got %d", 200, resp.StatusCode)
    }
    if h := resp.Header.Get("Content-Type"); h != "application/json" {
        t.Fatalf("Expected Content-Type %q but got %q", "application/json", h)
    }

    var res = map[string]interface{}{}
    if err = json.Unmarshal([]byte(body), &res); err != nil {
        t.Fatalf("Error unmarshaling JSON body: %s", err)
    }

    if !reflect.DeepEqual(&res, &names) {
        t.Fatalf("Expected response \n%v\n but got \n%v\n", res, names)
    }
}

//
//func BenchmarkCrc(b *testing.B) {
//
//    var c uint32
//    var checksum uint32 = 63
//    var a string
//    //cnt map[string][string]
//    for i :=0; i < b.N; i++ {
//        checksum = (crc32.ChecksumIEEE([]byte("send"+strconv.Itoa(i))))
//        //checksum = (crc32.ChecksumIEEE([]byte("marek")))
//        ca := checksum & 0x7
//        c = (checksum) % 8
//        if ca != c {
//            fmt.Printf("Checksum : %d %d %d \n", checksum, c, ca)
//        }
//        //fmt.Printf("Checksum : %d %d \n", checksum, c)
//    }
//    fmt.Printf("Checksum : %d %s \n", c, a)
//}