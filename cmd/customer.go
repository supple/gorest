package main
import (
    "net/http"
    "encoding/json"
    "fmt"
)

func cget() (interface{}, error) {
    client := &http.Client{}

    //resp, err := client.Get("http://example.com")
    // ...

    req, err := http.NewRequest("GET", "http://192.168.1.106:8080/api/v1/devices/abc", nil)
    // ...
    req.Header.Add("API-KEY", `zbCrVUXQSLseDVruJIBwYgke-cRaddKsc`)
    req.Header.Add("Content-type", `application/json`)
    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    into := make(map[string]interface{})
    if err := json.NewDecoder(resp.Body).Decode(into); err != nil {
        return nil, err
    }

    return into, nil
}

func main() {
    fmt.Println(cget())
}
