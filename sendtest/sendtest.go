package sendtest

import "net/http"

func TestNet() *http.Response {
    resp, err := http.Get("http://google.com")
    if err != nil {
        panic(err)
    }
    return resp
}